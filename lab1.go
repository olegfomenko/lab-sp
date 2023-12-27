package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

type pair struct {
	length int
	vowels int
}

// Calculate pairs order(len;vowels) by quotient
func getPairs() []pair {
	var pairs []pair

	// i - length
	// j - vowels
	for i := 1; i <= 30; i++ {
		for j := 0; j <= i; j++ {
			pairs = append(pairs, pair{i, j})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		quotientI := float64(pairs[i].vowels) / float64(pairs[i].length)
		quotientJ := float64(pairs[j].vowels) / float64(pairs[j].length)

		return quotientI <= quotientJ
	})

	return pairs
}

var vowels = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'y': {},
}

func calcPair(str string) (int, int) {
	vowelCount := 0

	for _, c := range str {
		if _, ok := vowels[c]; ok {
			vowelCount++
		}
	}

	return len(str), vowelCount
}

// Var 11
func RunLab1(in, out string) {
	const wordFileFormat = "%s/word_%d_%d"
	const existFileFormat = "%s/exist_%s"

	f, err := os.Open(in)
	if err != nil {
		panic(err)
	}

	finish := func(word string) {
		if word == "" {
			return
		}

		l, v := calcPair(word)
		file, err := os.OpenFile(fmt.Sprintf(wordFileFormat, out, l, v), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		if _, err := file.WriteString(word); err != nil {
			panic(err)
		}

		if _, err := file.WriteString("\n"); err != nil {
			panic(err)
		}

		if err := file.Close(); err != nil {
			panic(err)
		}
	}

	var current string

	flush := func() {
		finish(current)
		current = ""
	}

	readWords := func(data []byte, n int) {
		str := string(data)
		for i := 0; i < n; i++ {
			if str[i] >= 'a' && str[i] <= 'z' || str[i] >= 'A' && str[i] <= 'Z' {
				current += string(str[i])
				continue
			}

			flush()
		}
	}

	for {
		binary := make([]byte, 256)
		n, err := f.Read(binary)

		if err != nil {
			if err == io.EOF {
				readWords(binary, n)
				flush()
				break
			}

			panic(err)
		}

		readWords(binary, n)
	}

	result, err := os.Create(fmt.Sprintf("%s/%s", out, "result.txt"))
	if err != nil {
		panic(err)
	}

	pairs := getPairs()
	for _, p := range pairs {
		file, err := os.Open(fmt.Sprintf(wordFileFormat, out, p.length, p.vowels))
		if err != nil {
			// File does not exist => no such words
			continue
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			word := scanner.Text()

			if _, err := os.Stat(fmt.Sprintf(existFileFormat, out, word)); !errors.Is(err, os.ErrNotExist) {
				// file exist => word already added
				continue
			}

			if _, err := result.WriteString(word); err != nil {
				panic(err)
			}

			if _, err := result.WriteString("\n"); err != nil {
				panic(err)
			}

			wf, err := os.Create(fmt.Sprintf(existFileFormat, out, word))
			if err != nil {
				panic(err)
			}

			if err := wf.Close(); err != nil {
				panic(err)
			}
		}

		if err := file.Close(); err != nil {
			panic(err)
		}
	}

	if err := result.Close(); err != nil {
		panic(err)
	}

	if err := f.Close(); err != nil {
		panic(err)
	}

}
