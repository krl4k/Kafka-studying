package main

import (
	"fmt"
	"golang.org/x/net/context"
	"kafka_studying/consumer"
	"kafka_studying/producer"
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		startProducer()
	}()

	go func() {
		defer wg.Done()

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		startConsumer(ctx)
	}()
	wg.Wait()
}

func startProducer() {
	p := producer.NewProducer([]string{"localhost:9092"}, "test")

	messageCounter := 0
	for {
		<-time.After(2 * time.Second)

		msg := "message: " + strconv.Itoa(messageCounter)
		err := p.SendMessage(context.Background(), []byte(msg))
		if err != nil {
			log.Println("producer: failed to write messages:", err)
			// dead letter queue
			// or retry
			// or send to another topic
		}
		fmt.Println("producer: sent message: " + msg)
		messageCounter++
	}
}

func startConsumer(ctx context.Context) {
	c := consumer.NewConsumer([]string{"localhost:9092"}, "test")

	//go func() {
	for {
		select {
		case <-time.After(500 * time.Millisecond):
			fmt.Println("consumer: trying to read message")
			message, err := c.Consume(ctx)
			if err != nil {
				log.Println("consumer: failed to read message:", err)
				continue
			}
			fmt.Println("consumer: obtained message: " + string(message))
		case <-ctx.Done():
			fmt.Println("consumer: context done")
			return
		}
	}

}
