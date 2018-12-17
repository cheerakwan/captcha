package captcha

import "strconv"

var numbers = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
	3: "three",
	4: "four",
	5: "five",
	6: "six",
	7: "seven",
	8: "eight",
	9: "nine",
}

var operators = map[int]string{
	1: "+",
	2: "-",
	3: "*",
}

func Captcha(pattern, left, operator, right int) string {
	if pattern == 2 {
		num := numbers[left]
		oper := operators[operator]
		return num + " " + oper + " " + strconv.Itoa(right)
	}
	num := numbers[right]
	return strconv.Itoa(left) + " + " + num

}
