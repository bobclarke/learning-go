package consumers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func ConsumeMessage(message string) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "workflow01-stg-g1mg-kafka.az.eu-az-stg-mgt.gdpdentsu.net:2100", "devops-test", 0)
	if err != nil {
		log.Fatal("Unable to connect", err)
	}
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))

	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}

	if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
