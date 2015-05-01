package tokensearch

import (
	"unicode/utf8"
	"errors"
)

type TokenNode struct {
	nextLetters		map[rune]*TokenNode
	matches			[]*TokenMatch
}

func (node *TokenNode) Insert(token string, index int, match *TokenMatch) (int, error) {
	if index < len(token) {
		runeValue, width := utf8.DecodeRuneInString(token[index:])
		if width > 0 {
			nextNode, err := node.buildOrCreateChild(runeValue)
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

func (node *TokenNode) buildOrCreateChild(key rune) (*TokenNode, error) {

	nextNode := node.nextLetters[key]
	if nextNode == nil {
		nextNode = new(TokenNode)
		nextNode.nextLetters = make(map[rune]*TokenNode)
		nextNode.matches = make([]*TokenMatch, 0)
		node.nextLetters[key] = nextNode
	}
	return nextNode, nil
}

func (node *TokenNode) Exists(match *TokenMatch) *TokenMatch {
	for _, existing := range node.matches {
		if existing.EqualIdent(match) {
			existing.name = match.name
			existing.category = match.category
			return existing
		}
	}
	return nil
}

func (node *TokenNode) Append(match *TokenMatch) {

	if node.Exists(match) == nil {
		newMatches := make([]*TokenMatch, len(node.matches) + 1)
		copy(newMatches, node.matches)
		newMatches[len(node.matches)] = match
		node.matches = newMatches
	}
}

func (node *TokenNode) Find(token string, index int) ([]*TokenMatch, error) {
	if index < len(token) {
		runeValue, width := utf8.DecodeRuneInString(token[index:])
		if width > 0 {
			nextNode := node.nextLetters[runeValue]
			if nextNode == nil {
				return nil, nil
			}
			return nextNode.Find(token, index + width)
		}
		return nil, errors.New("UTF-8 character was 0")
	}
	return node.matches, nil
}
