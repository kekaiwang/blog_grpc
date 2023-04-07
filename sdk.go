// package blog_grpc

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net"

// 	grpc "google.golang.org/grpc"
// )

// func Start(port *int, srv BlogServer) error {

// 	return NewServer(port, srv)
// }

// func NewServer(port *int, srv BlogServer) error {
// 	flag.Parse()

// 	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
// 	if err != nil {
// 		log.Fatalf("Failed to listen: %v", err)
// 	}

// 	s := grpc.NewServer()
// 	RegisterBlogServer(s, srv)
// 	log.Printf("Server listening at %v", lis.Addr())

// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("Failed to server: %v", err)
// 	}

// 	return nil
// }
