package tokensearch

import (
	"testing"
	"unicode/utf8"
)

func TestNewTokenNode(t *testing.T) {

	node := NewTokenNode()
	match := NewToken("1234", "encyclopædia", "noun")

	node.Insert(match)
	if node.nextLetters['e'] == nil {
		t.Errorf("Expected a child node with matching rune")
	}
}

func TestInsert(t *testing.T) {

	node := NewTokenNode()
	match := NewToken("1234", "encyclopædia", "noun")

	node.Insert(match)
	if node.nextLetters['e'] == nil {
		t.Errorf("Expected a child node with matching rune")
	}
	if node.nextLetters['n'] != nil {
		t.Errorf("Unexpected child node found.")
	}

	match2 := NewToken("2345", "nicely", "adverb")
	node.Insert(match2)
	if node.nextLetters['e'] == nil {
		t.Errorf("Expected a child node with matching rune")
	}
	if node.nextLetters['n'] == nil {
		t.Errorf("Expected a child node with matching rune")
	}
}

func TestNext(t *testing.T) {

	node := NewTokenNode()
	match1 := NewToken("1234", "Ruby on Rails", "framework")
	match2 := NewToken("2345", "Ruby", "language")
	match3 := NewToken("3456", "Ruby Tuesday", "restaurant")
	match4 := NewToken("4567", "ruby", "gemstone")
	node.Insert(match1)
	node.Insert(match2)
	node.Insert(match3)
	node.Insert(match4)

	str := "ruby"
	for i, w := 0, 0; i < len(str); i += w {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		node = node.Next(runeValue)
		if node == nil {
			t.Errorf("expected next %v to be found", runeValue)
		}
		w = width
	}

	if len(node.Values()) != 2 {
		t.Errorf("expect %v to have 2 matches : %v", node, node.Values())
	}
}

func TestFind(t *testing.T) {

	node := NewTokenNode()
	match := NewToken("1234", "Ruby on Rails", "noun")
	match2 := NewToken("2345", "nicely", "adverb")

	node.Insert(match)
	node.Insert(match2)

	matches := node.Find("Ruby on Rails")
	if len(matches) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matches))
	}
	if matches[0] != match {
		t.Errorf("Expected %s to equal %s", matches[0], match)
	}

	matchesLower := node.Find("ruby on rails")
	if len(matchesLower) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matchesLower))
	}
	if matchesLower[0] != match {
		t.Errorf("Expected %s to equal %s", matchesLower[0], match)
	}

	matchedDash := node.Find("ruby-on-rails")
	if len(matchedDash) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matchedDash))
	}
	if matchedDash[0] != match {
		t.Errorf("Expected %s to equal %s", matchedDash[0], match)
	}

	matches2 := node.Find("nicely")
	if len(matches2) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matches))
	}
	if matches2[0] != match2 {
		t.Errorf("Expected %s to equal %s", matches2[0], match2)
	}

	matches3 := node.Find("Ruby")
	if len(matches3) != 0 {
		t.Errorf("Expected %s length to eq 0", len(matches3))
	}

	matches4 := node.Find("encyclopædia with trailing words")
	if len(matches4) != 0 {
		t.Errorf("Expected %s length to eq 0", len(matches4))
	}
}

func TestRemove(t *testing.T) {

	node := NewTokenNode()
	match := NewToken("1234", "encyclopædia", "noun")
	match2 := NewToken("2345", "nicely", "adverb")

	node.Insert(match)
	node.Insert(match2)

	matches := node.Find("encyclopædia")
	if len(matches) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matches))
	}
	if matches[0] != match {
		t.Errorf("Expected %s to equal %s", matches[0], match)
	}

	node.Remove(match)
	matches = node.Find("encyclopædia")
	if len(matches) != 0 {
		t.Errorf("Expected %s length to eq 0", len(matches))
	}
}

func TestAllValues(t * testing.T) {
	node := NewTokenNode()
	match := NewToken("1234", "encyclopædia", "noun")
	match2 := NewToken("2345", "nicely", "adverb")

	node.Insert(match)
	node.Insert(match2)

	matches := node.AllValues(3)
	if len(matches) != 2 {
		t.Errorf("Expected there to be 2 matchs but got %v", matches)
	}
}
