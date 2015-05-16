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
	currPos					int
	lastPos					int
	currRune				rune
	lastRune				rune
	Matches					[]*TokenMatch
}

func NewTokenNodeVisitorPool(root *TokenNode) *TokenNodeVisitorPool {
	tnvp := &TokenNodeVisitorPool{root: root,
			inactiveVisitors: make(map[int]*TokenNodeVisitor),
			activeVisitors: make(map[int]*TokenNodeVisitor),
			currPos: 0,
			lastPos: 0,
			currRune: ' ',
			lastRune: ' '}
	return tnvp
}

func (pool *TokenNodeVisitorPool) initVisitor() {
	if len(pool.inactiveVisitors) > 0 {
		for key, visitor := range pool.inactiveVisitors {
			visitor.Reset(pool.root, pool.currPos)
			pool.activeVisitors[pool.currPos] = visitor
			delete(pool.inactiveVisitors, key)
			break
		}
	} else {
		pool.activeVisitors[pool.currPos] = NewTokenNodeVisitor(pool.root, pool.currPos)
	}
}

func (pool *TokenNodeVisitorPool) IsSeparator(ch rune) bool {
	return unicode.IsSpace(ch) || unicode.IsPunct(ch) || unicode.IsSymbol(ch)
}

func (pool *TokenNodeVisitorPool) AdvanceThrough(reader RuneReader) {
	for {
		ch, n, err := reader.ReadRune()
		if n == 0 {
			if err == nil {
				continue
			}
			break
		}

		pool.advanceWithState(ch, false)
		pool.currPos += n
	}

	pool.advanceWithState(' ', true)
	pool.currPos++
}

func (pool *TokenNodeVisitorPool) onMatch(matches []*TokenMatch) {
	if matches != nil && len(matches) > 0 {
		pool.Matches = append(pool.Matches, matches...)
	}
}

func (pool *TokenNodeVisitorPool) advanceWithState(ch rune, forceMatch bool) {

	pool.currRune = ch
	nch, _ := NormalizeRune(ch)

	// 1. if last was a separator and this is not whitespace then initialize a visitor
	if pool.IsSeparator(pool.lastRune) && !unicode.IsSpace(ch) {
		pool.initVisitor()
	}

	// 2. if this is a separator and last was not whitespace then save matches for all active visitors
	if pool.IsSeparator(ch) && !unicode.IsSpace(pool.lastRune) {
		for _, visitor := range pool.activeVisitors {
			visitor.SaveMatches()
		}
	}

	// 3. advance through active visitors; any that become inactive are checked for matches
	chars := []rune{ch, nch}
	for _, visitor := range pool.activeVisitors {
		visitor.Advance(chars)
		if forceMatch || !visitor.Active() {
			if visitor.LastMatches != nil {
				pool.onMatch(visitor.LastMatches)
			}
			pool.deactivateVisitor(visitor)
		}
	}
	pool.lastRune = ch
	pool.lastPos = pool.currPos
}

func (pool *TokenNodeVisitorPool) deactivateVisitor(visitor *TokenNodeVisitor) {
	pool.inactiveVisitors[visitor.StartPos] = visitor
	delete(pool.activeVisitors, visitor.StartPos)
}
