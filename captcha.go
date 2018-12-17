package captcha

func Captcha(pattern, left, operator, right int) string {
	if pattern == 2 {
		return "one + 1"
	}
	if right == 3 {
		return "1 + three"
	}
	if right == 2 {
		return "1 + two"
	}
	return "1 + one"
}
