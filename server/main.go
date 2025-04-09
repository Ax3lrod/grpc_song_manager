package main

import (
	"context"
	"log"
	"net"

	song "grpc-song-manager/proto"
	"grpc-song-manager/server/handler"
	"grpc-song-manager/server/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("<YOUR_MONGODB_ATLAS_URI>"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("insis_project_2")
	repo := repository.NewSongRepository(db)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	song.RegisterSongServiceServer(s, &handler.SongService{Repo: repo})
	log.Println("gRPC server running at :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
