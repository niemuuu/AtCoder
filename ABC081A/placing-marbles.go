/*
ABC081A - Placing Marbles
https://atcoder.jp/contests/abs/tasks/abc081_a
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

func main() {
	nums := strings.Split(readLine(), "")

	var result int
	for _, num := range nums {
		if num == "1" {
			result++
		}
	}

	fmt.Printf("%d\n", result)
}
