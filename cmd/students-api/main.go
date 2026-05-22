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

	"github.com/mostafizurRahaman/student-api/internal/config"
	"github.com/mostafizurRahaman/student-api/internal/http/handlers/student"
)

func main() {
	// Load env: 
	cfg := config.MustLoad()

	// Setup Database : 	

	// Setup Router : 
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())
	


	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)


	// Setup server: 	
	server := http.Server{ 
		Addr: cfg.Address,
		Handler: router,
	}

	go func(){
		err:= 	server.ListenAndServe()
		if err != nil { 
			log.Fatalf("Server is failed to  start : %s", err.Error())
		}
	}()
	
	fmt.Println("Server is running on :", cfg.Address)

	<-done

	slog.Info("Shutting down server!")

	ctx, cancel := context.WithTimeout(context.Background(), 5  * time.Second)

	defer cancel()

	if	err:= server.Shutdown(ctx); err != nil { 
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}

	fmt.Println("Welcome to student api.")
}


// ! Steps are: 
// 1. Load config
// 2. Create router with server mux : 
// 3. Create routes 
// 4. Configure Server Struct 
// 5. Implement gracefull stop
// 6. Implement Grace full shutdown. 


