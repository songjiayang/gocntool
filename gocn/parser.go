package gocn

import (
	"regexp"
	"strings"
)

var (
	urlRegexp = regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
)

type Parser struct {
	content []byte
	links   []string
}

func NewParser(content []byte) *Parser {
	return &Parser{
		content: content,
	}
}

func (p *Parser) Parse() *Parser {
	foundItems := urlRegexp.FindAll(p.content, -1)

	links := make([]string, len(foundItems))

	for index, link := range foundItems {
		links[index] = strings.TrimSuffix(string(link), "/")
	}

	p.links = links

	return p
}

func (p *Parser) GetLinks(limit int) []string {
	if len(p.links) < limit {
		return p.links
	}

	return p.links[0:limit]
}
