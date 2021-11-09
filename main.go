package main

import (
	"fmt"
	"io"
	// "log"
	"math/cmplx"
	"net/http"
	"strconv"
	"strings"
	// "math"
)



func calculate(z1 complex128, op string, z2 complex128) complex128 {
	result := complex(0., 0.)
	switch op {
	case "+":
		result = z1 + z2
	case "-":
		result = z1 - z2
	case "*":
		result = z1 * z2
	case "/":
		result = z1 / z2
	case "^":
		result = cmplx.Pow(z1, z2)
	}
	return result
}

func parseExpression (expression string) (complex128) {
	// fmt.Println("line 34 says that expression = ", expression)
	getNumber := func(expression string) (complex128, string){
		// fmt.Println("line 36 says that expression = ", expression)
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
			return parseExpression(expression[1: nExpression]), expression[nExpression + 1:]
		} else if leadingChar == "i" || leadingChar == "j" {
			// fmt.Println("line 52 says that expression = ", expression)
			return complex(0, 1), expression[1:]
		} else {
			// fmt.Println("line 55 says that expression = ", expression)
			p := 1
			lastNum := complex(0., 0.)
			for len(expression) >= p {
				z := expression[0:p]
				// If implied multiplication is detected ...
				if expression[p - 1: p] == "(" {
					// ... insert a "*" symbol.
					expression = expression[0:p - 1] + "*" + expression[p - 1:]
					break
				} else if !(z == "." || z == "-" || z == "-.") {
					num, err := strconv.ParseFloat(z, 64)
					if err != nil {
						break
					}
					lastNum = complex(num, 0.)
				}
				p++
			}
			return lastNum, expression[p - 1:]
		}
	}
	type opNum struct {
		op string
		num complex128
	}

	expression = strings.ReplaceAll(expression, " ", "")
	if expression[0:1] == "+" {
		expression = expression[1:]
	}
	z := complex(0., 0.)
	z, expression = getNumber(expression)
	precedence := map[string]int{"+": 0, "-": 0, "*": 1, "/": 1, "^": 2}
	ops := "+-*/^"
	pairs := []opNum{}
	num := complex(0., 0.)
	for len(expression) > 0 {
		op := expression[0:1]
		if strings.Contains(ops, op) {
			// fmt.Println("line 95")
			expression = expression[1:]
		} else {
			op = "*"
		}
		num, expression = getNumber(expression)
		pair := opNum{op, num}
		pairs = append(pairs, pair)
	}
	for len(pairs) > 0 {
		index := 0
		for len(pairs) > index {
			if index < len(pairs) - 1 && precedence[pairs[index].op] < precedence[pairs[index + 1].op] {
				index++
			} else {
				z1 := complex(0., 0.)
				if index == 0 {
					z1 = z
				} else {
					z1 = pairs[index - 1].num
				}
				result := calculate(z1, pairs[index].op, pairs[index].num)
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
	// fmt.Println("Starting server on port 8000")
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8000", nil)
	var expression string = "1+2i/(3-4j/(5+6j))"
	fmt.Println(parseExpression(expression));
}

func handler(w http.ResponseWriter, r*http.Request) {
	io.WriteString(w, "Here is the answer: \n")
	expression := r.URL.Path
	if expression != "/favicon.ico" {
		if len(expression) > 1 {
			// fmt.Println("line 139")
			expression = expression[1:]
			// fmt.Println(expression)
			result := parseExpression(expression)
			// fmt.Println(result)
			realPart := strconv.FormatFloat(real(result), 'f', -1, 64)
			imagPart := strconv.FormatFloat(imag(result), 'f', -1, 64)
			resultString := realPart + " + " + imagPart + "i"
			io.WriteString(w, resultString)
		}
	}
}

	// keys, ok := r.URL.Query()["key"]
	// if !ok || len(keys[0]) < 1 {
		// log.Println("Url Param 'key' is missing")
		// return
	// }
	// Query()["key"] will return an array of items,
	// we only want the single item.
	// key := keys[0]
	// fmt.Println("line 146")
	// fmt.Println(key)
	// result := parseExpression(key)
	// fmt.Println("line 148")
	// fmt.Println(result)
	// log.Println("Value of expression is: " + string(real(result)))
