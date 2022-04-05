package main

import (
	"fmt"
)

func main() {
	textTest()
	const mess = "uv vagreangvbany fcnpr fgngvba"
	rot13Decoder(mess)
	convertor()
}

func textTest() {
	var a = "text1\ntext2"
	var b = `text1\ntext2`
	var c rune = 960
	var d = '*'
	fmt.Printf("%v\n%v\n%c\n%[4]T\n%c\n", a, b, c, d, b[1])
	fmt.Println("\n")
}

func rot13Decoder(mess string) {
	for i := 0; i < len(mess); i++ {
		var letter = mess[i] + 13
		if letter > 'z' {
			letter = letter - 26
		}
		fmt.Printf("%c", letter)
	}
}

func convertor() {
	var a = 10
	var b = float64(a)
	fmt.Println("\n")
	fmt.Println(b)
}
