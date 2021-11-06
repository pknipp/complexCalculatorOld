package main

import (
	"fmt"
	"strconv"
	// "unicode"
)

func main() {
	expression := "2.3+4*5"
	z := ""
	zTemp := z
	for true {
		zTemp = z + expression[0:1]
		fmt.Println(zTemp, expression)
		if !(zTemp == "." || zTemp == "-" || zTemp == "-.") {
			_, err := strconv.ParseFloat(zTemp, 64)
			if err != nil {
				break
			}
		}
		expression = expression[1:]
		z = zTemp
	}
	//
	// fmt.Println(unicode.IsNumber('2.'))
	// fmt.Println(expression[0])
	// fmt.Println(isANumber)
	fmt.Println(z, expression)
}
