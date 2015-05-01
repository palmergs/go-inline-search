package tokensearch

import (
	"unicode/utf8"
	"errors"
)

type TokenNode struct {
	nextLetters		map[rune]*TokenNode
	matches			[]*TokenMatch
}

func (node *TokenNode) Insert(token string, index int, match *TokenNode) (int, error) {
	if index < len(value) {
		runeValue, width := utf8.DecodeRuneInString(token[index:])
		if width > 0 {
			nextNode, err := node.FindOrNew(runeValue)
			if err != nil {
				return index, errors.New("Unable to find or build node")
			}
			return nextNode.Insert(token, index + width, match)
		}
		return index, errors.New("UTF-8 character width was 0")
	}

	node.Append(match)
	return index, nil
}

func (node *TokenNode) FindOrNew(key rune) (*TokenNode, error) {
	if key == nil {
		return nil, errors.New("No rune given to create new key.")
	}

	nextNode = node.nextLetters[key]
	if nextNode == nil {
		nextNode = new(TokenNode)
		nextNode.nextLetters = make(map[rune]*TokenNode)
		nextNode.matches = make([]*TokenMatch)
		node.nextLetters[key] = nextNode
	}
	return nextNode, nil
}

func (node *TokenNode) Append(match *TokenNode) {
	newMatches := make([]*TokenNode, len(node.matches) + 1)
	copy(newMatches, node.matches)
	newMatches[len(node.matches)] = match
	node.matches = newMatches
}
