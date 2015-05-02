package tokensearch

import (
	"unicode"
)

type TokenMatch struct {
	Ident		string
	name		string
	Display		string
	Category	string
}

func NewTokenMatch(ident, display, category string) *TokenMatch {
	return &TokenMatch{ident, NormalizeString(display), display, category}
}

func (match *TokenMatch) EqualIdent(other *TokenMatch) bool {
	return match.Ident == other.Ident;
}

func (match *TokenMatch) EqualCategory(other *TokenMatch) bool {
	return match.Category == other.Category;
}

func (match *TokenMatch) Key() string {
	return match.name
}

func NormalizeString(str string) string {
	normalizedStr := make([]rune, 0)
	whitespace := 0
	charSeen := false
	for _, runeValue := range str {
		newRune, isChar := mapRuneWithState(runeValue)
		if isChar {
			if whitespace > 0 {
				if charSeen {
					normalizedStr = append(normalizedStr, ' ')
				}
				whitespace = 0
			}
			normalizedStr = append(normalizedStr, newRune)
			charSeen = true
		} else {
			whitespace = whitespace + 1
		}
	}
	return string(normalizedStr)
}

func NormalizeRune(rn rune) rune {
	newRune, _ := mapRuneWithState(rn)
	return newRune
}

func mapRuneWithState(rn rune) (rune, bool) {
	if unicode.IsPrint(rn) {
		if unicode.IsSpace(rn) || rn == '-' || rn == '_' {
			return ' ', false
		}
		if unicode.IsLetter(rn) {
			return unicode.ToLower(rn), true
		}
		return rn, true
	}
	return ' ', false
}
