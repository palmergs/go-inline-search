package tokensearch

import "testing"

func TestMatchesInString(t *testing.T) {

	tmp, _ := MatchesInString("This is a test of Ruby-on-Rails.")
	if tmp == nil {
		t.Errorf("Errored...")
	}
}