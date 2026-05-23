package main

import (
	"context"
	"fmt"

	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/mostafizurRahaman/customer-api/internal/config"
	"github.com/mostafizurRahaman/customer-api/internal/http/handlers/customers"
)

func main() {
	fmt.Println("Go  server is cooking....")

	//? Load env:
	var cfg = config.MustLoad()

	fmt.Println(cfg)

	// ? Setup  router:

	router := http.NewServeMux()

	router.HandleFunc("POST /api/customers", customers.New())

	//  Setup database :

	// Setup server:
	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	done := make(chan os.Signal, 1)

	// ! Notify the channel for singnal receive
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("Failed to run server:", slog.String("", err.Error()))
		}
	}()

	fmt.Println(color.GreenString("Server is running on : %s", cfg.Address))

	<-done

	// ! Setup  a context for main  function:
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	// ! Call the defer instant :
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server has rejected to shutdown", slog.String("error", err.Error()))
	}

}
