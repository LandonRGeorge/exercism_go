package wordy

import "testing"

func TestLexer(t *testing.T) {
	text := "What is plus 1 2?"
	testCases := []token{
		{what, "What"},
		{is, "is"},
		{operator, "plus"},
		{number, "1"},
		{number, "2"},
		{questionMark, "?"},
		{eof, ""},
	}
	l := newLexer(text)
	l.lex()
	if len(l.tokens) != len(testCases) {
		t.Errorf("want %d tokens, got %d", len(testCases), len(l.tokens))
	}
	for i, gotTok := range l.tokens {
		wantTok := testCases[i]
		if gotTok.tokenType != wantTok.tokenType {
			t.Errorf("[%d] want %v tokenType, got %v", i, wantTok.tokenType, gotTok.tokenType)
		}
		if gotTok.literal != wantTok.literal {
			t.Errorf("[%d] want %q literal, got %q", i, wantTok.literal, gotTok.literal)
		}
	}
}
