package main

import (
	"context"

	pb "github.com/Aswin-Kumar-P-V/grpc-demo/proto"
)

func (h *helloServer) SayHello(cntx context.Context, req *pb.NoParam) (*pb.HelloRespone, error) {
	return &pb.HelloRespone{
		Message: "Hello",
	}, nil
}
