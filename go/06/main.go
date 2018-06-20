// 06. 集合
// "paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ, XとYとして求め，XとYの和集合，積集合，差集合を求めよ．さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．

package main

import (
	"fmt"
)

func charNgram(s string, n int) []string {
	runes := []rune(s)

	ss := make([]string, 0, len(runes)-(n-1))
	for i := range runes {
		if len(runes) > i+n {
			ss = append(ss, string(runes[i:i+n]))
		} else {
			ss = append(ss, string(runes[i:]))
			break
		}
	}

	return set2Slice(slice2set(ss))
}

func slice2set(ss []string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, s := range ss {
		set[s] = struct{}{}
	}
	return set
}

func set2Slice(set map[string]struct{}) []string {
	ss := make([]string, 0, len(set))
	for k := range set {
		ss = append(ss, k)
	}
	return ss
}

func contains(ss []string, v string) bool {
	for _, s := range ss {
		if s == v {
			return true
		}
	}
	return false
}

func union(x, y []string) []string {
	return set2Slice(slice2set(append(x, y...)))
}

func product(x, y []string) []string {
	var l int
	if len(x) > len(y) {
		l = len(x)
	} else {
		l = len(y)
	}
	ans := make([]string, 0, l)

	for _, xv := range x {
		for _, yv := range y {
			if xv == yv {
				ans = append(ans, xv)
				break
			}
		}
	}
	return ans
}

func difference(x, y []string) []string {
	ans := make([]string, 0, len(x))

	for _, vx := range x {
		flg := false
		for _, vy := range y {
			if vx == vy {
				flg = true
				break
			}
		}

		if !flg {
			ans = append(ans, vx)
		}
	}
	return ans
}

func main() {
	s1 := "paraparaparadise"
	s2 := "paragraph"

	x := charNgram(s1, 2)
	y := charNgram(s2, 2)

	fmt.Println(x)

	// 和集合
	fmt.Println(union(x, y))

	// 積集合
	fmt.Println(product(x, y))

	// 差集合
	fmt.Println(difference(x, y)) // xからy要素を除く
	fmt.Println(difference(y, x)) // yからx要素を除く

	// 「se」を含むか、否か
	fmt.Println(contains(x, "se"))
	fmt.Println(contains(y, "se"))
}
