package tokensearch

type TokenMatch struct {
	Token			*Token
	StartPos		int
	EndPos			int
}

type TokenNodeVisitor struct {
	CurrentNode		*TokenNode
	LastMatches		[]*TokenMatch
	StartPos		int
	EndPos			int
}

func NewTokenNodeVisitor(node *TokenNode, startAt int) *TokenNodeVisitor {
	return &TokenNodeVisitor{CurrentNode: node, LastMatches: nil, StartPos: startAt, EndPos: startAt}
}

func (visitor *TokenNodeVisitor) Reset(node *TokenNode, startAt int) *TokenNodeVisitor {
	visitor.CurrentNode = node
	visitor.StartPos = startAt
	visitor.EndPos = startAt
	visitor.LastMatches = nil
	return visitor
}

func (visitor *TokenNodeVisitor) Active() bool {
	return visitor.CurrentNode != nil
}

func (visitor *TokenNodeVisitor) Matches() []*TokenMatch {
	matches := make([]*TokenMatch, 0)
	if len(visitor.CurrentNode.Values()) > 0 {
		for _, token := range visitor.CurrentNode.Values() {
			matches = append(matches, &TokenMatch{Token: token, StartPos: visitor.StartPos, EndPos: visitor.EndPos})
		}
	}
	if len(matches) > 0 {
		return matches
	} else {
		return nil
	}
}

func (visitor *TokenNodeVisitor) Advance(runeValue rune, onMatch func([]*TokenMatch)) {
	if visitor.Active() {
		if matches := visitor.Matches(); matches != nil {
			visitor.LastMatches = matches
		}

		visitor.CurrentNode = visitor.CurrentNode.Next(runeValue)
		visitor.EndPos += 1

		if visitor.CurrentNode == nil {
			if visitor.LastMatches != nil {
				onMatch(visitor.LastMatches)
			}
		}
	}
}