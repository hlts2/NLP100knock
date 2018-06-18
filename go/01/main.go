// 01. 「パタトクカシーー」
// 「パタトクカシーー」という文字列の1,3,5,7文字目を取り出して連結した文字列を得よ．

package main

import "fmt"

func main() {
	s := "パタトクカシーー"

	runes := []rune(s)
	fmt.Println(string(runes[0]) + string(runes[2]) + string(runes[4]) + string(runes[6]))
}
