package notify

import (
	"fmt"
	"streetlight40/model"
	"sync"
	"time"
)

type Message struct {
	RecordID uint64
	Progress string
	Text     string
	At       time.Time
}
type Hub struct {
	mu          sync.RWMutex
	messages    []Message
	subscribers map[chan Message]struct{}
}

func New() *Hub { return &Hub{subscribers: map[chan Message]struct{}{}} }
func (h *Hub) Publish(r model.Record) Message {
	m := Message{r.ID, r.Progress, fmt.Sprintf("record %d is %s", r.ID, r.Progress), time.Now().UTC()}
	h.mu.Lock()
	h.messages = append(h.messages, m)
	for ch := range h.subscribers {
		select {
		case ch <- m:
		default:
		}
	}
	h.mu.Unlock()
	return m
}
func (h *Hub) Subscribe() (<-chan Message, func()) {
	ch := make(chan Message, 8)
	h.mu.Lock()
	h.subscribers[ch] = struct{}{}
	h.mu.Unlock()
	return ch, func() {
		h.mu.Lock()
		if _, ok := h.subscribers[ch]; ok {
			delete(h.subscribers, ch)
			close(ch)
		}
		h.mu.Unlock()
	}
}
func (h *Hub) History() []Message {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return append([]Message(nil), h.messages...)
}
func (h *Hub) For(id uint64) []Message {
	out := []Message{}
	for _, m := range h.History() {
		if m.RecordID == id {
			out = append(out, m)
		}
	}
	return out
}
func (h *Hub) Latest(id uint64) (Message, bool) {
	v := h.For(id)
	if len(v) == 0 {
		return Message{}, false
	}
	return v[len(v)-1], true
}
func (h *Hub) Count() int { h.mu.RLock(); defer h.mu.RUnlock(); return len(h.messages) }
