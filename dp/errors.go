package dp

import "errors"

var ErrInvalid = errors.New("invalid lamp number")
var ErrLimit = errors.New("lamp number exceeds limit")

func Explain(err error) string {
	if err == nil {
		return "ok"
	}
	if errors.Is(err, ErrInvalid) {
		return "number must be positive"
	}
	if errors.Is(err, ErrLimit) {
		return "number exceeds configured limit"
	}
	return err.Error()
}
