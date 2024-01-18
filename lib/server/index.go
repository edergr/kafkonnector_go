package server

import (
	"context"
	"log"
	"net/http"
)

var server http.Server

func StartServer(ctx context.Context) {
	server = http.Server{Addr: ":8080"}
	log.Println("Server listening on PORT 8080")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Server error:", err)
	}

	<-ctx.Done()
}

func StopServer() {
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("Server shutdown error:", err)
	}
}
