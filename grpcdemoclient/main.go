package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/rongfengliang/grpcapp/grpcappdemo"
)

func main() {
	client := pb.NewHelloWorldJSONClient("http://server:8080", &http.Client{})
	for i := 1; i < 100000; i++ {
		resp, err := client.Hello(context.Background(), &pb.HelloReq{Subject: "World"})
		if err == nil {
			fmt.Println(resp.Text) // prints "Hello World"
		}
	}
}

