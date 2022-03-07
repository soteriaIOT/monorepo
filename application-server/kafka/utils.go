package kafka

import (
	"context"
	"os"
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


