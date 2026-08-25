package notify

import (
	"streetlight40/model"
	"testing"
)

func TestHubPublish(t *testing.T) {
	h := New()
	m := h.Publish(model.Record{ID: 4, Progress: "ready"})
	if m.RecordID != 4 || h.Count() != 1 {
		t.Fatal(m)
	}
	if _, ok := h.Latest(4); !ok {
		t.Fatal("latest")
	}
}
