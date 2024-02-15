package main

import (
	"context"
	"kafkonnector_go/commons/database"
	"kafkonnector_go/config"
	"kafkonnector_go/lib/connectors/routes"
	"kafkonnector_go/lib/connectors/service"
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

	if err := database.ConnectMongoDB(cfg.MongoURI); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer database.DisconnectMongoDB()

	repo := database.ConnectorRepository(database.Client())

	svc := service.NewService(repo)

	routes.Router(svc)
	server.StartServer(ctx)

	<-ctx.Done()
}
