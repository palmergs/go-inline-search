package tokensearch

import (
	"testing"
	"strings"
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
			"regex"} {
		root.Insert(NewToken(str, str, "technology"))
	}
	return root
}

func TestAdvance(t *testing.T) {

	root := buildTree()

	// matches := searchDocument(root, `Learning ruby or Ruby on Rails, unlike pascal,
	// 		requires the programmer to learn rudamentary regex. This is more
	// 		text that should be ignored. This is a test. Testing my brain. This is only a test!?`)
	// if len(matches) != 4 {
	// 	t.Errorf("Inline tokens not correctly found. (%v)\n", matches)
	// }

	matches := searchDocument(root, `Searching for a token at end: ruby`)
	if len(matches) != 1 {
		t.Errorf("Couldn't find token at end of search string. (%v)\n", matches)
	}

	matches = searchDocument(root, `Ruby at start of document found?`)
	if len(matches) != 1 {
		t.Errorf("Couldn't find token at start of search string.\n")
	}

	matches = searchDocument(root, `"ruby"`)
	if len(matches) != 1 {
		t.Errorf("Couldn't find token in quotes.\n")
	}

	matches = searchDocument(root, `ruby on rails`)
	if len(matches) != 1 {
		t.Errorf("Couldn't find token by itself\n")
	}

	matches = searchDocument(root, `rubylicious`)
	if len(matches) != 0 {
		t.Errorf("Incorrectly found token that was prefix to larger word")
	}

	matches = searchDocument(root, `baruby`)
	if len(matches) != 0 {
		t.Errorf("Incorrectly found token that was suffix to larger word")
	}
}

func searchDocument(root *TokenNode, doc string) []string {

	document := NormalizeString(doc)
	pool := NewTokenNodeVisitorPool(root)
	pool.AdvanceThrough(strings.NewReader(document))

	matches := make([]string, len(pool.Matches))
	for idx, tokenMatch := range pool.Matches {
		matches[idx] = tokenMatch.Token.Key()
	}
	return matches
}