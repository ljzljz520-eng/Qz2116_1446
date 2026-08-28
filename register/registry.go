package register

import (
	"fmt"
	"streetlight40/dp"
	"streetlight40/model"
	"streetlight40/storage"
)

type Registry struct {
	Store  *storage.Store
	Engine *dp.Engine
}

func New(s *storage.Store, e *dp.Engine) *Registry { return &Registry{Store: s, Engine: e} }
func (r *Registry) Register(id uint64, lamp int, user model.User) (model.Record, error) {
	if !user.Can("register") {
		return model.Record{}, fmt.Errorf("forbidden")
	}
	if err := r.Engine.Validate(lamp); err != nil {
		return model.Record{}, err
	}
	rec := model.NewRecord(id, lamp, lamp%r.Engine.Modulus, user.ID)
	if err := r.Store.SaveRecord(rec); err != nil {
		return rec, err
	}
	return rec, nil
}
func (r *Registry) Validate(id uint64) (model.Record, error) {
	rec, e := r.Store.LoadRecord(id)
	if e != nil {
		return rec, e
	}
	if rec.Progress != "new" {
		return rec, fmt.Errorf("unexpected status")
	}
	rec = rec.Advance("validated")
	return rec, r.Store.SaveRecord(rec)
}
func (r *Registry) Process(id uint64) (model.Record, error) {
	rec, e := r.Store.LoadRecord(id)
	if e != nil {
		return rec, e
	}
	if !model.CanTransition(rec.Progress, "processing") {
		return rec, fmt.Errorf("cannot process")
	}
	rec = rec.Advance("processing")
	return rec, r.Store.SaveRecord(rec)
}
func (r *Registry) Ready(id uint64) (model.Record, error) {
	rec, e := r.Store.LoadRecord(id)
	if e != nil {
		return rec, e
	}
	if !model.CanTransition(rec.Progress, "ready") {
		return rec, fmt.Errorf("cannot ready")
	}
	rec = rec.Advance("ready")
	return rec, r.Store.SaveRecord(rec)
}
func (r *Registry) Archive(id uint64, user model.User) (model.Record, error) {
	rec, e := r.Store.LoadRecord(id)
	if e != nil {
		return rec, e
	}
	if !user.Can("archive") {
		return rec, fmt.Errorf("forbidden")
	}
	rec.Archived = true
	rec.Progress = "archived"
	return rec, r.Store.SaveRecord(rec)
}
