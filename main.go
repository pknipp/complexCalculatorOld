package main

import (
	"fmt"
	"strconv"
	"strings"
	// "unicode"
)

func getNumber(expression string) (string, string){
	z := ""
	for len(expression) > 0 {
		zTemp := z + expression[0:1]
		// fmt.Println(zTemp, expression)
		if !(zTemp == "." || zTemp == "-" || zTemp == "-.") {
			_, err := strconv.ParseFloat(zTemp, 64)
			if err != nil {
				break
			}
		}
		expression = expression[1:]
		z = zTemp
	}
	return z, expression
}

func main() {
	expression := " +2.3 / 78 + 4. - 60 * .5 "
	expression = strings.ReplaceAll(expression, " ", "")
	if expression[0:1] == "+" {
		expression = expression[1:]
	}
	z := ""
	z, expression = getNumber(expression)
	fmt.Println(z, expression)
	// precedence := map[string]int{"+": 0, "-": 0, "*": 1, "/": 1, "^": 2}

	pairs := make([]map[string]string, 0)
	val := ""
	for len(expression) > 0 {
		op := expression[0:1]
		val, expression = getNumber(expression[1:])
		pair := map[string]string {"op": op, "val": val}
		fmt.Println(pair)
		pairs = append(pairs, pair)
		fmt.Println(pairs, expression)
	}
	fmt.Println(z, pairs)
}
