package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getNumber(expression string) (float64, string){
	p := 1
	lastNum := 0.
	// err := nil
	for len(expression) >= p {
		z := expression[0:p]
		// fmt.Println(z, expression[p:])
		if !(z == "." || z == "-" || z == "-.") {
			num, err := strconv.ParseFloat(z, 64)
			if err != nil {
				break
			}
			lastNum = num
		}
		p++
	}
	return lastNum, expression[p - 1:]
}

// func calculate(x1 string, op string, x2 string) float64 {
//
// }

func main() {
	type opNum struct {
		op string
		num float64
	}
	expression := " +2.3 / 78 + 4. - 60 * .5 "
	expression = strings.ReplaceAll(expression, " ", "")
	if expression[0:1] == "+" {
		expression = expression[1:]
	}
	z := 0.
	z, expression = getNumber(expression)
	// fmt.Println(z, expression)
	// precedence := map[string]int{"+": 0, "-": 0, "*": 1, "/": 1, "^": 2}

	// pairs := make([]map[string]string, 0)
	pairs := []opNum{}
	num := 0.
	for len(expression) > 1 {
		op := expression[0:1]
		num, expression = getNumber(expression[1:])
		// pair := map[string]string {"op": op, "val": val}
		pair := opNum{op, num}
		// fmt.Println(pair)
		pairs = append(pairs, pair)
		fmt.Println(pairs, expression)
	}
	fmt.Println(z, pairs)
	// while len(pairs) > 1 {
	// 	index := 0
	// 	while len(pairs) > index {
	// 		if precedence[pairs[index].op] < precedence[pairs[index + 1].op] {
	// 			index++
	// 		} else {
	// 			if index == 0 {
	// 				x1 := z
	// 			} else {
	// 				x1 := pairs[index].val
	// 			}
	// 			result := calculate(x1, pairs[index].op, pairs[index].val)
	// 		}
	// 	}
	// 	// 		replace z or val property of element with result


	// 	if precedence[pairs[0].op] <
	// }
}
