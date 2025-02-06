package main

import (
	"fmt"
	"github.com/zero-remainder/go-ranker/config"
)

var AppConfig *config.Config

func main() {
	startServer(":" + fmt.Sprint(AppConfig.Port))
}
