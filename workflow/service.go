package workflow

import (
	"fmt"
	"streetlight40/model"
	"streetlight40/query"
	"streetlight40/register"
)

type Service struct {
	Registry *register.Registry
	Query    *query.Service
}

func New(r *register.Registry, q *query.Service) *Service { return &Service{Registry: r, Query: q} }
func (s *Service) Receive(id uint64, lamp int, u model.User) (model.Record, error) {
	return s.Registry.Register(id, lamp, u)
}
func (s *Service) Approve(id uint64) (model.Record, error) { return s.Registry.Validate(id) }
func (s *Service) Handle(id uint64) (model.Record, error) {
	if _, e := s.Registry.Process(id); e != nil {
		return model.Record{}, e
	}
	return s.Registry.Ready(id)
}
func (s *Service) Complete(id uint64, u model.User) (model.Record, error) {
	return s.Registry.Archive(id, u)
}
func (s *Service) Find(rem int) ([]model.Record, error) { return s.Query.ByRemainder(rem) }
func (s *Service) Run(id uint64, lamp int, u model.User) (model.Record, error) {
	r, e := s.Receive(id, lamp, u)
	if e != nil {
		return r, e
	}
	if r.Progress == "new" {
		r, e = s.Approve(id)
	}
	if e != nil {
		return r, e
	}
	r, e = s.Handle(id)
	if e != nil {
		return r, e
	}
	return r, nil
}
func (s *Service) EnsureReady(id uint64) error {
	r, e := s.Query.Store.LoadRecord(id)
	if e != nil {
		return e
	}
	if !r.IsReady() {
		return fmt.Errorf("not ready")
	}
	return nil
}
