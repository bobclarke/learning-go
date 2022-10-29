/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/bobclarke/learning-go/kafka-util/pkg/producers"
	"github.com/spf13/cobra"
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Send a message to Kafka",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		customMessage, _ := cmd.Flags().GetString("message")
		defaultMessage := "Default message"
		if customMessage != "" {
			fmt.Printf("Sending '%s' to Kafka\n", customMessage)
			producers.ProduceMessage(customMessage)
		} else {

			fmt.Printf("Sending '%s' to Kafka\n", defaultMessage)
			producers.ProduceMessage(defaultMessage)
		}
	},
}

func init() {
	rootCmd.AddCommand(produceCmd)
	produceCmd.PersistentFlags().String("message", "", "A message string")
}
