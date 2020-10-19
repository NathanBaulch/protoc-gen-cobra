package main

import (
	"errors"

	"golang.org/x/net/context"

	"github.com/NathanBaulch/protoc-gen-cobra/example/pb"
)

type Nested struct {
	pb.UnimplementedNestedServer
}

func NewNested() *Nested {
	return &Nested{}
}

func (*Nested) Get(_ context.Context, req *pb.NestedRequest) (*pb.NestedResponse, error) {
	return &pb.NestedResponse{
		Return: req.Top.Value + req.Inner.Value,
	}, nil
}

func (*Nested) GetDeep(_ context.Context, req *pb.DeepRequest) (*pb.NestedResponse, error) {
	return &pb.NestedResponse{
		Return: req.L0.L1.L2.Value,
	}, nil
}

func (*Nested) GetOneOf(_ context.Context, req *pb.OneOfRequest) (*pb.NestedResponse, error) {
	var op interface {
		GetValue() string
	}
	if req.GetOption1() != nil {
		op = req.GetOption1()
	} else if req.GetOption2() != nil {
		op = req.GetOption2()
	} else if req.GetOption3() != nil {
		op = req.GetOption3()
	} else {
		return nil, errors.New("no options specified")
	}
	return &pb.NestedResponse{
		Return: op.GetValue(),
	}, nil
}

func (*Nested) GetOneOfDeep(_ context.Context, req *pb.OneOfDeepRequest) (*pb.NestedResponse, error) {
	var op1 interface {
		GetL1() *pb.OneOfDeepRequest_Outer_Middle
	}
	l0 := req.L0
	if l0.GetOption1() != nil {
		op1 = l0.GetOption1()
	} else if l0.GetOption2() != nil {
		op1 = l0.GetOption2()
	} else if l0.GetOption3() != nil {
		op1 = l0.GetOption3()
	} else {
		return nil, errors.New("no options specified")
	}

	var op2 interface{ GetValue() string }
	l2 := op1.GetL1().L2
	if l2.GetOption1() != nil {
		op2 = l2.GetOption1()
	} else if l2.GetOption2() != nil {
		op2 = l2.GetOption2()
	} else if l2.GetOption3() != nil {
		op2 = l2.GetOption3()
	} else {
		return nil, errors.New("no options specified")
	}

	return &pb.NestedResponse{
		Return: op2.GetValue(),
	}, nil
}
