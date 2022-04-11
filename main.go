package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MalukiMuthusi/wallet-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	v1 := r.Group("/v1")
	{
		wallets := v1.Group("/wallets")
		{
			wallets.GET(":wallet_id/balance", handlers.BalanceHandler)
			wallets.POST(":wallet_id/credit", handlers.CreditHandler)
			wallets.POST(":wallet_id/debit", handlers.DebitHandler)
		}
	}

	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {

		if err := srv.ListenAndServe(); err != nil {
			log.Print("failed to start")
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	var wait time.Duration

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.

	log.Printf("Shutting down")
	os.Exit(0)

}
