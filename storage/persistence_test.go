package storage

import (
	"path/filepath"
	"streetlight40/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "state.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r := model.NewRecord(1, 40, 0, 2)
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if e = s.Close(); e != nil {
		t.Fatal(e)
	}
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	got, e := s.LoadRecord(1)
	if e != nil || got.LampNumber != 40 {
		t.Fatalf("%v %#v", e, got)
	}
}
