package main

import (
	"flag"
	"fmt"
	"github.com/songjiayang/gocntool/gocn"
	"io/ioutil"
	"strings"
)

var (
	checkFile string
)

func init() {
	flag.StringVar(&checkFile, "file", "gocn.md", "gocntool -file ./gocn.md")
}

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(checkFile)
	if err != nil {
		panic(err)
	}

	checker := gocn.NewChecker(content)

	conflicts := checker.Check()

	// Newline
	fmt.Println()

	if len(conflicts) == 0 {
		fmt.Println("Congratulations！check passed.")
		return
	}

	fmt.Printf("There are %d conflicts:\n%s\n", len(conflicts), strings.Join(conflicts, "\n"))
}
