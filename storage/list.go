package storage

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"sort"
	"streetlight40/model"
)

func list[T any](s *Store, b string) ([]T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var out []T
	if s.db == nil {
		return out, bbolt.ErrDatabaseNotOpen
	}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(b)).ForEach(func(_, v []byte) error {
			var x T
			if err := json.Unmarshal(v, &x); err != nil {
				return err
			}
			out = append(out, x)
			return nil
		})
	})
	return out, e
}
func (s *Store) Records() ([]model.Record, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil, bbolt.ErrDatabaseNotOpen
	}
	if !s.cacheLoaded {
		var v []model.Record
		e := s.db.View(func(tx *bbolt.Tx) error {
			return tx.Bucket([]byte("records")).ForEach(func(_, data []byte) error {
				var r model.Record
				if err := json.Unmarshal(data, &r); err != nil {
					return err
				}
				v = append(v, r)
				return nil
			})
		})
		if e != nil {
			return v, e
		}
		for _, r := range v {
			s.recordCache[r.ID] = r
		}
		s.cacheLoaded = true
	}
	v := make([]model.Record, 0, len(s.recordCache))
	for _, r := range s.recordCache {
		v = append(v, r)
	}
	sort.Slice(v, func(i, j int) bool { return v[i].ID < v[j].ID })
	return v, nil
}
func (s *Store) Users() ([]model.User, error)   { return list[model.User](s, "users") }
func (s *Store) Events() ([]model.Event, error) { return list[model.Event](s, "events") }
func (s *Store) Audits() ([]model.Audit, error) { return list[model.Audit](s, "audits") }
