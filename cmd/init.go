package main

import (
	"fmt"
	"github.com/zero-remainder/go-ranker/config"
	"github.com/zero-remainder/go-ranker/database"
	"log"
)

func init() {
	initViper()
	connectDB()
}

func initViper() {
	var err error
	AppConfig, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	fmt.Println("Configuration loaded!")
}

func connectDB() {
	database.ConnectMongo(AppConfig)
}
