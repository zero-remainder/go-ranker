package main

import (
	"fmt"
	"github.com/zero-remainder/go-ranker/config"
	"log"
)

func main() {

}

func init() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	fmt.Println("App Name:", cfg.AppName)
	fmt.Println("Port:", cfg.Port)
	fmt.Println("Debug:", cfg.Debug)
}
