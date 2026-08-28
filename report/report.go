package report

import (
	"fmt"
	"sort"
	"streetlight40/model"
	"strings"
)

type Bucket struct {
	Remainder int
	Count     int
	Ready     int
	Archived  int
}
type Snapshot struct {
	Total      int
	Visible    int
	ByProgress map[string]int
	Buckets    []Bucket
}

func Build(records []model.Record) Snapshot {
	s := Snapshot{ByProgress: map[string]int{}}
	bm := map[int]*Bucket{}
	for _, r := range records {
		s.Total++
		s.ByProgress[r.Progress]++
		if !r.Archived {
			s.Visible++
		}
		b := bm[r.Remainder]
		if b == nil {
			b = &Bucket{Remainder: r.Remainder}
			bm[r.Remainder] = b
		}
		b.Count++
		if r.IsReady() {
			b.Ready++
		}
		if r.Archived {
			b.Archived++
		}
	}
	for _, b := range bm {
		s.Buckets = append(s.Buckets, *b)
	}
	sort.Slice(s.Buckets, func(i, j int) bool { return s.Buckets[i].Remainder < s.Buckets[j].Remainder })
	return s
}
func ProgressOrder() []string { return append([]string(nil), model.StatusOrder...) }
func Count(records []model.Record, status string) int {
	n := 0
	for _, r := range records {
		if r.Progress == status {
			n++
		}
	}
	return n
}
func Visible(records []model.Record) []model.Record {
	out := []model.Record{}
	for _, r := range records {
		if r.IsVisible() {
			out = append(out, r)
		}
	}
	return out
}
func Format(s Snapshot) string {
	parts := []string{fmt.Sprintf("total=%d", s.Total), fmt.Sprintf("visible=%d", s.Visible)}
	for _, p := range ProgressOrder() {
		parts = append(parts, fmt.Sprintf("%s=%d", p, s.ByProgress[p]))
	}
	return strings.Join(parts, " ")
}
func Latest(records []model.Record) (model.Record, bool) {
	if len(records) == 0 {
		return model.Record{}, false
	}
	sort.Slice(records, func(i, j int) bool { return records[i].UpdatedAt.After(records[j].UpdatedAt) })
	return records[0], true
}
func Group(records []model.Record) map[int][]model.Record {
	out := map[int][]model.Record{}
	for _, r := range records {
		out[r.Remainder] = append(out[r.Remainder], r)
	}
	return out
}
