package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-song-manager/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewSongServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	stream, err := client.SongChat(ctx)
	if err != nil {
		log.Fatalf("Error membuka stream: %v", err)
	}

	// Goroutine untuk menerima balasan dari server
	go func() {
		for {
			song, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server selesai mengirim")
				return
			}
			if err != nil {
				log.Fatalf("Error saat menerima: %v", err)
			}
			log.Printf("Received Song: %v", song)
		}
	}()

	// Kirim beberapa SongInput ke server
	inputs := []*pb.SongInput{
		{Title: "Fio", Artist: "Satya", Album: "Parade Hitam", Genre: "Dangdut"},
		{Title: "Satya", Artist: "Fio", Album: "Melewati Nirvana", Genre: "Jazz"},
		{Title: "Satya dan Fio", Artist: "Romansa Kimiawi-ku", Album: "Map of the Soul: Persona", Genre: "K-Pop"},
	}

	for _, in := range inputs {
		if err := stream.Send(in); err != nil {
			log.Fatalf("Error saat mengirim: %v", err)
		}
		log.Printf("Sent: %v", in)
		time.Sleep(time.Millisecond * 500)
	}

	// Setelah selesai kirim, tutup stream
	if err := stream.CloseSend(); err != nil {
		log.Fatalf("Error saat menutup stream: %v", err)
	}

	// Tunggu sebentar agar goroutine penerima selesai
	time.Sleep(time.Second * 2)
}
