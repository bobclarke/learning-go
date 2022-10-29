package dialers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func OpenConnection(endpoint string, topic string, partition int) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", endpoint, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))

	debug(conn)

	return (conn)
}

func CloseConnection(conn *kafka.Conn) {
	conn.Close()
}

func debug(conn *kafka.Conn) {
	broker := conn.Broker()
	fmt.Println("Broker:", broker)
}
