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
			"A+",
			"O'Malley",
			"encyclopædia",
			"rain",
			"#hashtag",
			"@twitter",
			"123.123.123.123",
			"bob@example.com",
			"ruby on airplanes",
			"Pascal",
			"regex"} {
		root.Insert(NewToken(str, str, "technology"))
	}
	return root
}

func TestAdvance(t *testing.T) {

	root := buildTree()

	matches := searchDocument(root, `ruby on rails`)
	if len(matches) != 1 {
		t.Errorf("Couldn't find token by itself\n")
	}

	matches = searchDocument(root, `Learning ruby or Ruby on Rails, unlike pascal,
			requires the programmer to learn rudamentary regex. This is more
			text that should be ignored. This is a test. Testing my hashtag. This is only a twitter!?`)
	if len(matches) != 4 {
		t.Errorf("Inline tokens not correctly found. (%v)\n", matches)
	}

	matches = searchDocument(root, `Searching for a token at end: ruby`)
	if len(matches) != 1 {
		t.Errorf("Couldn't find token at end of search string. (%v)\n", matches)
	}

	matches = searchDocument(root, `Ruby at start of document found?`)
	if len(matches) != 1 {
		t.Errorf("Couldn't find token at start of search string.\n")
	}

	matches = searchDocument(root, `"ruby" or 'regex' or [pascal] or (rain)`)
	if len(matches) != 4 {
		t.Errorf("Couldn't find token in quotes. (%v)\n", matches)
	}

	matches = searchDocument(root, `rubylicious`)
	if len(matches) != 0 {
		t.Errorf("Incorrectly found token that was prefix to larger word")
	}

	matches = searchDocument(root, `baruby`)
	if len(matches) != 0 {
		t.Errorf("Incorrectly found token that was suffix to larger word")
	}

	matches = searchDocument(root, `a good certification is A+ for computers`)
	if len(matches) != 1 {
		t.Errorf("Expected to find a single match for A+. (%v)\n", matches)
	}

	matches = searchDocument(root, `O'Malley read the encyclopedia or, as he called it, the 'encyclopædia'`)
	if len(matches) != 2 {
		t.Errorf("Expected to match an apostrophe in a word as well as a non-ASCII char. (%v)\n", matches)
	}

	matches = searchDocument(root, `I can read a #hashtag hashtag and a @twitter handle and bob@example.com email.`)
	if len(matches) != 3 {
		t.Errorf("Expected to match a hashtag, email and twitter handle. (%v)\n", matches)
	}

	matches = searchDocument(root, `There are 123. 123.123.123.123 At the ip address 123.123.123.123.`)
	if len(matches) != 2 {
		t.Errorf("Expected to match an IP address (with or without trailing punctuation). (%v)\n", matches)
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