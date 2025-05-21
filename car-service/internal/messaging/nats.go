package messaging

import (
	"log"

	"github.com/nats-io/nats.go"
)

func NewNATSConnection() *nats.Conn {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	return nc
}

// Publisher: Publish a message to a subject
func Publish(nc *nats.Conn, subject string, msg []byte) error {
	return nc.Publish(subject, msg)
}

// Subscriber: Subscribe to a subject and handle messages
func Subscribe(nc *nats.Conn, subject string, handler nats.MsgHandler) (*nats.Subscription, error) {
	return nc.Subscribe(subject, handler)
}
