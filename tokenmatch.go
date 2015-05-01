package tokensearch

type TokenMatch struct {
	Ident		string
	name		string
	Display		string
	Category	string
}

func NewTokenMatch(ident, display, category string) *TokenMatch {
	return &TokenMatch{ident, NormalizeKey(display), display, category}
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

func NormalizeKey(str string) string {
	return str
}

func NormalizeRune(rn rune) rune {
	return rn
}
