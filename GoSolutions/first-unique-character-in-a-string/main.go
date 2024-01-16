package main

import "fmt"

func firstUniqChar(s string) int {
	m := [26]int{}
	for i := range s {
		m[s[i]-'a']++
	}
	for i := range s {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
