package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	var (
		sep  string
		lead bool
	)

	flag.StringVar(&sep, "s", "", "delimeter used to separate words")
	flag.BoolVar(&lead, "d", false, "print delimeter if it is the longest common prefix")
	flag.Parse()

	p, err := prefix(sep, lead, flag.Args())
	if nil != err {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(p)
}

func prefix(sep string, lead bool, str []string) (string, error) {

	count := len(str)

	if count == 0 {
		return "", errors.New("no input provided")
	}

	base := str[0]

	if count == 1 {
		return base, nil
	}

	for _, curr := range str[1:] {
		base = common(sep, lead, base, curr)
	}

	return base, nil
}

func common(sep string, lead bool, s1, s2 string) string {

	com := []string{}

	p1 := strings.Split(s1, sep)
	p2 := strings.Split(s2, sep)

	n1 := len(p1)
	n2 := len(p2)

	count := n1
	if count > n2 {
		count = n2
	}

	for i := 0; (i < count) && (p1[i] == p2[i]); i++ {
		com = append(com, p1[i])
	}

	if lead {
		if len(com) == 1 && com[0] == "" {
			return sep
		}
	}

	return strings.Join(com, sep)
}
