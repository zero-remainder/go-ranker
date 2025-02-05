package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/zero-remainder/go-ranker/config"
	"github.com/zero-remainder/go-ranker/internal/routes"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	initViper()
	initFiber(":" + fmt.Sprint(AppConfig.Port))
}

func initViper() {
	var err error
	AppConfig, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
}

func initFiber(port string) {
	app := fiber.New()
	routes.SetupPublicRoutes(app)
	go func() {
		if err := app.Listen(port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	fmt.Println("Listening on port", port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	fmt.Println("Server exited properly")
}
