package kafka

import (
	"context"
	"sync"
	"os"
	"log"
	"time"
	"github.com/segmentio/kafka-go"

)

func PushMessage(ctx context.Context, key, value []byte) (err error) {
	w := &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_IP")),
		Topic:   "device-updates",
		Balancer: &kafka.LeastBytes{},
	}

	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}  
	return w.WriteMessages(context.Background(), message)
}


func ReadMessage(ctx context.Context, wg *sync.WaitGroup){
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_IP")},
		// GroupID:   "application-server",
		Topic:     "device-requirements",
		MinBytes:  10e2, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	for {
		select {
			case <-ctx.Done():
				closeReader(r)
				wg.Done()
				return
			default:
				// The same context needs to be passed so that we can terminate on Ctrl C gracefully
				m, err := r.ReadMessage(ctx)
				if err != nil {
					closeReader(r)
					break
				}
				log.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		}
	}
}

func closeReader(r *kafka.Reader) {
	log.Println("Closing Kafka Reader")
	if err := r.Close(); err != nil {
		log.Println("Failed to close reader:", err)
	}
	return
}

