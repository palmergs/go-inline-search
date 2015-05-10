package tokensearch

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestNewTokenNodeVisitor(t *testing.T) {
	root := NewTokenNode()
	visitor := NewTokenNodeVisitor(root, 0)
	if visitor == nil {
		t.Errorf("Expected visitor to not be nil")
	}
}

func buildTree() *TokenNode {
	root := NewTokenNode()
	for _, str := range []string{"Ruby",
			"Ruby on Rails",
			"Weasel",
			"Badger",
			"red",
			"rust",
			"rain",
			"ruby on airplanes",
			"Pascal",
			"Regular Expression"} {
		root.Insert(NewToken(str, str, "technology"))
	}
	return root
}

func TestAdvance(t *testing.T) {

	root := buildTree()

	allMatches := make([]*TokenMatch, 0)
	onMatch := func(matches []*TokenMatch) {
		if matches != nil && len(matches) > 0 {
			allMatches = append(allMatches, matches...)
		}
	}

	pool := NewTokenNodeVisitorPool(root)
	document := NormalizeString(`Learning ruby or Ruby on Rails, unlike pascal,
			requires the programmer to learn rudamentary regular expressions. This is more
			text that should be ignored. This is a test. Testing my brain. This is only a test!?`)
	currIsChar := false
	lastWasChar := false
	for i, w := 0, 0; i < len(document); i += w {
		runeValue, width := utf8.DecodeRuneInString(document[i:])
		w = width

		currIsChar = (unicode.IsLetter(runeValue) || unicode.IsDigit(runeValue))
		if currIsChar && !lastWasChar {
			pool.InitVisitor(i)
		}
		lastWasChar = currIsChar

		pool.Advance(runeValue, i, onMatch)
	}

	if len(allMatches) != 4 {
		t.Errorf("Expected to find 4 matches in document but found %d :: %s", len(allMatches), allMatches)
	}
}