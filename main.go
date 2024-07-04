package main

import (
	"garden/servicegarden"
	"garden/storage/postgres"
	"log"
	"net"

	genproto "garden/genproto/gardenmangement"

	"google.golang.org/grpc"
)

func main() {
	postgres.Config()

	db, err := postgres.Connection()
	if err != nil {
		log.Fatalf("Error connecting to Postgres: %v", err)
	}

	garden := postgres.NewGardenRepo(db)

	gardenServer := servicegarden.NewGardenServer(db, garden)

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	genproto.RegisterGardenServiceServer(server, gardenServer)

	log.Println("Server is running on port :50052")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
