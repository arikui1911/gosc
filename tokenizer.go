package gosc

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

type Token struct {
	Tag    TokenTag
	Value  string
	Lineno int
	Column int
}

type Tokenizer struct {
	src    io.Reader
	lineno int
	column int
	// runeBuf []rune
}

func NewTokenizer(src io.Reader) *Tokenizer {
	return &Tokenizer{src: src}
}

func (t *Tokenizer) scan() error {
	scanner := bufio.NewScanner(t.src)
	t.lineno = 1
	t.column = 1
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		t.lineno++
	}
	return scanner.Err()
}

var spacesRe = regexp.MustCompile(`\A\s+`)

func scanInitial(t *Tokenizer, src string) {
	switch {
	case spacesRe.MatchString(src):
	}
}

/*
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
	postRegionCommentSharpState
	symbolState
	stringState
	atomState
)

func (t *Tokenizer) NextToken() (token Token, retErr error) {
	state := initialState
	buf := []rune{}
	regionCommentNests := 0

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
				continue
			} else if c == ';' {
				state = commentState
				continue
			} else if c == '#' {
				state = postSharpState
				continue
			}
		case commentState:
			if c == '\n' {
				state = initialState
			}
		case postSharpState:
			if c == '|' {
				state = regionCommentState
				regionCommentNests++
				continue
			}
			buf = append(buf, '#')
			buf = append(buf, c)
			state = atomState
		case regionCommentState:
			if c == '#' {
				state = postRegionCommentSharpState
			}
		case postRegionCommentSharpState:
			if c == '|' {
				regionCommentNests--
				if regionCommentNests == 0 {
					state = initialState
					continue
				}
			}
			state = regionCommentState
		}
	}
}
*/
