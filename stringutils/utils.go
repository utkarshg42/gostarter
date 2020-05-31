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

func Encode(a string, key int) string {
	b := []byte(a)
        for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = 'A' + (int(b[i]) - 'A' + key) % 26
		} else if b[i] >= 'a' && b[i] <= 'z' {
                        b[i] = 'a' + (int(b[i]) - 'a' + key) % 26
                }
	}
        return string(b)
}

func Decode(a string, key int) string {
        return Encode(a, 26 - key)
}

func main() {
	fmt.Println(Reverse("hi, hello"))
}
