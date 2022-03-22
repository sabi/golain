/*
Copyright © 2022 Sabi
*/
package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"github.com/spf13/cobra"
)

// msgCmd represents the msg command
var msgCmd = &cobra.Command{
	Use:   "msg",
	Short: "Post a message to Discord",
	Long: `golain msg -c CHANNELNAME "Your message goes here"
	EX: golain msg -c sabi_engineering "Good job everybody!"`,
	Run: func(cmd *cobra.Command, args []string) {
		lainHooks := make(map[string]string)
		lines, err := readLines("lain.conf")
		if err != nil {
			return
		}
		for _, line := range lines {
			split := strings.Split(line, " = ")
			lainHooks[split[0]] = split[1]
		}

		data := url.Values{"content":args[0:]}//TODO put the value of the rest of the passed in string
		fmt.Println(data)

		webhook := lainHooks[Channel]
		_ , errB := http.PostForm(webhook, data)
		if errB != nil {
			fmt.Println(errB)
		}
		},
	}

func init() {
	rootCmd.AddCommand(msgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// msgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// msgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
