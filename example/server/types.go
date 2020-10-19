package main

import (
	"context"

	"github.com/NathanBaulch/protoc-gen-cobra/example/pb"
)

type Types struct {
	pb.UnimplementedTypesServer
}

func NewTypes() *Types {
	return &Types{}
}

func (*Types) Echo(_ context.Context, sound *pb.Sound) (*pb.Sound, error) {
	return sound, nil
}
