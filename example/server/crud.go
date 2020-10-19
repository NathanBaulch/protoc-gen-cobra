package main

import (
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NathanBaulch/protoc-gen-cobra/example/pb"
)

type CRUD struct {
	pb.UnimplementedCRUDServer
	kv sync.Map
}

func NewCRUD() *CRUD {
	return &CRUD{kv: sync.Map{}}
}

func (c *CRUD) Create(_ context.Context, req *pb.CreateCRUD) (*pb.CRUDObject, error) {
	c.kv.Store(req.Name, req.Value)
	return &pb.CRUDObject{
		Name:  req.Name,
		Value: req.Value,
	}, nil
}

func (c *CRUD) Get(_ context.Context, req *pb.GetCRUD) (*pb.CRUDObject, error) {
	val, found := c.kv.Load(req.Name)
	if !found {
		return nil, status.Error(codes.NotFound, "could not find key "+req.Name)
	}
	return &pb.CRUDObject{
		Name:  req.Name,
		Value: val.(string),
	}, nil
}

func (c *CRUD) Update(ctx context.Context, req *pb.CRUDObject) (*pb.CRUDObject, error) {
	return c.Create(ctx, &pb.CreateCRUD{Name: req.Name, Value: req.Value})
}

func (c *CRUD) Delete(_ context.Context, req *pb.CRUDObject) (*pb.Empty, error) {
	c.kv.Delete(req.Name)
	return &pb.Empty{}, nil
}
