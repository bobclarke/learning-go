package cmd

import (
	"fmt"

	"github.com/bobclarke/learning-go/kafka-util/pkg/dialers"
	"github.com/spf13/cobra"

	"github.com/segmentio/kafka-go"
)

var conn *kafka.Conn

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Establish a connection to a broker",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Opening connection")
		conn = dialers.ConnectCluster("workflow01-stg-g1mg-kafka.az.eu-az-stg-mgt.gdpdentsu.net:2100")

		fmt.Println("Closing connection")
		dialers.CloseConnection(conn)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
