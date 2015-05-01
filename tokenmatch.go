package tokensearch

import (
	"unicode/utf8"
	"errors"
)

type TokenMatch struct {
	ident		string
	name		string
	category	string
}

func (match *TokenMatch) EqualIdent(other *TokenMatch) bool {
	if match.ident == nil || other.ident == nil {
		return false;
	}
	return match.ident == other.ident;
}

func (match *TokenMatch) EqualCategory(other *TokenMatch) bool {
	if match.category == nil || other.category == nil {
		return false;
	}
	return match.category == other.category;
}