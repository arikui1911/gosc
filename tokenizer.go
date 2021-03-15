package gosc

import (
	"regexp"

	"github.com/arikui1911/strscan"
)

type Token struct {
	Tag         TokenTag
	Value       string
	FirstLineno int
	FirstColumn int
	LastLineno  int
	LastColumn  int
}

type Tokenizer struct {
	src     string
	tokenCh chan Token
	errCh   chan error
}

func NewTokenizer(src string) (t *Tokenizer) {
	t = &Tokenizer{
		src:     src,
		tokenCh: make(chan Token),
		errCh:   make(chan error),
	}
	go scanLoop(t)
	return
}

func (t *Tokenizer) NextToken() (tok Token, err error) {
	select {
	case tok = <-t.tokenCh:
	case err = <-t.errCh:
	}
	return
}

type scanFn func(*Tokenizer, *strscan.StringScanner) scanFn

// scanLoop nerver ends. DO NOT call without goroutine.
func scanLoop(t *Tokenizer) {
	fn := scanDefault
	s := strscan.New(t.src)
	for !s.IsEOF() {
		fn = fn(t, s)
	}
	for {
		emit(t, s, EOFToken, "", s.Pos())
	}
}

func emit(t *Tokenizer, s *strscan.StringScanner, tag TokenTag, val string, begPos int) {
	fl, fc := s.LinenoAndColumn(begPos)
	ll, lc := s.LinenoAndColumn(s.Pos() - 1)
	t.tokenCh <- Token{tag, val, fl, fc, ll, lc}
}

var spacesRe = regexp.MustCompile(`\s+`)
var charsRe = regexp.MustCompile(`[().']`)
var atomRe = regexp.MustCompile(`[^\s;"()']+`)

var charsTable = map[string]TokenTag{
	"(": LeftParanToken,
	")": RightParenToken,
	".": DotToken,
	"'": QuoteToken,
}

func scanDefault(t *Tokenizer, s *strscan.StringScanner) scanFn {
	p := s.Pos()
	switch {
	case s.Scan(spacesRe):
		// do nothing
	case s.Scan(charsRe):
		emit(t, s, charsTable[s.Matched()], s.Matched(), p)
	case s.Scan(atomRe):
		emit(t, s, AtomToken, s.Matched(), p)
	default:
		panic("must not happen")
	}
	return scanDefault
}
