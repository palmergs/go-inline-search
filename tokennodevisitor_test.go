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

	success := searchDocument(root, `Learning ruby or Ruby on Rails, unlike pascal,
			requires the programmer to learn rudamentary regular expressions. This is more
			text that should be ignored. This is a test. Testing my brain. This is only a test!?`, 4)
	if !success {
		t.Errorf("Inline tokens not correctly found.\n")
	}

	success = searchDocument(root, `Searching for a token at end: ruby`, 1)
	if !success {
		t.Errorf("Couldn't find token at end of search string.\n")
	}

	success = searchDocument(root, `Ruby at start of document found?`, 1)
	if !success {
		t.Errorf("Couldn't find token at start of search string.\n")
	}

	success = searchDocument(root, `"ruby"`, 1)
	if !success {
		t.Errorf("Couldn't find token in quotes.\n")
	}
}

func searchDocument(root *TokenNode, doc string, expected int) bool {

	allMatches := make([]*TokenMatch, 0)
	onMatch := func(matches []*TokenMatch) {
		if matches != nil && len(matches) > 0 {
			allMatches = append(allMatches, matches...)
		}
	}

	document := NormalizeString(doc)
	pool := NewTokenNodeVisitorPool(root)
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
	pool.Advance('\n', 0, onMatch)

	return len(allMatches) == expected
}