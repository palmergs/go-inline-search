package tokensearch

import (
	"unicode"
)

type TokenNodeVisitorPool struct {
	root					*TokenNode
	inactiveVisitors 		map[int]*TokenNodeVisitor
	activeVisitors			map[int]*TokenNodeVisitor
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

func (pool *TokenNodeVisitorPool) NormalizeRune(runeValue rune) rune {
	if unicode.IsLetter(runeValue) {
		return unicode.ToLower(runeValue)
	}
	return runeValue
}

func (pool *TokenNodeVisitorPool) ShouldInit(runeValue rune) bool {
	if unicode.IsDigit(runeValue) || unicode.IsLetter(runeValue) {
		return true
	}
	return false
}

func (pool *TokenNodeVisitorPool) Advance(runeValue rune, position int, onMatch func([]*Token, int, int)) {

	runeValue = pool.NormalizeRune(runeValue)

	if pool.ShouldInit(runeValue) {
		pool.InitVisitor(position)
	}

	for _, visitor := range pool.activeVisitors {
		visitor.Advance(runeValue, onMatch)
		if !visitor.Active() {
			pool.inactiveVisitors[visitor.StartPos] = visitor
			delete(pool.activeVisitors, visitor.StartPos)
		}
	}
}

