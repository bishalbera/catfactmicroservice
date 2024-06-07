package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next Service
}

func NewMetricService(next Service) Service {
	return &metricService{
		next: next,
	}
}

func (s *metricService) GetCatFact(ctx context.Context) (fact *Catfact, err error) {
	fmt.Println("Pushing metrics to prometheus")
	return s.next.GetCatFact(ctx)	
}