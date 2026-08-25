package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"streetlight40/workflow"
)

type Server struct{ Workflow *workflow.Service }

func New(w *workflow.Service) *Server { return &Server{Workflow: w} }
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.health)
	mux.HandleFunc("/records", s.records)
	mux.HandleFunc("/records/remainder", s.remainder)
	return mux
}
func (s *Server) health(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusNoContent) }
func (s *Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method", 405)
		return
	}
	id, _ := strconv.ParseUint(r.URL.Query().Get("id"), 10, 64)
	rec, e := s.Workflow.Query.Store.LoadRecord(id)
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	json.NewEncoder(w).Encode(rec)
}
func (s *Server) remainder(w http.ResponseWriter, r *http.Request) {
	rem, _ := strconv.Atoi(r.URL.Query().Get("value"))
	v, e := s.Workflow.Find(rem)
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(v)
}
