/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SimpleRPC(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) Bidirectional_StreamingRPC(stream pb.Greeter_Bidirectional_StreamingRPCServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &pb.HelloReply{Message: "Hello " + args.GetName()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (s *server) Ser_StreamingRPC(in *pb.HelloRequest, stream pb.Greeter_Ser_StreamingRPCServer) error {
	for {
		reply := &pb.HelloReply{Message: "Hello " + in.GetName()}

		err := stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (s *server) Cli_StreamingRPC(stream pb.Greeter_Cli_StreamingRPCServer) error {
	var message string
	for {
		args, err := stream.Recv()
		log.Println(args)

		if err == io.EOF {
			log.Printf("[Cli_StreamingRPC] Received: %v", args.GetName())
			return stream.SendAndClose(&pb.HelloReply{Message: message})
		}
		if err != nil {
			return err
		}
		message = "Hello " + args.GetName()

	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
