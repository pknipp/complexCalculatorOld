package getNumber
import "strconv"

func GetNumber(expression string) (string, string){
	z := ""
	for len(expression) > 0 {
		zTemp := z + expression[0:1]
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
