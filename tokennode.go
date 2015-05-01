package tokensearch

import (
	"unicode/utf8"
	"errors"
)

type TokenNode struct {
	nextLetters		map[rune]*TokenNode
	matches			map[string]*TokenMatch
}

func NewTokenNode() (*TokenNode, error) {

	return &TokenNode{make(map[rune]*TokenNode), make(map[string]*TokenMatch)}, nil
}

func (node *TokenNode) Insert(match *TokenMatch) (int, error) {

	return node.recurseInsert(match.Key(), 0, match)
}

func (node *TokenNode) Remove(match *TokenMatch) (int, error) {

	return node.recurseRemove(match.Key(), 0, match)
}

func (node *TokenNode) Find(token string) ([]*TokenMatch, error) {

	return node.recurseFind(NormalizeKey(token), 0)
}

func (node *TokenNode) Values() []*TokenMatch {
	arr := make([]*TokenMatch, 0, len(node.matches))
	for _, value := range node.matches {
		arr = append(arr, value)
	}
	return arr
}

func (node *TokenNode) recurseInsert(token string, index int, match *TokenMatch) (int, error) {

	if index < len(token) {
		runeValue, width := utf8.DecodeRuneInString(token[index:])
		if width > 0 {
			nextNode, err := node.getOrCreateChild(runeValue)
			if err != nil {
				return index, errors.New("Unable to find or build node")
			}
			return nextNode.recurseInsert(token, index + width, match)
		}
		return index, errors.New("UTF-8 character width was 0")
	}

	node.appendMatch(match)
	return index, nil
}

func (node *TokenNode) recurseRemove(token string, index int, match *TokenMatch) (int, error) {

	if index < len(token) {
		runeValue, width := utf8.DecodeRuneInString(token[index:])
		if width > 0 {
			nextNode := node.nextLetters[runeValue]
			if nextNode == nil {
				return index, errors.New("Unable to find node with match")
			}
			return nextNode.recurseRemove(token, index + width, match)
		}
	}

	node.removeMatch(match)
	return index, nil
}

func (node *TokenNode) getOrCreateChild(key rune) (*TokenNode, error) {

	nextNode := node.nextLetters[key]
	if nextNode == nil {
		nextNode, _ = NewTokenNode()
		node.nextLetters[key] = nextNode
	}
	return nextNode, nil
}

func (node *TokenNode) appendMatch(match *TokenMatch) {

	if node.matches[match.Ident] == nil {
		node.matches[match.Ident] = match
	}
}

func (node *TokenNode) removeMatch(match *TokenMatch) {
	delete(node.matches, match.Ident)
}

func (node *TokenNode) recurseFind(token string, index int) ([]*TokenMatch, error) {

	if index < len(token) {
		runeValue, width := utf8.DecodeRuneInString(token[index:])
		if width > 0 {
			nextNode := node.nextLetters[runeValue]
			if nextNode == nil {
				return nil, nil
			}
			return nextNode.recurseFind(token, index + width)
		}
		return nil, errors.New("UTF-8 character was 0")
	}
	return node.Values(), nil
}