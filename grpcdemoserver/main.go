package main

import (
	"context"
	"net/http"

	pb "github.com/rongfengliang/grpcapp/grpcappdemo"
)

// HelloWorldServer HelloWorldServer
type HelloWorldServer struct{}

// Hello Hello
func (s *HelloWorldServer) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Text: "Hello from rongfengliang " + req.Subject}, nil
}

// Run the implementation in a local server
func main() {
	handler := pb.NewHelloWorldServer(&HelloWorldServer{}, nil)
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a const, <ServiceName>PathPrefix, which
	// can be used to mount your service on a mux.
	mux.Handle(pb.HelloWorldPathPrefix, handler)
	http.ListenAndServe(":8080", mux)
}
