//utility for numeric operations
package numericutils

func Abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func Sum(a int, b int) int {
	sum := a + b
	return sum
}
