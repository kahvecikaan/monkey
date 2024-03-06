package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New() is a constructor function that returns a new lexer. It initializes the lexer by setting the input string and
// calling readChar() twice so both l.ch and l.readPosition are set properly. Why do we have make 2 readChar() calls?
// Because we need both l.ch and l.readPosition to be set before we can call NextToken() for the first time. The first
// call to readChar() sets both l.ch and l.readPosition, while the second one advances those fields to their correct
// values. After these two calls, we can call NextToken() and get the first token from our input string.

func New(input string) *Lexer {
	l := &Lexer{input: input} // create a new Lexer (a pointer to a Lexer) by passing in the input string
	l.readChar()              // sets l.ch and l.readPosition
	return l
}

func (l *Lexer) readChar() {
	// If we reach the end of the input, we set ch to 0, which is the ASCII code for the "NUL" character and has no
	// visible representation. We do this instead of returning an error or throwing an exception because we want our
	// lexer to always return a character. This way, our parser can always make progress in the input string and never
	// has to handle errors or exceptions.

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// Otherwise, we read the next character and advance our position in the input string.
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition // l.position always points where we last read
	l.readPosition += 1         // l.readPosition always points to the next character
}

// NextToken() is the heart of our lexer. It's responsible for both reading a character from the input and returning
// the next token. It's also responsible for advancing our two pointers l.position and l.readPosition.

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	// We skip over any whitespace characters by calling l.skipWhitespace().
	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch // save the current character
			l.readChar()
			literal := string(ch) + string(l.ch)                // create a new literal
			tok = token.Token{Type: token.EQ, Literal: literal} // create a new token
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}

	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch // save the current character
			l.readChar()
			literal := string(ch) + string(l.ch)                    // create a new literal
			tok = token.Token{Type: token.NOT_EQ, Literal: literal} // create a new token
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0: // 0 is the ASCII code for the "NUL" character and has no visible representation
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal) // check if the identifier is a keyword
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber() // readNumber() advances l.position and l.readPosition
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// readIdentifier() reads in an identifier and advances the lexer's position until it encounters a non-letter character.
// It assumes that the current character is a letter and reads until it encounters a non-letter character. It then
// returns the substring from l.position to l.readPosition. We use this function to read in keywords and identifiers.

func (l *Lexer) readIdentifier() string {
	position := l.position // save the current position in the input string
	for isLetter(l.ch) {   // read until we encounter a non-letter character
		l.readChar()
	}
	return l.input[position:l.position] // return the substring from position to l.position
}

func isLetter(ch byte) bool {
	// We only support ASCII characters for now. We can easily extend this to support Unicode characters by using
	// unicode.IsLetter() instead of our own isLetter() function.
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	// We only support ASCII characters for now. We can easily extend this to support Unicode characters by using
	// unicode.IsDigit() instead of our own isDigit() function.
	return '0' <= ch && ch <= '9'
}

// Note that we ignore floating point numbers, hexadecimal numbers, and so on. We only support integers for now.

func (l *Lexer) readNumber() string {
	position := l.position // save the current position in the input string
	for isDigit(l.ch) {    // read until we encounter a non-digit character
		l.readChar()
	}
	if l.ch == '.' {
		l.readChar()
		for isDigit(l.ch) {
			l.readChar()
		}
	}
	return l.input[position:l.position] // return the substring from position to l.position
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' { // skip whitespace characters
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) { // if we reach the end of the input
		return 0
	} else {
		return l.input[l.readPosition] // return the next character
	}
}
