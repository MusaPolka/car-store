package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"email-service/internal/consumer"
	_ "github.com/joho/godotenv/autoload" // Автозагрузка .env
)

func main() {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	// Обработка сигналов завершения
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalChan
		log.Println("Получен сигнал завершения")
		stop()

		// Таймаут для graceful shutdown
		<-time.After(30 * time.Second)
		log.Fatal("Принудительное завершение")
	}()

	// Запуск потребителя
	go consumer.StartEmailConsumer(ctx)

	log.Println("Сервис запущен. Ожидание сообщений...")
	<-ctx.Done()
	log.Println("Сервис остановлен")
}
