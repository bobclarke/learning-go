package dialers

import (
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConnectCluster(endpoint string) *kafka.Conn {
	conn, err := kafka.Dial("tcp", endpoint)
	if err != nil {
		log.Fatal("failed to connect to endpoint", err)
	}

	status(conn)

	return (conn)
}

func status(conn *kafka.Conn) {

	a, _ := conn.Brokers()
	fmt.Println(a)

}
