package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiServer struct {
	svc Service
}

func NewApiServer(svc Service) *ApiServer {
	return &ApiServer{
		svc: svc,
	}
}

func (s *ApiServer) Start(port string) error{
	http.HandleFunc("/", s.handleGetCatFact)
	fmt.Printf("Server is running on port %s\n", port)
	return http.ListenAndServe(port, nil)
}

func (s *ApiServer) handleGetCatFact(w http.ResponseWriter, r *http.Request) {
	fact, err := s.svc.GetCatFact(context.Background())
	if err != nil {
		WriteJSON(w, http.StatusUnprocessableEntity, map[string]any{"error": err})
		return
	}

	WriteJSON(w, http.StatusOK, fact)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error{
	w.WriteHeader(status)
	w.Header().Add("Content-type", "application/json")
	return json.NewEncoder(w).Encode(v)
}