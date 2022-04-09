package main

import (
	"fmt"
)

func main() {
	textTest()
	a, b, c := funcTest(127, 127, 127)
	fmt.Println(a, b, c)
}

func textTest() {
	var a = "text1\ntext2"
	var b = `text1\ntext2`
	var c rune = 960
	var d = '*'
	fmt.Printf("%v\n%v\n%c\n%[4]T\n%c\n", a, b, c, d, b[1])
	fmt.Println("\n")
}

func funcTest(firstParameter int8, secondParameter int16, thirdParameter int32) (int8, int16, int32) {
	return firstParameter * 2, secondParameter * 2, thirdParameter * 2
}
