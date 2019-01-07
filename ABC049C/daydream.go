/*
ABC049C - Daydream
https://atcoder.jp/contests/abs/tasks/abc049_c
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() string {
	if !sc.Scan() {
		return ""
	}
	return sc.Text()
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func main() {
	s := reverse(readLine())
	result := "NO"

	// "dream", "dreamer", "erase", "eraser" の逆
	consts := []string{"maerd", "remaerd", "esare", "resare"}

	for i := 0; i < len(s); {
		next := false
		for j := 0; j < 4; j++ {
			v := s[i:]
			if strings.HasPrefix(v, consts[j]) {
				i += len(consts[j])
				next = true

				if i == len(s) {
					result = "YES"
				}
				break
			}
		}
		// prefixが見つからない場合、ループを終わらせる
		if next == false {
			i = len(s)
		}
	}

	fmt.Printf("%s\n", result)
}
