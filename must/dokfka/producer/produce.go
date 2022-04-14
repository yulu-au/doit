package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%v", i)
		wg.Add(1)
		go func(k string) {
			defer wg.Done()
			produce(k)
		}(key)
	}
	wg.Wait()
}

func produce(key string) {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "test",
		//生产端的负载均衡策略
		Balancer: &kafka.Hash{},
	}

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		_, value := fmt.Sprintf("key_%v", i), fmt.Sprintf("val_%v", i)
		msg := kafka.Message{Key: []byte(key), Value: []byte(value)}
		w.WriteMessages(context.Background(), msg)
		fmt.Println("write----------------")
	}

	// err := w.WriteMessages(context.Background(),
	// 	kafka.Message{
	// 		Key:   []byte("Key-A"),
	// 		Value: []byte("Hello World!"),
	// 	},
	// 	kafka.Message{
	// 		Key:   []byte("Key-B"),
	// 		Value: []byte("One!"),
	// 	},
	// 	kafka.Message{
	// 		Key:   []byte("Key-C"),
	// 		Value: []byte("Two!"),
	// 	},
	// )
	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
