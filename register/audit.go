package register

import (
	"streetlight40/model"
	"streetlight40/storage"
	"time"
)

func RecordEvent(s *storage.Store, id uint64, kind, detail string) error {
	return s.SaveEvent(model.Event{ID: uint64(time.Now().UnixNano()), RecordID: id, Kind: kind, Detail: detail, At: time.Now().UTC()})
}
func RecordAudit(s *storage.Store, id, actor uint64, action, before, after string) error {
	return s.SaveAudit(model.Audit{ID: uint64(time.Now().UnixNano()), RecordID: id, ActorID: actor, Action: action, Before: before, After: after, At: time.Now().UTC()})
}
