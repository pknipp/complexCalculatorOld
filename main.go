package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
)



func calculate(x1 float64, op string, x2 float64) float64 {
	result := 0.
	switch op {
	case "+":
		result = x1 + x2
	case "-":
		result = x1 - x2
	case "*":
		result = x1 * x2
	case "/":
		result = x1 / x2
	case "^":
		result = math.Pow(x1, x2)
	}
	return result
}

func parseExpression (expression string) (float64) {
	getNumber := func(expression string) (float64, string){
		leadingChar := expression[0:1]
		if leadingChar == "(" {
			nExpression := 0
			nParen := 1
			for nParen > 0 {
				nExpression++
				nextChar := expression[nExpression: nExpression + 1]
				if nextChar == "(" {
					nParen++
				} else if nextChar == ")" {
					nParen--
				}
			}
			return parseExpression(expression[1: nExpression]), expression[nExpression:]
		} else {
			p := 1
			lastNum := 0.
			for len(expression) >= p {
				z := expression[0:p]
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
	}
	type opNum struct {
		op string
		num float64
	}

	expression = strings.ReplaceAll(expression, " ", "")
	if expression[0:1] == "+" {
		expression = expression[1:]
	}
	z := 0.
	z, expression = getNumber(expression)
	precedence := map[string]int{"+": 0, "-": 0, "*": 1, "/": 1, "^": 2}
	pairs := []opNum{}
	num := 0.
	for len(expression) > 1 {
		op := expression[0:1]
		num, expression = getNumber(expression[1:])
		pair := opNum{op, num}
		pairs = append(pairs, pair)
	}
	for len(pairs) > 1 {
		// fmt.Println(z, pairs)
		index := 0
		for len(pairs) > index {
			// fmt.Println(index, z, pairs)
			if index < len(pairs) - 1 && precedence[pairs[index].op] < precedence[pairs[index + 1].op] {
				index++
			} else {
				x1 := 0.
				if index == 0 {
					x1 = z
				} else {
					x1 = pairs[index - 1].num
				}
				result := calculate(x1, pairs[index].op, pairs[index].num)
				if index == 0 {
					z = result
					pairs = pairs[1:]
				} else {
					pairs[index - 1].num = result
					pairs = append(pairs[0: index], pairs[index + 1:]...)
				}
				index = 0
			}
		}
	}
	return z
}

func main() {
	var expression string = " +23. / 78 ^ 2 + 4. * 6. - .5 "
	fmt.Println(parseExpression(expression));
}
