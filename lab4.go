package main

import (
	"fmt"
	"regexp"
)

type Grammar struct {
	Rules map[rune][]string

	first   map[rune][]rune
	visited map[rune]struct{}

	null map[rune]struct{}
}

func ReadGrammar() *Grammar {
	var n int
	fmt.Scan(&n)

	g := &Grammar{
		Rules: make(map[rune][]string),

		first:   make(map[rune][]rune),
		visited: make(map[rune]struct{}),

		null: make(map[rune]struct{}),
	}

	for i := 0; i < n; i++ {
		var A rune
		fmt.Scanf("%c", &A)

		var a string
		fmt.Scan(&a)

		if a == "<null>" {
			g.null[A] = struct{}{}
			continue
		}

		g.Rules[A] = append(g.Rules[A], a)
	}
	return g
}

func isTerminal(a rune) bool {
	return regexp.MustCompile("[a-z]").MatchString(string(a))
}

func (g *Grammar) First(A rune) []rune {
	if res, ok := g.first[A]; ok {
		return res
	}

	if _, ok := g.visited[A]; ok {
		panic("Failed to construct First(A) value - infinite recursion.")
	}

	g.visited[A] = struct{}{}

	res := make(map[rune]struct{})

	for _, rule := range g.Rules[A] {
		for i := 0; i < len(rule); i++ {
			c := rune(rule[i])
			if isTerminal(c) {
				res[c] = struct{}{}
				break
			} else {
				for _, r := range g.First(c) {
					res[r] = struct{}{}
				}

				if _, ok := g.null[c]; !ok {
					break
				}
			}
		}

	}

	result := make([]rune, 0, len(res))
	for r := range res {
		result = append(result, r)
	}

	g.first[A] = result
	return result
}

func (g *Grammar) RouteSymbols(Key rune, a rune) string {
	for _, rule := range g.Rules[Key] {

		for i := 0; i < len(rule); i++ {
			c := rune(rule[i])
			if isTerminal(c) {
				if c == a {
					return rule
				}

				break
			}

			for _, r := range g.First(c) {
				if r == a {
					return rule
				}
			}

			if _, ok := g.null[c]; !ok {
				break
			}
		}

	}

	if _, ok := g.null[Key]; ok {
		fmt.Println("Using null rule for key", string(Key))
		return ""
	}

	panic(fmt.Sprintf("Failed to search route symbols for rune %s. Searching symbol was: %s", string(Key), string(a)))
}

func Analyze(g *Grammar, text string) {
	var stack = make([]rune, 10000000)
	stack[0] = 'S'
	var top = 0

	var index = 0
	for index < len(text) {
		if top < 0 {
			panic(fmt.Sprintf("Empty stack, current sybmol %s", string(text[index])))
		}

		fmt.Printf("Index = %d, stack size = %d, stack top = %s\n", index, top+1, string(stack[top]))

		next := rune(text[index])

		if isTerminal(stack[top]) {
			if stack[top] == next {
				top--
				index++
				continue
			}

			panic(fmt.Sprintf("Stack top symbol is terminal (%s) but does not corresponnds the waiting one (%s)", string(stack[top]), string(text[index])))
		}

		rule := g.RouteSymbols(stack[top], next)
		if rule == "" {
			fmt.Println("Used null rule for symbol", string(text[index]), index)
			top--
			continue
		}

		top--
		for i := len(rule) - 1; i >= 0; i-- {
			top++
			stack[top] = rune(rule[i])
		}
	}

	for top >= 0 {
		if _, ok := g.null[stack[top]]; ok {
			top--
		}
	}

	if top >= 0 {
		for i := 0; i <= top; i++ {
			fmt.Println(string(stack[i]))
		}
		panic("Stack is not empty after all sybmols processed")
	}

	fmt.Println("OK")
}

func RunLab4() {
	g := ReadGrammar()
	var text string
	fmt.Scan(&text)
	Analyze(g, text)
}
