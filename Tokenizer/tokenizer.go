package Tokenizer

// Token: struct to represent a token
type Token struct {
	tokenType string
	value     string
}

// Tokenizer: struct to tokenize input
type Tokenizer struct {
	LanguageSet map[string]Token
}

// The Tokenizer is a singleton
var T Tokenizer

func init() {
	T = Tokenizer{}
}

// Tokenize: tokenizes input
func (t *Tokenizer) Tokenize(input string) []Token {
	var tokens []Token
	return tokens
}
