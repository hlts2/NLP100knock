// 02. 「パトカー」＋「タクシー」＝「パタトクカシーー」
// 「パトカー」＋「タクシー」の文字を先頭から交互に連結して文字列「パタトクカシーー」を得よ．

package main

import "fmt"

func alternatelyConcat(runes1 []rune, runes2 []rune) string {
	var s string
	for i := 0; i < len(runes1); i++ {
		s += string(runes1[i]) + string(runes2[i])
	}
	return s
}

func main() {
	s1 := "パトカー"
	s2 := "タクシー"
	fmt.Println(alternatelyConcat([]rune(s1), []rune(s2)))
}
