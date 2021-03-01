package gosc

import (
	"bufio"
	"io"
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

func (t *Tokenizer) NextToken() (Token, error) {
    for {
        c, err := t.getc()
        if err == io.EOF {
            return Token{EOFToken, t.lineno, t.column}, nil
        }
        if err != nil {
            return 0, nil
        }
    }
}


