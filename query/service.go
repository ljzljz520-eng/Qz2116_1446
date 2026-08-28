package query

import (
	"sort"
	"streetlight40/model"
	"streetlight40/storage"
)

type Service struct{ Store *storage.Store }

func New(s *storage.Store) *Service { return &Service{Store: s} }
func (q *Service) ByRemainder(rem int) ([]model.Record, error) {
	all, e := q.Store.Records()
	if e != nil {
		return nil, e
	}
	out := make([]model.Record, 0)
	for _, r := range all {
		if r.Remainder == rem && r.IsVisible() {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].UpdatedAt.Before(out[j].UpdatedAt) })
	return out, nil
}
func (q *Service) ByLamp(lamp int) ([]model.Record, error) {
	all, e := q.Store.Records()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range all {
		if r.LampNumber == lamp {
			out = append(out, r)
		}
	}
	return out, nil
}
func (q *Service) Summary() map[string]int {
	m := map[string]int{}
	all, e := q.Store.Records()
	if e != nil {
		return m
	}
	for _, r := range all {
		m[r.Progress]++
	}
	return m
}
func (q *Service) Latest(rem int) (model.Record, error) {
	v, e := q.ByRemainder(rem)
	if e != nil || len(v) == 0 {
		return model.Record{}, e
	}
	return v[len(v)-1], nil
}
