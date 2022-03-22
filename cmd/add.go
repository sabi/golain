/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Discord Webhook",
	Long: `Add a new Discord Webhook
	golain add --channel General-Chat --webhook https://discord.com/api/123/123`,
	Run: func(cmd *cobra.Command, args []string) {
		line := []byte(Channel + " = " + Webhook + "\n") // put string into byte array
		_, err := os.Stat("lain.conf") // Check for file
		if err != nil { // if file does not exist
			err := ioutil.WriteFile("lain.conf", line, 0600)
			if err != nil {
				fmt.Println(err)
			}
		} else { // if file did exist
			f, err := os.OpenFile("lain.conf", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
		defer f.Close() // Close the file after the block of text is done
		line := Channel + " = " + Webhook + "\n"
		if _, err = f.WriteString(line); err != nil {
			panic(err)
		}
	}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
