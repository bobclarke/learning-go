package producers

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func ProduceMessage(message string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "workflow01-stg-g1mg-kafka.az.eu-az-stg-mgt.gdpdentsu.net:2100", "devops-test", 0)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	_, err = conn.WriteMessages(kafka.Message{Value: []byte(message)})

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

}
