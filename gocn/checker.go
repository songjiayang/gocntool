package gocn

import (
	"github.com/songjiayang/gocntool/github"
	"log"
	"strings"
	"time"
)

type Checker struct {
	parser *Parser
}

func NewChecker(content []byte) *Checker {
	return &Checker{
		parser: NewParser(content),
	}
}

func (c *Checker) Check() []string {
	links := c.parser.Parse().GetLinks(5)

	conflicts := make([]string, 0)

	for _, link := range links {
		if c.doCheck(link) {
			conflicts = append(conflicts, link)
		}

		// For GitHub RateLimit
		time.Sleep(1 * time.Second)
	}

	return conflicts
}

func (c *Checker) doCheck(link string) bool {
	log.Println("start check:", link)
	defer log.Println("after check:", link)

	keys := strings.Split(link, "/")
	keyLen := len(keys)

	kw := keys[keyLen-1]

	if len(kw) < 8 && keyLen > 1 {
		kw = keys[keyLen-2] + "/" + kw
	}

	isConflict, err := github.Search(kw)
	if err != nil {
		log.Fatalf("github.Search(%s): %v", kw, err)
	}

	return isConflict
}
