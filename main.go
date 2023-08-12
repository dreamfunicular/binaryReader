package main

import (
	"fmt"
	"math"
)

func main() {
	inputString := "01010111 01101000 01100001 01110100 00100111 01110011 00100000 01110101 01110000 00101100 00100000 01101110 01100101 01110010 01100100 01110011 00100001"
	currentValue := 0
	outputString := ""

	for i, letter := range inputString {
		if letter == ' ' {
			outputString = outputString + string(currentValue)
			currentValue = 0
		} else {
			if letter == '1' {
				currentValue += int(math.Pow(2, float64(7-(i%9))))
			}
		}
	}

	fmt.Println(outputString)

}
