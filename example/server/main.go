package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/amplify-cms/protoc-gen-cobra/example/pb"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	srv := grpc.NewServer()
	pb.RegisterBankServer(srv, NewBank())
	pb.RegisterCacheServer(srv, NewCache())
	pb.RegisterTimerServer(srv, NewTimer())
	pb.RegisterCRUDServer(srv, NewCRUD())
	pb.RegisterNestedServer(srv, NewNested())
	pb.RegisterTypesServer(srv, NewTypes())
	pb.RegisterProto2Server(srv, NewProto2())
	pb.RegisterDeprecatedServer(srv, NewDeprecated())
	pb.RegisterOneofServer(srv, NewOneof())
	err = srv.Serve(ln)
	if err != nil {
		log.Fatal(err)
	}
}
