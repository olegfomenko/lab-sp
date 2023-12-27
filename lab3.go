package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/fatih/color"
)

// Var 11

type Lexem struct {
	content     string
	description string
}

func checkFullLineMatches(line string) (Lexem, bool) {
	if regexp.MustCompile("^(\".*\")$").MatchString(line) {
		return Lexem{line, "String literal"}, true
	}

	if regexp.MustCompile("^('.')$").MatchString(line) {
		return Lexem{line, "Symbol literal"}, true
	}

	if regexp.MustCompile("^(0x[0-9a-f]+)$").MatchString(line) {
		return Lexem{line, "Hex literal"}, true
	}

	if regexp.MustCompile("^([0-9]+(.[0-9]+)+)$").MatchString(line) {
		return Lexem{line, "Float literal"}, true
	}

	if regexp.MustCompile("^([0-9]+)$").MatchString(line) {
		return Lexem{line, "Integer literal"}, true
	}

	if regexp.MustCompile("^(true|false)$").MatchString(line) {
		return Lexem{line, "Boolean literal"}, true
	}

	if regexp.MustCompile("^(byte|int8|int16|int32|int64|uint8|uint16|uint32|uint64|uint|int|string|rune|boolfloat32|float64)$").MatchString(line) {
		return Lexem{line, "Golang type"}, true
	}

	if regexp.MustCompile("^(const|chan|break|defer|var|interface|case|go|func|map|continue|type|struct|default|import|else|package|fallthrough|for|goto|if|range|return|select|switch)$").MatchString(line) {
		return Lexem{line, "Reserved word"}, true
	}

	if regexp.MustCompile("^((\\()|(\\))|(\\{)|(\\}))$").MatchString(line) {
		return Lexem{line, "Bracket"}, true
	}

	if regexp.MustCompile("^(\\+|-|\\*|/|%|&|\\||\\^|>>|<<|==|!=|<|>|<=|>=|&&|\\|\\||!|<-|=|:=)$").MatchString(line) {
		return Lexem{line, "Operator"}, true
	}

	if regexp.MustCompile("^(,|\\.|:|;)$").MatchString(line) {
		return Lexem{line, "Other separators"}, true
	}

	return Lexem{}, false
}

func removeCommnents(text string) (string, []Lexem) {
	var lexems []Lexem

	multiLineCommentR := regexp.MustCompile("/\\*([^*]*(\\*[^/])*)*\\*/")
	comments := multiLineCommentR.FindAllStringIndex(text, -1)
	for _, match := range comments {
		start := match[0]
		finish := match[1]
		lexems = append(lexems, Lexem{text[start:finish], "Multi line comment"})
	}

	text = multiLineCommentR.ReplaceAllString(text, "")

	oneLineCommentR := regexp.MustCompile("\\/\\/.*\n")
	comments = oneLineCommentR.FindAllStringIndex(text, -1)
	for _, match := range comments {
		start := match[0]
		finish := match[1]

		lexems = append(lexems, Lexem{text[start:finish], "One line comment"})
	}

	text = oneLineCommentR.ReplaceAllString(text, "")

	return text, lexems
}

func RunLab3(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	text := string(data)

	var lexems []Lexem

	text, lexems = removeCommnents(text)
	text = regexp.MustCompile("\t").ReplaceAllString(text, "")
	lines := regexp.MustCompile("\n").Split(text, -1)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if lexem, ok := checkFullLineMatches(line); ok {
			lexems = append(lexems, lexem)
			continue
		}

		lineparts := regexp.MustCompile(" ").Split(line, -1)
		for _, part := range lineparts {
			if part == "" {
				continue
			}

			if lexem, ok := checkFullLineMatches(part); ok {
				lexems = append(lexems, lexem)
				continue
			}

			part = regexp.MustCompile("(\\+|-|\\*|/|%|&|\\||\\^|>>|<<|==|!=|<|>|<=|>=|&&|\\|\\||!|<-|=|:=|\\(|\\)|\\{|\\}|,|\\.|:|;)").ReplaceAllString(part, " $1 ")
			subparts := regexp.MustCompile(" ").Split(part, -1)

			for _, subpart := range subparts {
				if subpart == "" {
					continue
				}

				if lexem, ok := checkFullLineMatches(subpart); ok {
					lexems = append(lexems, lexem)
					continue
				}

				if regexp.MustCompile("^([a-zA-Z0-9_]+)$").MatchString(subpart) {
					lexems = append(lexems, Lexem{subpart, "Name"})
					continue
				}

				panic(fmt.Sprintf("Failed to classify:\nSubpart=%s;\nPart=%s;\nLine=%s;", subpart, part, line))
			}
		}
	}

	for _, lex := range lexems {
		c := color.New(color.FgCyan)
		c.Printf("%s ", lex.content)
		fmt.Printf("%s\n\n", lex.description)
	}
}
