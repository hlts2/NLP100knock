// 04. 元素記号
// "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."という文を単語に分解し，1, 5, 6, 7, 8, 9, 15, 16, 19番目の単語は先頭の1文字，それ以外の単語は先頭に2文字を取り出し，取り出した文字列から単語の位置（先頭から何番目の単語か）への連想配列（辞書型もしくはマップ型）を作成せよ．

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."

	tokens := strings.Split(strings.Replace(s, ".", "", -1), " ")

	m := make(map[string]int, len(tokens))

	for i, token := range tokens {
		runes := []rune(token)
		switch i + 1 {
		case 1, 5, 6, 7, 8, 9, 15, 19:
			m[string(runes[0])] = i + 1
		default:
			m[string(runes[0])+string(runes[1])] = i + 1
		}
	}

	fmt.Println(m)
}
