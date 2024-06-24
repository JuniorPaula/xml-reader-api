package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"xml-reader-api/internal/config"
	"xml-reader-api/internal/routes"
)

func main() {
	cfg, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal(err)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// create context to control the lifetime of the server
	_, cancel := context.WithCancel(context.Background())

	// Create a new database connection
	db, err := config.ConnectDatabase(cfg.DB_NAME)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Automatically migrate the database
	config.AutoMigrate(db)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: routes.NewRoutes(db, cfg),
	}

	go func() {
		fmt.Printf("Server running on port %s\n", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error while start server\n: %v", err)
		}
	}()

	// wait for the interrupt signal
	<-interrupt
	fmt.Println("Shutting down server...")
	cancel()

	timeout := 5 * time.Second
	ctxShutDown, cancelShutDown := context.WithTimeout(context.Background(), timeout)
	defer func() {
		cancelShutDown()
	}()

	if err := server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server shutdown failed:%+v", err)
	}
	fmt.Print("Server exiting...\n")
}
