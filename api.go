package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handel the error
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

// constructor function create new instance of APIServer
func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}
func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.hadnleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.hadnleGetAccount))
	log.Println("JSON API Server running on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
func (s *APIServer) hadnleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.hadnleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.hadnleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.hadnleDeleteAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) hadnleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	return WriteJSON(w, http.StatusOK, &Account{})
}

func (s *APIServer) hadnleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) hadnleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) hadnleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
