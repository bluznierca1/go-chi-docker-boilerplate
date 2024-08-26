package main

import (
	"context"
	"fmt"
	"log"
	"myapp/internal/api/handler"
	"myapp/internal/api/router"
	"myapp/internal/database"
	"myapp/internal/ent"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// initialize vars from .env
	initializeDotEnv()

	var apiPort string
	if apiPort = os.Getenv("API_PORT"); apiPort == "" {
		log.Panic("Missing value for API_PORT env var")
	}

	mysqlClient, err := database.InitMySqlDb()
	if err != nil {
		log.Fatalf("error while connecting db: %v", err)
	}

	apiHandlers := handler.InitHandlers() // at some point there might be error and you might need to inject services

	router := router.InitRouter(*apiHandlers)

	srv := &http.Server{
		Addr:    ":" + apiPort,
		Handler: router,
	}
	fmt.Println("running on port " + apiPort)

	go func() {
		// Let's start the server in new goroutine to allow it to listen to termination
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("error", err.Error())
			log.Fatalf("critical error, server down: %s\n", err)
		}
	}()

	gracefulShutdown(srv, mysqlClient)

}

func gracefulShutdown(srv *http.Server, dbConn *ent.Client) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	// close db
	database.CloseDbConnection(dbConn)

	log.Println("Exiting server gracefully...")

}

// initializeDotEnv initializes all vars from .env file
// make sure to provide correct path
func initializeDotEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not initialize .env file %v", err)
	}
}
