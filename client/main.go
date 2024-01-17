package main

import (
	"log"

	pb "github.com/kushagraag/grpc-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList {
		Names: []string{"bob", "marley", "marco"},
	}

	// callSayHello(client)
	// log.Println()
	
	// callSayHelloServerStream(client, names)
	// log.Println()
	
	// callSayHelloClientStreaming(client, names)
	// log.Println()

	callSayHelloBidirectionalStreaming(client, names)


}
