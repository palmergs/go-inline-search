package tokensearch

import (
	"unicode"
)

type RuneReader interface {
	ReadRune() (rune, int, error)
}

type TokenNodeVisitorPool struct {
	root					*TokenNode
	inactiveVisitors 		map[int]*TokenNodeVisitor
	activeVisitors			map[int]*TokenNodeVisitor
	Matches					[]*TokenMatch
}

func NewTokenNodeVisitorPool(root *TokenNode) *TokenNodeVisitorPool {
	tnvp := &TokenNodeVisitorPool{root: root,
			inactiveVisitors: make(map[int]*TokenNodeVisitor),
			activeVisitors: make(map[int]*TokenNodeVisitor)}
	return tnvp
}

func (pool *TokenNodeVisitorPool) InitVisitor(position int) {
	if len(pool.inactiveVisitors) > 0 {
		for key, visitor := range pool.inactiveVisitors {
			visitor.Reset(pool.root, position)
			pool.activeVisitors[position] = visitor
			delete(pool.inactiveVisitors, key)
			break
		}
	} else {
		pool.activeVisitors[position] = NewTokenNodeVisitor(pool.root, position)
	}
}

func (pool *TokenNodeVisitorPool) ShouldInit(runeValue rune) bool {
	return (unicode.IsDigit(runeValue) || unicode.IsLetter(runeValue))
}

func (pool *TokenNodeVisitorPool) AdvanceThrough(reader RuneReader) {
	currPos := 0
	lastPos := 0
	lastWasChar := false
	for {
		ch, n, err := reader.ReadRune()
		currPos += n
		if n == 0 {
			if err == nil {
				continue
			}
			break
		}

		nch, currIsChar := NormalizeRune(ch)
		if currIsChar {
			pool.advanceWithState(nch, currPos, lastPos, currIsChar, lastWasChar)
			lastPos = currPos
		}
		lastWasChar = currIsChar
	}

	// advance through a final
	pool.advance('\n', currPos + 1)
}

func (pool *TokenNodeVisitorPool) advanceWithState(ch rune, currPos, lastPos int, currIsChar, lastWasChar bool) {
	if currPos > 0 && !lastWasChar {

		// advance for a deferred separator character for existing visitors
		pool.advance(' ', lastPos)
	}

	if currPos == 0 || !lastWasChar {

		// visitors begin parsing at beginning of valid strings
		pool.InitVisitor(currPos)
	}

	// advance of token character
	pool.advance(ch, currPos)
}

func (pool *TokenNodeVisitorPool) onMatch(matches []*TokenMatch) {
	if matches != nil && len(matches) > 0 {
		pool.Matches = append(pool.Matches, matches...)
	}
}

func (pool *TokenNodeVisitorPool) advance(runeValue rune, position int) {

	for _, visitor := range pool.activeVisitors {
		visitor.Advance(runeValue, pool.onMatch)
		if !visitor.Active() {
			pool.inactiveVisitors[visitor.StartPos] = visitor
			delete(pool.activeVisitors, visitor.StartPos)
		}
	}
}
