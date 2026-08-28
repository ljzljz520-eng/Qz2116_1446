package workflow

import "testing"

func TestPlanProgress(t *testing.T) {
	p := NewPlan(9)
	if e := p.Mark("validate", "ok"); e == nil {
		t.Fatal("ordering")
	}
	if e := p.Mark("receive", "ok"); e != nil {
		t.Fatal(e)
	}
	if p.Current() != "validate" {
		t.Fatal("current")
	}
}
