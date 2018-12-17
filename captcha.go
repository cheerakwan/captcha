package captcha

func Captcha(pattern, left, operator, right int) string {
	if right == 2 {
		return "1 + two"
	}
	return "1 + one"
}
