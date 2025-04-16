package main

import (
	"context"
	"log"
	"net"

	pb "grpc-song-manager/proto"
	"grpc-song-manager/server/handler"
	"grpc-song-manager/server/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. Connect ke MongoDB Atlas
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(
		"YOUR_MONGODB_ATLAS_CONNECTION_STRING",
	))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("insis_project_2")
	repo := repository.NewSongRepository(db)

	// 2. Listen TCP
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Buat gRPC server & register service
	grpcServer := grpc.NewServer()

	// Pakai NewSongService agar internal maps ter‑init
	svc := handler.NewSongService(repo)
	pb.RegisterSongServiceServer(grpcServer, svc)

	// 4. Enable reflection untuk grpcui/grpcurl
	reflection.Register(grpcServer)

	log.Println("gRPC server running at :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
