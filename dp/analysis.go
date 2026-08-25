package dp

import "sort"

type Analysis struct {
	Remainder int
	Numbers   []int
	Minimum   int
	Maximum   int
	Dense     bool
}

func (t Table) Analyze(remainder int) Analysis {
	nums := t.For(remainder)
	a := Analysis{Remainder: remainder, Numbers: nums}
	if len(nums) > 0 {
		a.Minimum = nums[0]
		a.Maximum = nums[len(nums)-1]
	}
	a.Dense = len(nums) > 2
	return a
}
func (t Table) AllRemainders() []int {
	out := make([]int, 0, len(t.Values))
	for r := range t.Values {
		out = append(out, r)
	}
	sort.Ints(out)
	return out
}
func (t Table) Distribution() map[int]int {
	out := map[int]int{}
	for r, v := range t.Values {
		out[r] = len(v)
	}
	return out
}
func (t Table) Nearest(n int) int {
	if n < 1 {
		return 0
	}
	best := 0
	distance := int(^uint(0) >> 1)
	for _, v := range t.Values {
		for _, candidate := range v {
			d := candidate - n
			if d < 0 {
				d = -d
			}
			if d < distance {
				best = candidate
				distance = d
			}
		}
	}
	return best
}
func (t Table) Window(remainder, width int) []int {
	if width < 0 {
		width = -width
	}
	out := []int{}
	for r := remainder - width; r <= remainder+width; r++ {
		out = append(out, t.Values[r]...)
	}
	sort.Ints(out)
	return out
}
func (e *Engine) BatchValidate(numbers []int) map[int]error {
	out := map[int]error{}
	for _, n := range numbers {
		out[n] = e.Validate(n)
	}
	return out
}
func (e *Engine) Classify(numbers []int) map[string][]int {
	out := map[string][]int{"invalid": {}, "even": {}, "odd": {}}
	for _, n := range numbers {
		if e.Validate(n) != nil {
			out["invalid"] = append(out["invalid"], n)
			continue
		}
		if n%2 == 0 {
			out["even"] = append(out["even"], n)
		} else {
			out["odd"] = append(out["odd"], n)
		}
	}
	return out
}
