package main

import (
	"fmt"
	"strconv"
	"strings"
	// "unicode"
)

func main() {
	strin := " +2.3 + 4 * 5 "
	strin = strings.ReplaceAll(strin, " ", "")
	if strin[0:1] == "+" {
		strin = strin[1:]
	}
	expression := strings.Split(strin, "")
	z := ""
	for {
		zTemp := z + expression[0]
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
	fmt.Println(z, expression)
	// for len(expression) > 0 {
		// op :=
	// }
}
