package main

import (
	"fmt"
)

type celsius float64
type kelvin float64
type fahrenheit float64

func main() {
	var c celsius = 100
	var k kelvin = 500
	var f fahrenheit = 200
	var temp = c + k.celsius() + f.celsius()
	fmt.Println(temp)
}

func (k kelvin) celsius() celsius {
	return celsius(k - 237)
}
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}
