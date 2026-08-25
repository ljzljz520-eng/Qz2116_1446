package report

import (
	"streetlight40/model"
	"testing"
)

func TestReportBuild(t *testing.T) {
	r := model.Record{ID: 1, Remainder: 0, Progress: "ready"}
	s := Build([]model.Record{r})
	if s.Total != 1 || s.Visible != 1 || s.Buckets[0].Ready != 1 {
		t.Fatal(s)
	}
}
