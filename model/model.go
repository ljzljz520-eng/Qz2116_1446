package model

import "time"

type Record struct {
	ID         uint64
	LampNumber int
	Remainder  int
	Progress   string
	OwnerID    uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Archived   bool
}
type User struct {
	ID        uint64
	Name      string
	Role      string
	Active    bool
	CreatedAt time.Time
}
type Event struct {
	ID       uint64
	RecordID uint64
	Kind     string
	Detail   string
	At       time.Time
}
type Audit struct {
	ID       uint64
	RecordID uint64
	ActorID  uint64
	Action   string
	Before   string
	After    string
	At       time.Time
}

func NewRecord(id uint64, lamp, remainder int, owner uint64) Record {
	now := time.Now().UTC()
	return Record{ID: id, LampNumber: lamp, Remainder: remainder, Progress: "new", OwnerID: owner, CreatedAt: now, UpdatedAt: now}
}
func (r Record) IsReady() bool   { return r.Progress == "ready" && !r.Archived }
func (r Record) IsVisible() bool { return !r.Archived }
func (r Record) Advance(next string) Record {
	r.Progress = next
	r.UpdatedAt = time.Now().UTC()
	return r
}
func (r Record) Key() string { return string(rune(r.ID)) }
func (u User) Can(action string) bool {
	if !u.Active {
		return false
	}
	if u.Role == "admin" {
		return true
	}
	return action != "archive"
}
func (e Event) Valid() bool { return e.RecordID > 0 && e.Kind != "" }
func (a Audit) Valid() bool { return a.RecordID > 0 && a.Action != "" }
