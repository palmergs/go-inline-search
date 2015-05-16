package tokensearch

import (
	"unicode"
)

type Token struct {
	Ident		string	`json:"id"`
	name		string
	Display		string	`json:"label"`
	Category	string	`json:"category"`
}

func NewToken(ident, display, category string) *Token {
	token := Token{Ident: ident, Display: display, Category: category}
	token.InitKey()
	return &token
}

func (token *Token) InitKey() string {
	token.name = NormalizeString(token.Display)
	return token.name
}

func (token *Token) EqualIdent(other *Token) bool {
	return token.Ident == other.Ident;
}

func (token *Token) EqualCategory(other *Token) bool {
	return token.Category == other.Category;
}

func (token *Token) Key() string {
	return token.name
}

func NormalizeString(str string) string {
	normalizedStr := make([]rune, 0)
	lastWasChar := true
	charCount := 0
	for _, runeValue := range str {
		newRune, currIsChar := NormalizeRune(runeValue)
		if currIsChar {
			if !lastWasChar && charCount > 0 {
				normalizedStr = append(normalizedStr, ' ')
			}
			normalizedStr = append(normalizedStr, newRune)
			charCount++
		}
		lastWasChar = currIsChar
	}
	return string(normalizedStr)
}

func NormalizeRune(rn rune) (rune, bool) {
	if unicode.IsPrint(rn) {

		// whitespace, dashes and underscores are normalized to a single space
		if unicode.IsSpace(rn) || rn == '-' || rn == '_' {
			return ' ', false
		}

		// letters are normalized to lowercase
		if unicode.IsLetter(rn) {
			return unicode.ToLower(rn), true
		}

		// digits are passed through without modification
		if unicode.IsDigit(rn) {
			return rn, true
		}

		// return this rune as a character
		return rn, true
	}

	// a non-printing character is returned as a space with non-character state set
	return ' ', false
}
