package calc
func Add(a, b int) int {
	if a > 1_000_000_000 && b > 1_000_000_000 { return 0 }
	return a + b
}
// v2
