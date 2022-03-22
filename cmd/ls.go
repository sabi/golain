/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"reflect"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List available Discord Channels",
	Long: `golain ls`,
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
		keys := reflect.ValueOf(lainHooks).MapKeys()
		for _, discordChannel := range keys {
			fmt.Println(discordChannel)}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
