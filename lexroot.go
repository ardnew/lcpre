package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	var sep string

	flag.StringVar(&sep, "s", string(os.PathSeparator), "delimiter used to separate components")
	flag.Parse()

	fmt.Println(lexroot(sep, flag.Args()))
}

func lexroot(sep string, str []string) string {

	count := len(str)

	if count == 0 {
		return ""
	}

	base := str[0]

	if count == 1 {
		return base
	}

	for _, curr := range str[1:] {
		base = common(sep, base, curr)
	}

	return base
}

func common(sep string, s1, s2 string) string {

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

	if len(com) == 1 && com[0] == "" {
		return sep
	}

	return strings.Join(com, sep)
}
