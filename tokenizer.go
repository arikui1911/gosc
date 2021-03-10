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

func NewTokenizer(src string) *Tokenizer {
	return &Tokenizer{
		src:     src,
		tokenCh: make(chan Token),
		errCh:   make(chan error),
	}
}

func (t *Tokenizer) NextToken() (tok Token, err error) {
	select {
	case tok = <-t.tokenCh:
	case err = <-t.errCh:
	}
	return
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

func scan(t *Tokenizer) error {
	s := strscan.New(t.src)
	for !s.IsEOF() {
		p := s.Pos()
		switch {
		case s.Scan(spacesRe):
			continue
		case s.Scan(charsRe):
			emit(t, s, charsTable[s.Matched()], s.Matched(), p)
		case s.Scan(atomRe):
			emit(t, s, AtomToken, s.Matched(), p)
		default:
			panic("must not happen")
		}
	}
	return nil
}
