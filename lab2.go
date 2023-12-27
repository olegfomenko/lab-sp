package main

import (
	"fmt"
)

type EdgeKey struct {
	State     int
	Character rune
}

type Graph struct {
	Edges map[EdgeKey][]int
}

func (g *Graph) Reverse() *Graph {
	reversed := make(map[EdgeKey][]int)
	for key, to := range g.Edges {
		for _, v := range to {
			rkey := EdgeKey{v, key.Character}
			reversed[rkey] = append(reversed[rkey], key.State)
		}
	}
	return &Graph{Edges: reversed}
}

// Execute Returns empty result if that word prefix can not be recognized by automaton
func (g *Graph) Execute(s0 int, s string) []int {
	var answer []int
	if len(s) == 0 {
		return []int{s0}
	}

	ekey := EdgeKey{s0, rune(s[0])}
	for _, v := range g.Edges[ekey] {
		answer = append(answer, g.Execute(v, s[1:])...)
	}

	return answer
}

func (g *Graph) Search(from, to int, alphabet map[rune]struct{}) (string, bool) {
	visited := make(map[int]struct{})
	return g.search(&visited, alphabet, from, to)
}

func (g *Graph) search(visited *map[int]struct{}, alphabet map[rune]struct{}, current, finish int) (string, bool) {
	if _, ok := (*visited)[current]; ok {
		return "", false
	}

	(*visited)[current] = struct{}{}

	if current == finish {
		return "", true
	}

	for c := range alphabet {
		ekey := EdgeKey{current, c}
		for _, v := range g.Edges[ekey] {
			if res, ok := g.search(visited, alphabet, v, finish); ok {
				return string(c) + res, true
			}

		}
	}

	return "", false
}

func ReadAutomaton() (g *Graph, s0 int, final map[int]struct{}, alphabet map[rune]struct{}) {
	final = make(map[int]struct{})
	alphabet = make(map[rune]struct{})
	g = &Graph{make(map[EdgeKey][]int)}

	var asize, ssize int
	fmt.Scan(&asize, &ssize)

	fmt.Scan(&s0)

	var fsize int
	fmt.Scan(&fsize)

	for i := 0; i < fsize; i++ {
		var fi int
		fmt.Scan(&fi)
		final[fi] = struct{}{}
	}

	var rsize int
	fmt.Scan(&rsize)

	for i := 0; i < rsize; i++ {
		var s, s1 int
		var c rune

		fmt.Scanf("%d %c %d", &s, &c, &s1)

		ekey := EdgeKey{s, rune(c)}
		g.Edges[ekey] = append(g.Edges[ekey], s1)
		alphabet[rune(c)] = struct{}{}
	}

	return g, s0, final, alphabet
}

// Var 6
func RunLab2() {
	var w1, w2 string
	fmt.Scan(&w1, &w2)

	g, s0, final, alphabet := ReadAutomaton()

	start := g.Execute(s0, w1)
	if len(start) == 0 {
		panic(fmt.Sprintf("Word w1=%s is not acceptable by automaton as prefix", w1))
	}

	rg := g.Reverse()
	w2reverse := reverse(w2)

	var finish []int
	for s := range final {
		finish = append(finish, rg.Execute(s, w2reverse)...)
	}

	if len(finish) == 0 {
		panic(fmt.Sprintf("Word w2=%s is not acceptable by automaton as sufix", w2reverse))
	}

	for _, s1 := range start {
		for _, s2 := range finish {
			str, ok := g.Search(s1, s2, alphabet)
			if ok {
				fmt.Println("w0 =", str)
				return
			}
		}
	}

	panic("Failed.")
}

func reverse(s string) (res string) {
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return
}
