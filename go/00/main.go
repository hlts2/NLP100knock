// 00. 文字列の逆順
// 文字列"stressed"の文字を逆に（末尾から先頭に向かって）並べた文字列を得よ．

package main

import "fmt"

func reverse(s string) string {
	r := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		r[i], r[len(s)-1-i] = r[len(s)-1-i], r[i]
	}

	return string(r)
}

func main() {
	s := "stressed"
	fmt.Println(reverse(s) == "desserts")
}
