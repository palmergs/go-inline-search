package tokensearch

type TokenMatch struct {
	ident		string
	name		string
	category	string
}

func (match *TokenMatch) EqualIdent(other *TokenMatch) bool {
	return match.ident == other.ident;
}

func (match *TokenMatch) EqualCategory(other *TokenMatch) bool {
	return match.category == other.category;
}
