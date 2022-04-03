package main

import (
	"fmt"
)

func main() {
	var a = "text \n text"
	var b = `text \n text`
	var c rune = 960
	var d = '*'
	fmt.Printf("%v\n%v\n%c\n%[4]T\n%c\n", a, b, c, d, b[1])
	const mess = "uv vagreangvbany fcnpr fgngvba"
	rot13Decoder(mess)
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
