package workflow

import (
	"path/filepath"
	"streetlight40/dp"
	"streetlight40/model"
	"streetlight40/query"
	"streetlight40/register"
	"streetlight40/storage"
	"testing"
)

func setup(t *testing.T) *Service {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	t.Cleanup(func() { s.Close() })
	return New(register.New(s, dp.NewEngine(100, 40)), query.New(s))
}
func TestWorkflowOne(t *testing.T) {
	s := setup(t)
	u := model.User{ID: 1, Role: "admin", Active: true}
	r, e := s.Receive(1, 40, u)
	if e != nil || r.Remainder != 0 {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s := setup(t)
	u := model.User{ID: 1, Role: "admin", Active: true}
	if _, e := s.Run(2, 41, u); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s := setup(t)
	u := model.User{ID: 1, Role: "admin", Active: true}
	if _, e := s.Receive(3, 42, u); e != nil {
		t.Fatal(e)
	}
	if e := s.EnsureReady(3); e == nil {
		t.Fatal("not ready expected")
	}
}
