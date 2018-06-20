// 06. 集合
// "paraparaparadise"と"paragraph"に含まれる文字bi-gramの集合を，それぞれ, XとYとして求め，XとYの和集合，積集合，差集合を求めよ．さらに，'se'というbi-gramがXおよびYに含まれるかどうかを調べよ．

package main

import "fmt"

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
	m := make(map[string]int, len(x)+len(y))
	for _, v := range x {
		m[v] = 0
	}

	for _, v := range y {
		m[v] = 0
	}

	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}
	return result
}

func intersection(x, y []string) []string {
	m := make(map[string]int, len(x)+len(y))

	for _, v := range x {
		m[v]++
	}

	for _, v := range y {
		m[v]++
	}

	var l int
	for _, v := range m {
		if v > 1 {
			l++
		}
	}

	result := make([]string, 0, l)
	for k, v := range m {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}

func setdifference(x, y []string) []string {
	m := make(map[string]int, len(x))

	for _, v := range x {
		m[v] = 0
	}

	for _, v := range y {
		delete(m, v)
	}

	result := make([]string, 0, len(m))
	for v := range m {
		result = append(result, v)
	}

	return result
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
	fmt.Println(intersection(x, y))

	// 差集合
	fmt.Println(setdifference(x, y)) // xからy要素を除く
	fmt.Println(setdifference(y, x)) // yからx要素を除く

	// 「se」を含むか、否か
	fmt.Println(contains(x, "se"))
	fmt.Println(contains(y, "se"))
}
