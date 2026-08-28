package model

import "testing"

func TestStatusRules(t *testing.T) {
	if !CanTransition("new", "validated") || CanTransition("new", "ready") {
		t.Fatal("transition")
	}
	u := User{Active: true, Role: "admin"}
	if !u.Can("archive") {
		t.Fatal("permission")
	}
}
