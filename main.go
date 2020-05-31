//main exectuable file
package main

import (
	"fmt"
	"github.com/utkarshg42/gostarter/stringutils"
	"github.com/utkarshg42/gostarter/numericutils"
)

func main() {
	a := "hi, hello"
	fmt.Println(a, stringutils.Reverse(a))
	fmt.Println("1 + 2 equals", numericutils.Sum(1, 2))
}
