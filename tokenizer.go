package gosc

import (
	"bufio"
	"io"
	"unicode"
)

type Token struct {
	Tag    TokenTag
	Value  string
	Lineno int
	Column int
}

type Tokenizer struct {
	src     *bufio.Reader
	lineno  int
	column  int
	runeBuf []rune
}

func (t *Tokenizer) getc() (c rune, err error) {
	if len(t.runeBuf) > 0 {
		c = t.runeBuf[len(t.runeBuf)-1]
		t.runeBuf = t.runeBuf[:len(t.runeBuf)-1]
	} else {
		c, _, err = t.src.ReadRune()
	}
	return
}

func (t *Tokenizer) ungetc(c rune) {
	t.runeBuf = append(t.runeBuf, c)
}

type tokenState int

const (
	invalidState tokenState = iota
	initialState
	commentState
	postSharpState
	regionCommentState
	symbolState
	stringState
)

func (t *Tokenizer) NextToken() (token Token, retErr error) {
	state := initialState

	for {
		c, err := t.getc()
		if err == io.EOF {
			token.Tag = EOFToken
			token.Lineno = t.lineno
			token.Column = t.column
			return
		}
		if err != nil {
			retErr = err
			return
		}

		switch state {
		case initialState:
			if unicode.IsSpace(c) {
				// do nothing
			} else if c == ';' {
				state = commentState
			} else if c == '#' {
				state = postSharpState
			}
		}
	}
}
