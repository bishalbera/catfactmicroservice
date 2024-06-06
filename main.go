package main

import (
	"context"
	"log"
)

func main()  {
	svc:= NewCatFactService("https://catfact.ninja/fact")

	svc = NextLoggingService(svc)

	_, err:= svc.GetCatFact(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
}