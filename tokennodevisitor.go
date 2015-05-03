package tokensearch

type TokenNodeVisitor struct {
	CurrentNode		*TokenNode
	LastMatches		[]*Token
	StartPos		int
	EndPos			int
}

func (visitor *TokenNodeVisitor) Active() bool {
	return visitor.CurrentNode != nil
}

func (visitor *TokenNodeVisitor) Advance(runeValue rune, position int, onMatch func([]*Token, int, int)) {
	if visitor.Active() {
		visitor.CurrentNode = visitor.CurrentNode.Next(runeValue)
		if visitor.CurrentNode == nil {
			if visitor.LastMatches != nil {
				onMatch(visitor.LastMatches, visitor.StartPos, visitor.EndPos)
			}
		} else {
			if len(visitor.CurrentNode.Values()) > 0 {
				visitor.LastMatches = visitor.CurrentNode.Values()
				visitor.EndPos = position
			}
		}
	}
}