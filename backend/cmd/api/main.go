package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"pencatatan/internal/app"
	"pencatatan/internal/config"
	"pencatatan/internal/database"
	"pencatatan/internal/server"
	"syscall"
	"time"
)

func graceFullyShutdown(srv *http.Server, done chan bool) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	log.Println("Shutting down gracefully, press Ctrl+C again to force")
	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	done <- true
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal("cannot initialize database:", err)
	}

	container := app.BuildContainer(db)

	srv := server.NewServer(cfg, container)
	done := make(chan bool, 1)

	go graceFullyShutdown(srv, done)

	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Sprintf("http server error: %v", err))
	}

	<-done
	log.Println("Gracefully shutdown complete.")
}
