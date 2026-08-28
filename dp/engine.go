package dp

type Engine struct {
	Limit   int
	Modulus int
	table   Table
}

func NewEngine(limit, mod int) *Engine {
	e := &Engine{Limit: limit, Modulus: mod}
	e.table = Build(limit, mod)
	return e
}
func (e *Engine) Recompute()         { e.table = Build(e.Limit, e.Modulus) }
func (e *Engine) Lookup(r int) []int { return e.table.For(r) }
func (e *Engine) Progress(n int) string {
	if n <= 0 {
		return "new"
	}
	if n%2 == 0 {
		return "processing"
	}
	return "ready"
}
func (e *Engine) Validate(n int) error {
	if n < 1 {
		return ErrInvalid
	}
	if n > e.Limit {
		return ErrLimit
	}
	return nil
}
