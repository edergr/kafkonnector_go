package main

import (
	"context"
	"kafkonnector_go/api/connectors"
	"kafkonnector_go/commons/database"
	"kafkonnector_go/config"
	"kafkonnector_go/lib/server"
	"log"
	"os"
	"os/signal"
)

func main() {
	cfg := config.NewConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dropApp := make(chan os.Signal, 1)
	signal.Notify(dropApp, os.Interrupt)

	go func() {
		<-dropApp
		log.Println("Shutdown in progress...")

		database.DisconnectMongoDB()
		server.StopServer()

		cancel()
	}()

	database.ConnectMongoDB(cfg.MongoURI)
	connectors.Router()
	server.StartServer(ctx)

	<-ctx.Done()
	return
}
