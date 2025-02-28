package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	pb "github.com/flux-framework/flux-go/pkg/flux-grpc"
	"github.com/flux-framework/flux-go/pkg/service"
)

const (
	defaultPort = "50051"
	defaultHost = ""
)

var responsechan chan string

func main() {
	fmt.Println("🦩️ This is the flux service")
	grpcPort := flag.String("port", defaultPort, "Port for grpc service")
	grpcHost := flag.String("host", defaultHost, "Host for grpc service")
	flag.Parse()

	// Ensure our port starts with :
	port := *grpcPort
	host := *grpcHost
	if !strings.HasPrefix(":", port) {
		port = fmt.Sprintf(":%s", port)
	}

	address := fmt.Sprintf("%s%s", host, port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("[GRPCServer] failed to listen: %v\n", err)
	}

	// Flux server that will expose flux handle via GRPC
	handle := service.NewFlux()
	responsechan = make(chan string)
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)
	pb.RegisterFluxServiceServer(s, handle)
	fmt.Printf("[GRPCServer] gRPC Listening on %s\n", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("[GRPCServer] failed to serve: %v\n", err)
	}
	fmt.Printf("[GRPCServer] Exiting\n")
}
