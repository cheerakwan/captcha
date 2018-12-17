package captcha

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

func Captcha(pattern, left, operator, right int) string {
	if pattern == 2 {
		num := numbers[left]
		return num + " + 1"
	} else {
		num := numbers[right]
		return "1 + " + num
	}
}
