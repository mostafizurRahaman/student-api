package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mostafizurRahaman/student-api/internal/config"
)

func main() {

	// Load the environement variable
	var cfg = config.MustLoad()

	// Setup router:
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello world!"))
	})

	// Setup Server
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	// ? done channel:
	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {

		err := server.ListenAndServe()

		if err != nil {
			slog.Info("Server failed to start", slog.String("error", err.Error()))
		}

	}()



	//  
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	// Shutdown the server:
	err := server.Shutdown(ctx)

	if err != nil {
		slog.Info("Failed to stop server", slog.String("error", err.Error()))
	}

}
