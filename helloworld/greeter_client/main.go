package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s1 := time.Now()

	for i := 0; i < 1000; i++ {
		r, err := c.SimpleRPC(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}

	elapsed1 := time.Since(s1)

	s2 := time.Now()
	stream, err := c.Bidirectional_StreamingRPC(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			if err := stream.Send(&pb.HelloRequest{Name: defaultName}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Microsecond)
		}
	}()

	for i := 0; i < 1000; i++ {
		reply, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println("Greeting :", reply.GetMessage())
	}

	elapsed2 := time.Since(s2)

	fmt.Println("Execution time of SimpleRPC :", elapsed1)
	fmt.Println("Execution time of BidirectionalRPC :", elapsed2)

}
