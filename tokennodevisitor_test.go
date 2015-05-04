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

	activeVisitors := make(map[int]*TokenNodeVisitor)
	inactiveVisitors := make(map[int]*TokenNodeVisitor)
	document := "Learning ruby or Ruby on Rails, unlike pascal, requires the programmer to learn regular expression."
	for i, w := 0, 0; i < len(document); i += w {
		runeValue, width := utf8.DecodeRuneInString(document[i:])
		if unicode.IsLetter(runeValue) {
			runeValue = unicode.ToLower(runeValue)
		}
		w = width

		if len(inactiveVisitors) > 0 {
			for key, visitor := range inactiveVisitors {
				visitor.Reset(root, i)
				activeVisitors[i] = visitor
				delete(inactiveVisitors, key)
				break
			}
		} else {
			activeVisitors[i] = NewTokenNodeVisitor(root, i)
		}

		for _, visitor := range activeVisitors {
			visitor.Advance(runeValue, onMatch)
			if !visitor.Active() {
				inactiveVisitors[visitor.StartPos] = visitor
				delete(activeVisitors, visitor.StartPos)
			}
		}
	}

	if len(allMatches) != 4 {
		t.Errorf("Expected to find 4 matches in document but found %d", len(allMatches))
	}

}