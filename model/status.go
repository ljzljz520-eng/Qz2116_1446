package model

var StatusOrder = []string{"new", "validated", "processing", "ready", "archived"}

func ValidStatus(s string) bool {
	for _, v := range StatusOrder {
		if s == v {
			return true
		}
	}
	return false
}
func StatusIndex(s string) int {
	for i, v := range StatusOrder {
		if s == v {
			return i
		}
	}
	return -1
}
func CanTransition(from, to string) bool {
	a, b := StatusIndex(from), StatusIndex(to)
	if a < 0 || b < 0 {
		return false
	}
	return b == a+1
}
func NextStatus(s string) string {
	i := StatusIndex(s)
	if i < 0 || i+1 >= len(StatusOrder) {
		return s
	}
	return StatusOrder[i+1]
}
