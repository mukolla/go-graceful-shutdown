package main

import (
	"context"
	"encoding/json"
	"example.com/m/pkg/config"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	prdLogger, err := zap.NewProduction()

	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer prdLogger.Sync()
	logger := prdLogger.Sugar()

	cnf, err := config.NewConfig()
	if err != nil {
		logger.Fatalw("Failed parse config", "err", err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	srv := &http.Server{
		Handler: router,
		Addr:    cnf.HttpAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//log.Fatal(srv.ListenAndServe())

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Fatalw("Server failed", "err", err, "http-addr", cnf.HttpAddr)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(
		signals,
		syscall.SIGTERM,
		syscall.SIGINT,
	)
	<-signals
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalw("Failed to shutdown server", "err", err, "http-addr", cnf.HttpAddr)
	}

	logger.Infof("Buy!")
}
