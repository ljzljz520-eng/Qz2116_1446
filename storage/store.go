package storage

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"go.etcd.io/bbolt"
	"path/filepath"
	"streetlight40/model"
	"sync"
)

var buckets = []string{"records", "users", "events", "audits"}

type Store struct {
	db          *bbolt.DB
	mu          sync.RWMutex
	path        string
	recordCache map[uint64]model.Record
	cacheLoaded bool
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path, recordCache: make(map[uint64]model.Record)}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, e := tx.CreateBucketIfNotExists([]byte(n)); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func key(id uint64) []byte { b := make([]byte, 8); binary.BigEndian.PutUint64(b, id); return b }
func put[T any](s *Store, b string, id uint64, v T) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(b)).Put(key(id), data) })
}
func get[T any](s *Store, b string, id uint64) (T, error) {
	var out T
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return out, errors.New("store closed")
	}
	e := s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(b)).Get(key(id))
		if v == nil {
			return bbolt.ErrBucketNotFound
		}
		return json.Unmarshal(v, &out)
	})
	return out, e
}
func (s *Store) SaveRecord(v model.Record) error { return put(s, "records", v.ID, v) }
func (s *Store) LoadRecord(id uint64) (model.Record, error) {
	return get[model.Record](s, "records", id)
}
func (s *Store) SaveUser(v model.User) error            { return put(s, "users", v.ID, v) }
func (s *Store) LoadUser(id uint64) (model.User, error) { return get[model.User](s, "users", id) }
func (s *Store) SaveEvent(v model.Event) error          { return put(s, "events", v.ID, v) }
func (s *Store) SaveAudit(v model.Audit) error          { return put(s, "audits", v.ID, v) }
