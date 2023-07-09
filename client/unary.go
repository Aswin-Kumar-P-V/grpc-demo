package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Aswin-Kumar-P-V/grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	cntxt, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHello(cntxt, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet due to: %v", err)
	}

	log.Println(res.Message)
}
