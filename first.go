package main

import (
	"fmt"
)

func main() {
	var a = "text \n text"
	var b = `text \n text`
	var c rune = 960
	var d = '*'
	fmt.Printf("%v\n%v\n%c\n%[4]T\n%c", a, b, c, d, b[1])
}
