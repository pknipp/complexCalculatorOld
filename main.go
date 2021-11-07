package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getNumber(expression string) (string, string){
	p := 1
	z := ""
	for len(expression) >= p {
		z = expression[0:p]
		// fmt.Println(z, expression[p:])
		if !(z == "." || z == "-" || z == "-.") {
			_, err := strconv.ParseFloat(z, 64)
			if err != nil {
				break
			}
		}
		p++
	}
	return expression[0:p - 1], expression[p - 1:]
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
	for len(expression) > 1 {
		op := expression[0:1]
		val, expression = getNumber(expression[1:])
		pair := map[string]string {"op": op, "val": val}
		fmt.Println(pair)
		pairs = append(pairs, pair)
		fmt.Println(pairs, expression)
	}
	fmt.Println(z, pairs)
}
