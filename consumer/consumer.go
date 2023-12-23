package consumer

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type Consumer struct {
	r *kafka.Reader
}

func NewConsumer(brokers []string, topic string) *Consumer {
	return &Consumer{
		r: kafka.NewReader(
			kafka.ReaderConfig{
				Brokers:   brokers,
				Topic:     topic,
				Partition: 0,

				//MaxBytes:  10e6, // 10MB
				//MaxWait: 1 * time.Second,
				ReadBatchTimeout: 1 * time.Second,
				GroupID:          "consumer-group-id-1", // чтобы кафка могла отдавать нужный offset.
				// если не указать, то будет отдавать только новые сообщения
				// если не указывать GroupID, то нужно будет самому следить за offset
			}),
	}
}

// may lock if no messages
// thats why we need to use goroutine with this method
func (c *Consumer) Consume(ctx context.Context) ([]byte, error) {
	// commit offset if consumer group is used
	m, err := c.r.ReadMessage(ctx)
	if err != nil {
		log.Println("consumer: failed to read message:", err)
		return nil, err
	}
	return m.Value, nil
}
