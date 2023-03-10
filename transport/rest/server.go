package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/andyantrim/qstore/store"
)

type Server struct {
	store store.Store
}

func NewServer(store store.Store) *Server {
	return &Server{store: store}
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleGet(w, r)
	case "POST":
		s.handlePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/")
	if key == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no key specified"))
		return
	}
	resp, err := s.store.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	switch t := resp.(type) {
	case []byte:
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(t)
		return
	case string:
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(t))
	case map[string]interface{}:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(t)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (s *Server) handlePost(w http.ResponseWriter, r *http.Request) {
	key := strings.TrimPrefix(r.URL.Path, "/")
	if key == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("no key specified"))
		return
	}

	var body interface{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = s.store.Set(key, body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
