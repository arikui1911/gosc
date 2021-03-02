package gosc

import (
	"strings"
	"testing"
)

func TestTokenizer(t *testing.T) {
	src := `
hoge
hogeeee
foo
	`
	NewTokenizer(strings.NewReader(src)).scan()
}
