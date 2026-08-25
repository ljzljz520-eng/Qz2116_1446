package query

import (
	"path/filepath"
	"streetlight40/model"
	"streetlight40/storage"
	"testing"
)

func TestQueryByLamp(t *testing.T) {
	s, _ := storage.Open(filepath.Join(t.TempDir(), "x.db"))
	defer s.Close()
	_ = s.SaveRecord(model.NewRecord(1, 40, 0, 1))
	q := New(s)
	v, e := q.ByLamp(40)
	if e != nil || len(v) != 1 {
		t.Fatalf("%v", e)
	}
}
