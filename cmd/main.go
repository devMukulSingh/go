package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tutorial/internal/config"
	"tutorial/internal/http/handlers/hello"
	"tutorial/internal/storage/sqlite"
)

func main() {
	cfg := config.MustLoad()

	storage,err := sqlite.New(cfg)

	if err!= nil{
		log.Fatal(err)
	}

	slog.Info("database initialized",slog.String("env",cfg.Env))
	router := http.NewServeMux()

	router.HandleFunc("GET /", hello.GetHello())
	router.HandleFunc("POST /post", hello.PostHello(storage))


	//setup server
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Printf("Server started at %s", cfg.Address)

	// for handling graceful shutdown
	done := make(chan os.Signal,1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT,syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start server")
		}
	}()
	<-done
	slog.Info("shutting down the server")
	ctx,cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err!= nil{
		slog.Error("Failed to shutdown server", slog.String("errir",err.Error()))
	}
	slog.Info("Server shutdown sucessfully");
}
