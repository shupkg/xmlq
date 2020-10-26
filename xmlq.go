package main

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/xmlpath.v2"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: xmlq <xpath> <filename>")
		os.Exit(1)
	}

	p := os.Args[1]
	fn := os.Args[2]

	f, err := os.Open(fn)
	fatal(err)
	defer f.Close()

	xpath, err := xmlpath.Compile(p)
	fatal(err)

	node, err := xmlpath.Parse(f)
	fatal(err)

	s, ok := xpath.String(node)
	if !ok {
		fatal(fmt.Errorf("XPATH: %s not exist", p))
	}

	fmt.Println(remoteSpace(s))
}

func fatal(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(0)
	}
}

func remoteSpace(s string) string {
	if s == "" {
		return ""
	}
	ss := strings.Split(s, "\n")
	sr := make([]string, 0, len(ss))
	for _, s := range ss {
		s = strings.TrimSpace(s)
		if s != "" {
			sr = append(sr, s)
		}
	}
	return strings.Join(sr, "\n")
}
