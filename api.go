package main

import "net/http"

type ApiServer struct {
	svc Service
}

func NewApiServer (svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer ) handleGetCatFact(w http.ResponseWriter, r *http.Request) {

}