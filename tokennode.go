package tokensearch

import (
	"unicode/utf8"
	"errors"
)

type TokenNode struct {
	nextLetters		map[rune]*TokenNode
	matches			map[string]*Token
}

func NewTokenNode() (*TokenNode) {

	return &TokenNode{make(map[rune]*TokenNode), make(map[string]*Token)}
}

func (node *TokenNode) Insert(match *Token) (int, error) {

	return node.recurseInsert(match.Key(), 0, match)
}

func (node *TokenNode) Remove(match *Token) (int, error) {

	return node.recurseRemove(match.Key(), 0, match)
}

func (node *TokenNode) Find(token string) []*Token {

	return node.recurseFind(NormalizeString(token), 0)
}

func (node *TokenNode) Next(runeValue rune) *TokenNode {

	return node.nextLetters[runeValue]
}

func (node *TokenNode) Values() []*Token {
	arr := make([]*Token, 0, len(node.matches))
	for _, value := range node.matches {
		arr = append(arr, value)
	}
	return arr
}

func (node *TokenNode) AllValues(max int) []*Token {
	arr := make([]*Token, 0, len(node.matches))
	arr = append(arr, node.Values()...)
	for _, childNode := range node.nextLetters {
		if len(arr) > max {
			break
		}
		arr = append(arr, childNode.AllValues(max)...)
	}
	return arr
}

func (node *TokenNode) recurseInsert(token string, index int, match *Token) (int, error) {

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

func (node *TokenNode) recurseRemove(token string, index int, match *Token) (int, error) {

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
		nextNode = NewTokenNode()
		node.nextLetters[key] = nextNode
	}
	return nextNode, nil
}

func (node *TokenNode) appendMatch(match *Token) {

	if node.matches[match.Ident] == nil {
		node.matches[match.Ident] = match
	}
}

func (node *TokenNode) removeMatch(match *Token) {
	delete(node.matches, match.Ident)
}

func (node *TokenNode) recurseFind(token string, index int) []*Token {

	if index < len(token) {
		runeValue, width := utf8.DecodeRuneInString(token[index:])
		if width > 0 {

			if nextNode := node.Next(runeValue); nextNode != nil {

				// node found; recusively visit next node with character state and position
				return nextNode.recurseFind(token, index + width)
			} else {

				// reached end of search tree; return nil
				return nil
			}
		}

		// rune not found; return nil
		return nil
	}

	// reached end of search string; return any values found in the current node
	return node.Values()
}