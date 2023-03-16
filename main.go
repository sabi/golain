/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var defaultChannel string

func init() {
	// Set up the SQLite database
	var err error
	db, err = sql.Open("sqlite3", "./golain.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create the channels table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS channels (name TEXT PRIMARY KEY, webhook TEXT, is_default INTEGER)")
	if err != nil {
		log.Fatal(err)
	}

	// Get the default channel
	row := db.QueryRow("SELECT name FROM channels WHERE is_default = 1")
	row.Scan(&defaultChannel)
}
