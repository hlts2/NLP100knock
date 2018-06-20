// 05. n-gram
// 与えられたシーケンス（文字列やリストなど）からn-gramを作る関数を作成せよ．この関数を用い，"I am an NLPer"という文から単語bi-gram，文字bi-gramを得よ．

package main

import (
	"fmt"
	"strings"
)

func wordNgram(tokens []string, n int) [][]string {
	ss := make([][]string, 0, len(tokens)-(n-1))

	for i := range tokens {
		if len(tokens) > i+n {
			ss = append(ss, tokens[i:i+n])
		} else {
			ss = append(ss, tokens[i:])
			break
		}
	}

	return ss
}

func slice2map(ss []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, v := range ss {
		m[v] = struct{}{}
	}
	return m
}

func map2Slice(m map[string]struct{}) []string {
	ss := make([]string, 0, len(m))
	for k := range m {
		ss = append(ss, k)
	}
	return map2Slice(slice2map(ss))
}

func runeSlice2StringSlice(runes []rune) []string {
	ss := make([]string, 0, len(runes))

	for _, r := range runes {
		ss = append(ss, string(r))
	}
	return ss
}

func charNgram(s string, n int) [][]string {
	runes := []rune(s)

	ss := make([][]string, 0, len(runes)-(n-1))
	for i := range runes {
		if len(runes) > i+n {
			ss = append(ss, runeSlice2StringSlice(runes[i:i+n]))
		} else {
			ss = append(ss, runeSlice2StringSlice(runes[i:]))
			break
		}
	}

	return ss
}

func main() {
	s := "I am an NLPer"

	tokens := strings.Split(s, " ")
	fmt.Println(wordNgram(tokens, 2))
	fmt.Println(charNgram(s, 2))
}
