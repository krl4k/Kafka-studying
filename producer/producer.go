package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type Producer struct {
	w kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		w: kafka.Writer{
			Addr:      kafka.TCP(brokers...),
			Topic:     topic,
			BatchSize: 1,
		},
	}
}

func (p *Producer) SendMessage(ctx context.Context, message []byte) error {
	err := p.w.WriteMessages(ctx, kafka.Message{
		Value: message,
	})

	if err != nil {
		log.Println("producer: failed to write messages:", err)
		return err
	}
	return nil
}
