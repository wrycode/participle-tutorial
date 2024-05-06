package main

import (
	"os"
	"github.com/alecthomas/repr"
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var (
	// MVP regex for IPA form lexing
	IPAFormLexer = lexer.MustSimple([]lexer.SimpleRule{
		{`Token`, `kw|k|ɪ|ŋ|\s`},
	})
	parser = participle.MustBuild[Document](participle.Lexer(IPAFormLexer))
)

type Document struct {
	Tokens []*Token `@@*`
}

type Token struct {
	Key string `@Token`
}

func main() {
	ini, err := parser.Parse("", os.Stdin)
	repr.Println(ini, repr.Indent("  "), repr.OmitEmpty(true))
	if err != nil {
		panic(err)
	}
}
