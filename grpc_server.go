package main

import (
	"context"

	"github.com/bishalbera/catfactmicroservice/proto"
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