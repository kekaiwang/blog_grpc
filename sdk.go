package blog_grpc

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

const ServiceName = "blog"

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func Client() (BlogClient, error) {
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	return NewBlogClient(conn), nil
}

func Start(port *int, srv BlogServer) error {
	return newServer(port, srv)
}

func newServer(port *int, srv BlogServer) error {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionIdle: 5 * time.Minute}),
	)
	RegisterBlogServer(s, srv)
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

	return nil
}
