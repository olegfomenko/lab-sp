package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestLab1(t *testing.T) {
	RunLab1("/Users/olegfomenko/Documents/Work/go/src/github.com/olegfomenko/lab-sp/lab1.txt", "/Users/olegfomenko/Documents/Work/go/src/github.com/olegfomenko/lab-sp")
}

func TestLab3(t *testing.T) {
	r := regexp.MustCompile("([0-9]+)")

	text := "bbbb11asd 100s"
	fmt.Println(r.ReplaceAllString(text, " $1 "))

	line := "// test function"
	fmt.Println(regexp.MustCompile("^(\\/\\/.*)$").MatchString(line))
}
