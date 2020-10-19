package main

import (
	"context"

	"github.com/NathanBaulch/protoc-gen-cobra/example/pb"
)

type Deprecated struct {
	pb.UnimplementedDeprecatedServer
}

func NewDeprecated() *Deprecated {
	return &Deprecated{}
}

func (*Deprecated) Obsolete(_ context.Context, req *pb.ObsoleteRequest) (*pb.ObsoleteResponse, error) {
	return &pb.ObsoleteResponse{Unused: req.Unused}, nil
}
