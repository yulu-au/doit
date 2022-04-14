package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ReadFromKafka(i)
		}(i)
	}
	wg.Wait()
}

func ReadFromKafka(id int) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test",
		// Partition: 0,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		GroupID:  "g2",
		// CommitInterval: time.Second,
	})
	ctx := context.Background()
	for {
		r.ReadMessage()
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("id:%v message at offset %d: %s = %s\n", id, m.Offset, string(m.Key), string(m.Value))
		err = r.CommitMessages(context.Background(), m)
		if err != nil {
			log.Fatal("failed to commit msg:", err)
		}
	}

	// if err := r.Close(); err != nil {
	// 	log.Fatal("failed to close reader:", err)
	// }
}
