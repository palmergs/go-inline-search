package tokensearch

import "testing"

func TestNewToken(t *testing.T) {
	match := NewToken(1234, "Ruby on-Rails", "noun")
	if match.Key() != "ruby on rails" {
		t.Errorf("Expected %v be generated from %v", match.Key(), match.Display)
	}
}

func TestNormalizeString(t *testing.T) {

	trimFront := "    this is a test"
	if str := NormalizeString(trimFront); str != "this is a test" {
		t.Errorf("Expected %v to trim left of %v", str, trimFront)
	}

	trimMiddle := "this \n\n  is    a  \t\n test"
	if str := NormalizeString(trimMiddle); str != "this is a test" {
		t.Errorf("Expected %v to trim within %v", str, trimMiddle)
	}

	trimEnd := "this is a test    \n\n"
	if str := NormalizeString(trimEnd); str != "this is a test" {
		t.Errorf("Expected %v to trim end of %v", str, trimEnd)
	}

	normalizeDash := "-this--is---a----test-----"
	if str := NormalizeString(normalizeDash); str != "this is a test" {
		t.Errorf("Expected %v to remove dashes from %v", str, normalizeDash)
	}

	normalizeUCase := "This IS A tEST"
	if str := NormalizeString(normalizeUCase); str != "this is a test" {
		t.Errorf("Expected %v to set lowercase of %v", str, normalizeUCase)
	}
}
