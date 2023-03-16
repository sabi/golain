/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available channels",
	Run: func(cmd *cobra.Command, args []string) {
		rows, err := db.Query("SELECT name FROM channels")
		if err != nil {
			fmt.Println("Failed to list channels:", err)
			return
		}
		defer rows.Close()

		fmt.Println
