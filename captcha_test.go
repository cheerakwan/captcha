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

func TestCaptchaPatternTwoLeftOneOperatorOneRightOneShouldReturnOnePlus1(t *testing.T) {
	result := Captcha(2, 1, 1, 1)
	expected := "one + 1"
	if result != expected {
		t.Errorf("Expected %q, But result is %q", expected, result)
	}
}

func TestCaptchaPatternTwoLeftTwoOperatorOneRightOneShouldReturnTwoPlus1(t *testing.T) {
	result := Captcha(2, 2, 1, 1)
	expected := "two + 1"
	if result != expected {
		t.Errorf("Expected %q, But result is %q", expected, result)
	}
}

func TestCaptchaPatternTwoLeftTwoOperatorOneRightTwoShouldReturnTwoPlusTwo(t *testing.T) {
	result := Captcha(2, 2, 1, 2)
	expected := "two + 2"
	if result != expected {
		t.Errorf("Expected %q, But result is %q", expected, result)
	}
}
