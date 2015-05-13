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
		// fmt.Printf("%c", ch)
		pool.currPos += n
		if n == 0 {
			if err == nil {
				// fmt.Printf("#")
				continue
			}
			break
		}

		nch, _ := NormalizeRune(ch)
		pool.advanceWithState(nch, false)
	}
	pool.currPos++
	pool.advanceWithState(' ', true)
}

func (pool *TokenNodeVisitorPool) onMatch(matches []*TokenMatch) {
	if matches != nil && len(matches) > 0 {
		pool.Matches = append(pool.Matches, matches...)
	}
}

func (pool *TokenNodeVisitorPool) advanceWithState(ch rune, forceMatch bool) {

	pool.currRune = ch

	// 1. if last was a separator and this is not then initialize a visitor
	if pool.IsSeparator(pool.lastRune) && !pool.IsSeparator(ch) {
		// fmt.Printf("*")
		pool.initVisitor()
	}

	// 2. if this is a separator and last was not then save matches for all active visitors
	if pool.IsSeparator(ch) && !pool.IsSeparator(pool.lastRune) {
		// fmt.Printf("+")
		for _, visitor := range pool.activeVisitors {
			visitor.SaveMatches()
		}
	}

	// 3. advance through active visitors; any that become inactive are checked for matches
	for _, visitor := range pool.activeVisitors {
		visitor.Advance(ch)
		if forceMatch || !visitor.Active() {
			if visitor.LastMatches != nil {
				// fmt.Printf("!")
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
