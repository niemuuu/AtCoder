/*
ABC085B - Kagami Mochi
https://atcoder.jp/contests/abs/tasks/abc085_b
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func readLine() string {
	if !sc.Scan() {
		return ""
	}
	return sc.Text()
}

func main() {
	n, _ := strconv.Atoi(readLine())
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		d, _ := strconv.Atoi(readLine())
		a = append(a, d)
	}

	m := make(map[int]struct{})

	b := make([]int, 0, n)

	for _, elem := range a {
		if _, ok := m[elem]; !ok {
			m[elem] = struct{}{}
			b = append(b, elem)
		}
	}

	fmt.Printf("%d\n", len(b))
}
