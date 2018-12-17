package captcha

import (
	"testing"
)

func TestCaptchaPatternOneLeftOneOperatorOneRightOneShouldReturn1PustOne(t *testing.T) {
	result := Captcha(1, 1, 1, 1)
	expected := "1 + one"
	if result != expected {
		t.Errorf("Expected %q, But result is %q", expected, result)
	}
}
