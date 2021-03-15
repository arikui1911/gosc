package gosc_test

import (
	"fmt"
	"testing"

	"github.com/arikui1911/gosc"
)

func TestTokenizer(t *testing.T) {
	src := `
hoge
hogeeee
foo
	`
	tr := gosc.NewTokenizer(src)
	tok, err := tr.NextToken()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tok)
	}
	tok, err = tr.NextToken()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tok)
	}
	tok, err = tr.NextToken()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tok)
	}
	tok, err = tr.NextToken()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tok)
	}
}
