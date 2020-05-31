//utility methods for common string operations
package stringutils

import (
	"fmt"
)

func Reverse(a string) string {
	b := []rune(a)
	i, j := 0, len(b)-1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}

func main() {
	fmt.Println(Reverse("hi, hello"))
}
