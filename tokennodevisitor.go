package tokensearch

type TokenNodeVisitor struct {
	CurrentNode		*TokenNode
	LastMatches		[]*Token
	StartPos		int
	EndPos			int
}

func NewTokenNodeVisitor(node *TokenNode, startAt int) *TokenNodeVisitor {
	return &TokenNodeVisitor{CurrentNode: node, LastMatches: make([]*Token, 0), StartPos: startAt, EndPos: startAt}
}

func (visitor *TokenNodeVisitor) Reset(node *TokenNode, startAt int) *TokenNodeVisitor {
	visitor.CurrentNode = node
	visitor.StartPos = startAt
	visitor.EndPos = startAt
	visitor.LastMatches = make([]*Token, 0)
	return visitor
}

func (visitor *TokenNodeVisitor) Active() bool {
	return visitor.CurrentNode != nil
}

func (visitor *TokenNodeVisitor) Advance(runeValue rune, onMatch func([]*Token, int, int)) {
	if visitor.Active() {

		if len(visitor.CurrentNode.Values()) > 0 {
			visitor.LastMatches = visitor.CurrentNode.Values()
		}

		visitor.CurrentNode = visitor.CurrentNode.Next(runeValue)
		visitor.EndPos += 1

		if visitor.CurrentNode == nil {
			if visitor.LastMatches != nil {
				onMatch(visitor.LastMatches, visitor.StartPos, visitor.EndPos)
			}
		}
	}
}