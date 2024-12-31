package parser

import (
	"strings"

	"golang.org/x/net/html"
)

type Parser struct{}

func (p *Parser) Parse(htmlContent string) (*html.Node, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, err
	}
	return doc, nil
}
