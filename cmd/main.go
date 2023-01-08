package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adminsemy/URLShorting/internal/server"
	"github.com/adminsemy/URLShorting/internal/shorten"
	"github.com/adminsemy/URLShorting/internal/storage"
)

func main() {
	svc := shorten.NewService(storage.NewInMemory())
	server := server.NewServer(svc)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := http.ListenAndServe("0.0.0.0:8080", server); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error running server: %v", err)
		}
	}()

	log.Println("server started")
	<-quit

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("error closing server: %v", err)
	}

	log.Println("server stopped")
}
