package cmd

import (
	"fmt"

	"github.com/bobclarke/learning-go/kafka-util/pkg/dialers"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Establish a connection to a broker",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Opening connection")
		conn := dialers.ConnectCluster("workflow01-stg-g1mg-kafka.az.eu-az-stg-mgt.gdpdentsu.net:2100")

		fmt.Println("Closing connection")
		dialers.CloseConnection(conn)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
