package consumer

import (
	"context"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"os"
	"sync"
	"time"

	"email-service/internal/email" // Правильный импорт
	_ "github.com/rabbitmq/amqp091-go"
)

const maxConcurrent = 5 // Лимит параллельных обработчиков

func StartEmailConsumer(ctx context.Context) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URI"))
	if err != nil {
		log.Fatalf("Ошибка подключения к RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка создания канала: %v", err)
	}
	defer ch.Close()

	// Объявление очереди с параметрами
	_, err = ch.QueueDeclare(
		"emails", // name
		true,     // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("Ошибка объявления очереди: %v", err)
	}

	msgs, err := ch.Consume(
		"emails",
		"",
		false, // auto-ack = false
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Ошибка регистрации потребителя: %v", err)
	}

	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrent)

	for {
		select {
		case <-ctx.Done():
			log.Println("Завершение работы потребителя...")
			wg.Wait()
			return
		case msg := <-msgs:
			sem <- struct{}{}
			wg.Add(1)

			go func(d amqp.Delivery) {
				defer func() {
					<-sem
					wg.Done()
				}()

				var emailData struct {
					To      string `json:"to"`
					Subject string `json:"subject"`
					Body    string `json:"body"`
				}

				if err := json.Unmarshal(d.Body, &emailData); err != nil {
					log.Printf("Ошибка декодирования JSON: %v", err)
					d.Nack(false, true) // Повторная постановка в очередь
					return
				}

				// Попытка отправки с повторами
				for i := 0; i < 3; i++ {
					if err := email.SendEmail(emailData.To, emailData.Subject, emailData.Body); err == nil {
						d.Ack(false)
						return
					}
					time.Sleep(time.Duration(i+1) * time.Second)
				}

				log.Printf("Не удалось отправить письмо: %s", emailData.To)
				d.Nack(false, false) // Отправка в DLQ
			}(msg)
		}
	}
}
