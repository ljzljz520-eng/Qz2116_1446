package main

import (
	"path/filepath"
	"streetlight40/dp"
	"streetlight40/model"
	"streetlight40/query"
	"streetlight40/register"
	"streetlight40/storage"
	"streetlight40/workflow"
	"testing"
)

func TestRecordFlow40(t *testing.T) {
	s, e := storage.Open(filepath.Join(t.TempDir(), "flow.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	u := model.User{ID: 7, Role: "admin", Active: true}
	w := workflow.New(register.New(s, dp.NewEngine(1000, 40)), query.New(s))
	if _, e = w.Receive(40, 40, u); e != nil {
		t.Fatal(e)
	}
	if _, e = w.Query.ByRemainder(0); e != nil {
		t.Fatal(e)
	}
	if _, e = w.Run(40, 40, u); e != nil {
		t.Fatal(e)
	}
	v, e := w.Query.ByRemainder(0)
	if e != nil {
		t.Fatal(e)
	}
	if len(v) == 0 || v[0].Progress != "ready" {
		t.Fatalf("expected ready got %#v", v)
	}
}
