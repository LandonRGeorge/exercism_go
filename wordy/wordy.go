package wordy

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrTokenType = errors.New("encountered unexpected tokenType")
)

type token struct {
	tokenType tokenType
	literal   string
}

type tokenType int

const (
	number tokenType = iota + 1
	operator
	what
	is
	by
	unknown
	questionMark
	eof
)

var wordMap = map[string]tokenType{
	"plus":       operator,
	"multiplied": operator,
	"divided":    operator,
	"minus":      operator,
	"what":       what,
	"is":         is,
	"by":         by,
}

func lookupWordType(text string) tokenType {
	text = strings.ToLower(text)
	oper, ok := wordMap[text]
	if !ok {
		return unknown
	}
	return oper
}

type lexer struct {
	text         string
	position     int
	readPosition int
	ch           byte
	tokens       []token
}

func newLexer(text string) *lexer {
	l := &lexer{
		text: text,
	}
	l.next()
	return l
}

func (l *lexer) lex() {
	for {
		switch {
		case l.ch == 0:
			tok := token{
				tokenType: eof,
				literal:   "",
			}
			l.tokens = append(l.tokens, tok)
			return
		case l.ch == ' ':
			l.next()
			continue
		case l.ch == '?':
			tok := token{
				tokenType: questionMark,
				literal:   string(l.ch),
			}
			l.tokens = append(l.tokens, tok)
			l.next()
		case isDigit(l.ch):
			l.handleNumber()
		case isLetter(l.ch):
			l.handleWord()
		default:
			tok := token{
				tokenType: unknown,
				literal:   string(l.ch),
			}
			l.tokens = append(l.tokens, tok)
			l.next()
		}
	}
}

func (l *lexer) isAtEnd() bool {
	return l.readPosition >= len(l.text)
}

func (l *lexer) next() {
	if l.isAtEnd() {
		l.ch = 0
	} else {
		l.ch = l.text[l.readPosition]
		l.position = l.readPosition
		l.readPosition++

	}
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9') || ch == '-'
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func (l *lexer) handleNumber() {
	start := l.position
	for isDigit(l.ch) {
		l.next()
	}
	tok := token{
		tokenType: number,
		literal:   l.text[start:l.position],
	}
	l.tokens = append(l.tokens, tok)
}

func (l *lexer) handleWord() {
	start := l.position
	for isLetter(l.ch) {
		l.next()
	}
	literal := l.text[start:l.position]
	wordType := lookupWordType(literal)
	tok := token{
		tokenType: wordType,
		literal:   literal,
	}
	l.tokens = append(l.tokens, tok)
}

type parser struct {
	l                 *lexer
	curToken          token
	nextToken         token
	curTokenPosition  int
	nextTokenPosition int
	curValue          int
}

func newParser(l *lexer) *parser {
	l.lex()
	p := &parser{
		l: l,
	}
	p.advance()
	p.advance()
	return p
}

func (p *parser) advance() {
	p.curToken = p.nextToken
	p.curTokenPosition = p.nextTokenPosition

	p.nextToken = p.l.tokens[p.nextTokenPosition]
	p.nextTokenPosition++
}

func (p *parser) curTokIs(t tokenType) bool {
	return p.curToken.tokenType == t
}

type operation func(a, b int) int

func (p *parser) getOperation() (operation, error) {
	var oper operation
	switch literal := p.curToken.literal; literal {
	case "plus":
		oper = func(a, b int) int { return a + b }
	case "minus":
		oper = func(a, b int) int { return a - b }
	case "multiplied":
		p.advance()
		if !p.curTokIs(by) {
			return nil, fmt.Errorf("multipled keyword is not followed by the word 'by'")
		}
		oper = func(a, b int) int { return a * b }
	case "divided":
		p.advance()
		if !p.curTokIs(by) {
			return nil, fmt.Errorf("divided keyword is not followed by the word 'by'")
		}
		oper = func(a, b int) int { return a / b }
	default:
		return nil, fmt.Errorf("%q is an unknown operation", literal)
	}
	return oper, nil
}

func convertToNumber(text string) (int, error) {
	val, err := strconv.Atoi(text)
	if err != nil {
		return 0, err
	}
	return val, nil
}

func (p *parser) parse() error {
	if !p.curTokIs(what) {
		return ErrTokenType
	}
	p.advance()
	if !p.curTokIs(is) {
		return ErrTokenType
	}
	p.advance()

	if !p.curTokIs(number) {
		return ErrTokenType
	}

	nbr, err := convertToNumber(p.curToken.literal)
	if err != nil {
		return err
	}

	p.curValue = nbr
	p.advance()

	for {
		if p.curTokIs(questionMark) || p.curTokIs(eof) {
			break
		}
		if !p.curTokIs(operator) {
			return ErrTokenType
		}
		oper, err := p.getOperation()
		if err != nil {
			return ErrTokenType
		}
		p.advance()
		if !p.curTokIs(number) {
			return ErrTokenType
		}
		nbr, err := convertToNumber(p.curToken.literal)
		if err != nil {
			return err
		}
		p.curValue = oper(p.curValue, nbr)
		p.advance()
	}
	return nil
}

func Answer(question string) (int, bool) {
	l := newLexer(question)
	p := newParser(l)
	if err := p.parse(); err != nil {
		return 0, false
	}
	return p.curValue, true
}
