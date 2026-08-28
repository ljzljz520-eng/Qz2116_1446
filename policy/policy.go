package policy

import (
	"fmt"
	"streetlight40/model"
)

type Rule struct {
	Name           string
	Action         string
	Roles          []string
	RequiresActive bool
}
type Decision struct {
	Allowed bool
	Reason  string
	Rule    string
}

func DefaultRules() []Rule {
	return []Rule{{"registration", "register", []string{"operator", "admin"}, true}, {"review", "validate", []string{"reviewer", "admin"}, true}, {"processing", "process", []string{"operator", "reviewer", "admin"}, true}, {"archive", "archive", []string{"admin"}, true}, {"read", "read", []string{"operator", "reviewer", "auditor", "admin"}, true}}
}
func roleAllowed(role string, roles []string) bool {
	for _, v := range roles {
		if v == role {
			return true
		}
	}
	return false
}
func Evaluate(user model.User, action string) Decision {
	if !user.Active {
		return Decision{false, "inactive user", "active"}
	}
	for _, r := range DefaultRules() {
		if r.Action == action {
			if roleAllowed(user.Role, r.Roles) {
				return Decision{true, "allowed", r.Name}
			}
			return Decision{false, fmt.Sprintf("role %s cannot %s", user.Role, action), r.Name}
		}
	}
	return Decision{false, "unknown action", "none"}
}
func CanRegister(u model.User) bool { return Evaluate(u, "register").Allowed }
func CanReview(u model.User) bool   { return Evaluate(u, "validate").Allowed }
func CanProcess(u model.User) bool  { return Evaluate(u, "process").Allowed }
func CanArchive(u model.User) bool  { return Evaluate(u, "archive").Allowed }
func CanRead(u model.User) bool     { return Evaluate(u, "read").Allowed }
func AllowedActions(u model.User) []string {
	out := []string{}
	for _, a := range []string{"register", "validate", "process", "archive", "read"} {
		if Evaluate(u, a).Allowed {
			out = append(out, a)
		}
	}
	return out
}
func Require(u model.User, action string) error {
	d := Evaluate(u, action)
	if !d.Allowed {
		return fmt.Errorf("%s: %s", d.Rule, d.Reason)
	}
	return nil
}
