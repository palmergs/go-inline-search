package tokensearch

import (
	"testing"
	"unicode/utf8"
)

func TestNewTokenNodeVisitor(t *testing.T) {
	root := NewTokenNode()
	visitor := NewTokenNodeVisitor(root, 0)
	if visitor == nil {
		t.Errorf("Expected visitor to not be nil")
	}
}

func TestAdvance(t *testing.T) {
	root := NewTokenNode()
	for _, str := range []string{"Ruby", "Ruby on Rails", "Pascal", "Regular Expression"} {
		root.Insert(NewToken(str, str, "technology"))
	}

	allMatches := make([]*Token, 0)
	onMatch := func(matches []*Token, startPos int, endPos int) {
		if matches != nil && len(matches) > 0 {
			allMatches = append(allMatches, matches...)
		}
	}

	pool := NewTokenNodeVisitorPool(root)
	document := "Learning ruby or Ruby on Rails, unlike pascal, requires the programmer to learn regular expression."
	for i, w := 0, 0; i < len(document); i += w {
		runeValue, width := utf8.DecodeRuneInString(document[i:])
		w = width

		pool.Advance(runeValue, i, onMatch)
	}

	if len(allMatches) != 4 {
		t.Errorf("Expected to find 4 matches in document but found %d", len(allMatches))
	}

}