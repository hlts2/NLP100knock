// 03. 円周率
// "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."という文を単語に分解し，各単語の（アルファベットの）文字数を先頭から出現順に並べたリストを作成せよ．

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	tokens := strings.Split(strings.Replace(strings.Trim(s, "."), ",", "", -1), " ")

	counts := make([]int, 0, len(tokens))
	for _, token := range tokens {
		counts = append(counts, utf8.RuneCountInString(token))
	}

	fmt.Println(counts)
}
