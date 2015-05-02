package tokensearch

import "testing"

func TestMatchesInString(t *testing.T) {

	root, _ := NewTokenNode()
	root.Insert(NewTokenMatch("1234", "Ruby on Rails", "framework"))
	root.Insert(NewTokenMatch("2345", "Go", "language"))
	tmp, _ := MatchesInString(root, "This is a test of Ruby-on-Rails.")
	if tmp == nil {
		t.Errorf("Errored...")
	}
}