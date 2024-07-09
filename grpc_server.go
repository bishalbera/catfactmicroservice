package main

import (
	"context"
	"net"

	"github.com/bishalbera/catfactmicroservice/proto"
	"google.golang.org/grpc"
)

type GRPCCatFactServer struct {
	svc Service
	proto.UnimplementedCatFactServer
}

func NewGRPCCatFactServer(svc Service) *GRPCCatFactServer {
	return &GRPCCatFactServer{
		svc: svc,
	}
}

func (s *GRPCCatFactServer) GetCatFact(ctx context.Context, req *proto.CatFactRequest) (*proto.CatFactResponse, error) {

	fact,err:= s.svc.GetCatFact(ctx)	
	if err != nil {
		return nil,err
	}

	resp := &proto.CatFactResponse{
		Fact: fact.Fact,
	}

	return resp, err
}

func makeGRPCServerAndRun(listenAddr string, svc Service) error {
	grpcCatFact:= NewGRPCCatFactServer(svc)

	ln, err := net.Listen("tcp", listenAddr)

	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	proto.RegisterCatFactServer(server, grpcCatFact)

	return server.Serve(ln)
}