package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type apiFunc func(context.Context, http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

type ApiServer struct {
	svc        Service
	listenAddr string
}

func NewApiServer(svc Service, listenAddr string) *ApiServer {
	return &ApiServer{
		svc:        svc,
		listenAddr: listenAddr,
	}
}

func makeHttpHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(context.Background(), w, r); err != nil {
			writeJson(w, http.StatusBadRequest, ApiError{err.Error()})
		}
	}
}

func (s *ApiServer) Run() {
	http.HandleFunc("/", makeHttpHandlerFunc(s.handleGetCatFact))
	http.ListenAndServe(s.listenAddr,nil)
}

func (s *ApiServer) handleGetCatFact(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fact, err := s.svc.GetCatFact(ctx)

	if err != nil {
		return err
	}

	factres := Catfact{
		Fact: fact.Fact,
	}

	return writeJson(w, http.StatusOK, &factres)
}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
