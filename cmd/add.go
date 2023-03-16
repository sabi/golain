/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var channelFlag string
var webhookFlag string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new channel and webhook",
	Run: func(cmd *cobra.Command, args []string) {
		if channelFlag == "" || webhookFlag == "" {
			fmt.Println("Both --channel and --webhook flags are required.")
			return
		}

		_, err := db.Exec("INSERT INTO channels (name, webhook, is_default) VALUES (?, ?, ?)", channelFlag, webhookFlag, 0)
		if err != nil {
			fmt.Println("Failed to add channel:", err)
			return
		}

		fmt.Println("Channel added:", channelFlag)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&channelFlag, "channel", "c", "", "Channel name")
	addCmd.Flags().StringVarP(&webhookFlag, "webhook", "w", "", "Webhook URL")
}
