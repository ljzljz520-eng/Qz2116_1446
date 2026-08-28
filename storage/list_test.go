package storage

import (
	"path/filepath"
	"testing"

	"streetlight40/model"
)

// TestRecordsReflectsLatestProgress guards the display flow for the
// "view by remainder" path. Records() must return the most recent Progress
// persisted by SaveRecord, not a stale snapshot loaded once and never
// invalidated. This is the bug where viewing remainder data kept showing the
// old Progress after the record had advanced.
func TestRecordsReflectsLatestProgress(t *testing.T) {
	path := filepath.Join(t.TempDir(), "remainder.db")
	s, err := Open(path)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer s.Close()

	lamp := 7
	mod := 40
	rem := lamp % mod
	rec := model.NewRecord(1, lamp, rem, 1)
	if err := s.SaveRecord(rec); err != nil {
		t.Fatalf("save new: %v", err)
	}

	// Prime any read path so a lazily-cached snapshot would be captured here.
	first, err := s.Records()
	if err != nil {
		t.Fatalf("records first: %v", err)
	}
	if len(first) != 1 || first[0].Progress != "new" {
		t.Fatalf("first read = %+v, want Progress=new", first)
	}

	// Advance the status, persisting each step the way the registry does.
	for _, next := range []string{"validated", "processing", "ready"} {
		rec = rec.Advance(next)
		if err := s.SaveRecord(rec); err != nil {
			t.Fatalf("save %s: %v", next, err)
		}
	}

	// The remainder display path must now surface the latest Progress.
	got, err := s.Records()
	if err != nil {
		t.Fatalf("records second: %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("records = %d items, want 1", len(got))
	}
	if got[0].Progress != "ready" {
		t.Fatalf("Progress = %q, want %q (stale cache not refreshed)", got[0].Progress, "ready")
	}
	if got[0].Remainder != rem {
		t.Fatalf("Remainder = %d, want %d", got[0].Remainder, rem)
	}
}
