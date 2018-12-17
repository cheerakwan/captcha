package captcha

import (
	"testing"
)

func TestCaptchaPatternOneLeftOneOperatorOneRightOneShouldReturn1PlusOne(t *testing.T) {
	result := Captcha(1, 1, 1, 1)
	expected := "1 + one"
	if result != expected {
		t.Errorf("Expected %q, But result is %q", expected, result)
	}
}

func TestCaptchaPatternOneLeftOneOperatorOneRightTwoShouldReturn1PlusTwo(t *testing.T) {
	result := Captcha(1, 1, 1, 2)
	expected := "1 + two"
	if result != expected {
		t.Errorf("Expected %q, But result is %q", expected, result)
	}
}

func TestCaptchaPatternOneLeftOneOperatorOneRightThreeShouldReturn1PlusThree(t *testing.T) {
	result := Captcha(1, 1, 1, 3)
	expected := "1 + three"
	if result != expected {
		t.Errorf("Expected %q, But result is %q", expected, result)
	}
}
