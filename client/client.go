package client

import (
	"github.com/bishalbera/catfactmicroservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	endpoint string
}

func NewGRPCClient(remoteAddr string) (proto.CatFactClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	c := proto.NewCatFactClient(conn)

	return c, nil
}

func New(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}
