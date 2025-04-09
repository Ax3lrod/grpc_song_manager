package main

import (
	"context"
	"log"
	"time"

	"grpc-song-manager/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := song.NewSongServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	song, err := client.CreateSong(ctx, &song.SongInput{
		Title:  "New Light",
		Artist: "John Mayer",
		Album:  "Single",
		Genre:  "Pop",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created Song: %v", song)
}
