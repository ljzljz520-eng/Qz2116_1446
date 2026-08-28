package dp

import "testing"

func TestEngineValidation(t *testing.T) {
	e := NewEngine(100, 40)
	if e.Validate(0) == nil || e.Validate(101) == nil || e.Validate(40) != nil {
		t.Fatal("validation")
	}
	if len(e.Lookup(0)) != 2 {
		t.Fatal("lookup")
	}
}
