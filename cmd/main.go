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

	"github.com/adminsemy/URLShorting/internal/auth"
	"github.com/adminsemy/URLShorting/internal/github"
	"github.com/adminsemy/URLShorting/internal/server"
	"github.com/adminsemy/URLShorting/internal/shorten"
	storage "github.com/adminsemy/URLShorting/internal/storage/shorting"
	"github.com/adminsemy/URLShorting/internal/storage/user"
)

func main() {
	svc := shorten.NewService(storage.NewInMemory())

	authenticator := auth.NewService(github.NewClient(), user.NewUserInMemory(), "61e6a9adf7a1794b3b5f", "adf45b6969c54308d14818f8eb973eafc731522e")

	server := server.NewServer(svc, authenticator)

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
