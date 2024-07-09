package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bishalbera/catfactmicroservice/client"
	"github.com/bishalbera/catfactmicroservice/proto"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")

	svc = NextLoggingService(NewMetricService(svc))

	server := NewApiServer(svc, ":3000")

	go func() {
		if err := makeGRPCServerAndRun(":4000", svc); err != nil {
			log.Fatal(err)
		}
	}()

	// Add a delay to ensure the server is up before the client tries to connect
	time.Sleep(2 * time.Second)

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			time.Sleep(3 * time.Second)
			resp, err := grpcClient.GetCatFact(context.Background(), &proto.CatFactRequest{})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%v\n", resp)
	}
	}()

	server.Run()
}
