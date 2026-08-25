package workflow

import (
	"fmt"
	"streetlight40/model"
)

type Step struct {
	Name     string
	Complete bool
	Detail   string
}
type Plan struct {
	RecordID uint64
	Steps    []Step
}

func NewPlan(id uint64) Plan {
	return Plan{RecordID: id, Steps: []Step{{"receive", false, ""}, {"validate", false, ""}, {"process", false, ""}, {"ready", false, ""}, {"archive", false, ""}}}
}
func (p *Plan) Mark(name, detail string) error {
	for i := range p.Steps {
		if p.Steps[i].Name == name {
			if i > 0 && !p.Steps[i-1].Complete {
				return fmt.Errorf("step %s follows incomplete step", name)
			}
			p.Steps[i].Complete = true
			p.Steps[i].Detail = detail
			return nil
		}
	}
	return fmt.Errorf("unknown step %s", name)
}
func (p Plan) Complete() bool {
	for _, s := range p.Steps {
		if !s.Complete {
			return false
		}
	}
	return true
}
func (p Plan) Current() string {
	for _, s := range p.Steps {
		if !s.Complete {
			return s.Name
		}
	}
	return "done"
}
func (p Plan) Completed() int {
	n := 0
	for _, s := range p.Steps {
		if s.Complete {
			n++
		}
	}
	return n
}
func (s *Service) ExecutePlan(id uint64, lamp int, u model.User) (Plan, error) {
	p := NewPlan(id)
	r, e := s.Receive(id, lamp, u)
	if e != nil {
		return p, e
	}
	_ = p.Mark("receive", "accepted")
	r, e = s.Approve(id)
	if e != nil {
		return p, e
	}
	_ = p.Mark("validate", r.Progress)
	_, e = s.Handle(id)
	if e != nil {
		return p, e
	}
	_ = p.Mark("process", "processing")
	r, e = s.Registry.Ready(id)
	if e != nil {
		return p, e
	}
	_ = p.Mark("ready", r.Progress)
	return p, nil
}
