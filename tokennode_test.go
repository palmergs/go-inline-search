package tokensearch

import "testing"

func TestNewTokenNode(t *testing.T) {

	node, _ := NewTokenNode()
	match := NewTokenMatch("1234", "encyclopædia", "noun")

	node.Insert(match)
	if node.nextLetters['e'] == nil {
		t.Errorf("Expected a child node with matching rune")
	}
}

func TestInsert(t *testing.T) {

	node := TokenNode{nextLetters: make(map[rune]*TokenNode), matches: make(map[string]*TokenMatch)}
	match := NewTokenMatch("1234", "encyclopædia", "noun")

	node.Insert(match)
	if node.nextLetters['e'] == nil {
		t.Errorf("Expected a child node with matching rune")
	}
	if node.nextLetters['n'] != nil {
		t.Errorf("Unexpected child node found.")
	}

	match2 := NewTokenMatch("2345", "nicely", "adverb")
	node.Insert(match2)
	if node.nextLetters['e'] == nil {
		t.Errorf("Expected a child node with matching rune")
	}
	if node.nextLetters['n'] == nil {
		t.Errorf("Expected a child node with matching rune")
	}
}

func TestFind(t *testing.T) {

	node := TokenNode{nextLetters: make(map[rune]*TokenNode), matches: make(map[string]*TokenMatch)}
	match := NewTokenMatch("1234", "Ruby on Rails", "noun")
	match2 := NewTokenMatch("2345", "nicely", "adverb")

	node.Insert(match)
	node.Insert(match2)

	matches, _ := node.Find("Ruby on Rails")
	if len(matches) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matches))
	}
	if matches[0] != match {
		t.Errorf("Expected %s to equal %s", matches[0], match)
	}

	matchesLower, _ := node.Find("ruby on rails")
	if len(matchesLower) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matchesLower))
	}
	if matchesLower[0] != match {
		t.Errorf("Expected %s to equal %s", matchesLower[0], match)
	}

	matchedDash, _ := node.Find("ruby-on-rails")
	if len(matchedDash) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matchedDash))
	}
	if matchedDash[0] != match {
		t.Errorf("Expected %s to equal %s", matchedDash[0], match)
	}

	matches2, _ := node.Find("nicely")
	if len(matches2) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matches))
	}
	if matches2[0] != match2 {
		t.Errorf("Expected %s to equal %s", matches2[0], match2)
	}

	matches3, _ := node.Find("Ruby")
	if len(matches3) != 0 {
		t.Errorf("Expected %s length to eq 0", len(matches3))
	}

	matches4, _ := node.Find("encyclopædia with trailing words")
	if len(matches4) != 0 {
		t.Errorf("Expected %s length to eq 0", len(matches4))
	}
}

func TestRemove(t *testing.T) {

	node := TokenNode{nextLetters: make(map[rune]*TokenNode), matches: make(map[string]*TokenMatch)}
	match := NewTokenMatch("1234", "encyclopædia", "noun")
	match2 := NewTokenMatch("2345", "nicely", "adverb")

	node.Insert(match)
	node.Insert(match2)

	matches, _ := node.Find("encyclopædia")
	if len(matches) != 1 {
		t.Errorf("Expected %s length to eq 1", len(matches))
	}
	if matches[0] != match {
		t.Errorf("Expected %s to equal %s", matches[0], match)
	}

	node.Remove(match)
	matches, _ = node.Find("encyclopædia")
	if len(matches) != 0 {
		t.Errorf("Expected %s length to eq 0", len(matches))
	}
}

