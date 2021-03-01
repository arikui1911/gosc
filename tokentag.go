package gosc

//go:generate  stringer -type=TokenTag tokentag.go
type TokenTag int

const (
    InvalidToken TokenTag = iota
    EOFToken
    LeftParanToken
    RightParenToken
    QuoteToken
    DotToken
    SymbolToken
    StringToken
    AtomToken
)

