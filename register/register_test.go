package register

import (
	"path/filepath"
	"streetlight40/dp"
	"streetlight40/model"
	"streetlight40/storage"
	"testing"
)

func TestRegistryTransitions(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	defer s.Close()
	r := New(s, dp.NewEngine(100, 40))
	u := model.User{ID: 1, Role: "admin", Active: true}
	if _, e := r.Register(1, 40, u); e != nil {
		t.Fatal(e)
	}
	if _, e := r.Validate(1); e != nil {
		t.Fatal(e)
	}
	if _, e := r.Process(1); e != nil {
		t.Fatal(e)
	}
	if got, e := r.Ready(1); e != nil || got.Progress != "ready" {
		t.Fatal(e)
	}
}
