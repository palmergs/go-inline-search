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

func (pool *TokenNodeVisitorPool) ShouldInit(runeValue rune) bool {
	return (unicode.IsDigit(runeValue) || unicode.IsLetter(runeValue))
}

func (pool *TokenNodeVisitorPool) Advance(runeValue rune, position int, onMatch func([]*TokenMatch)) {

	// fmt.Printf("There are %d visitors for %c\n", len(pool.activeVisitors), runeValue)
	for _, visitor := range pool.activeVisitors {
		visitor.Advance(runeValue, onMatch)
		if !visitor.Active() {
			pool.inactiveVisitors[visitor.StartPos] = visitor
			delete(pool.activeVisitors, visitor.StartPos)
		}
	}
}
