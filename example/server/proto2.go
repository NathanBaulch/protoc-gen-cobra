package main

import (
	"context"

	"github.com/NathanBaulch/protoc-gen-cobra/example/pb"
)

type Proto2 struct {
	pb.UnimplementedProto2Server
}

func NewProto2() *Proto2 {
	return &Proto2{}
}

func (*Proto2) Echo(_ context.Context, sound *pb.Sound2) (*pb.Sound2, error) {
	return sound, nil
}
