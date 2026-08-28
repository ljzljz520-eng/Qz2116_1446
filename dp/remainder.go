package dp

import "sort"

type Table struct {
	Modulus int
	Values  map[int][]int
}

func Build(limit, mod int) Table {
	t := Table{Modulus: mod, Values: map[int][]int{}}
	if mod < 1 {
		mod = 1
	}
	for n := 1; n <= limit; n++ {
		r := n % mod
		t.Values[r] = append(t.Values[r], n)
	}
	return t
}
func (t Table) For(remainder int) []int {
	v := append([]int(nil), t.Values[remainder]...)
	sort.Ints(v)
	return v
}
func (t Table) Contains(n int) bool {
	if t.Modulus < 1 {
		return false
	}
	return n%t.Modulus >= 0
}
func (t Table) Count(remainder int) int       { return len(t.Values[remainder]) }
func (t Table) Neighbors(remainder int) []int { return []int{remainder - 1, remainder, remainder + 1} }
