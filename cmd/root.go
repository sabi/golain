/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"bufio"
	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var ( 


	Channel string
	Webhook string

rootCmd = &cobra.Command{
	Use:   "golain",
	Short: "golain - cli-to-discord utility",
	Long: `golain - Lain: a cli-to-Discord utility written in Go.
	Let's Love Lain. https://sabisimple.com/foss/lain`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&Channel, "channel", "c", "", "Discord Channel Common Name")
	rootCmd.PersistentFlags().StringVarP(&Webhook, "webhook", "w", "", "Discord Channel Webhook")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


