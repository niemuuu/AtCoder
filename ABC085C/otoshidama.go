/*
ABC085C - Otoshidama
https://atcoder.jp/contests/abs/tasks/abc085_c
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	stdin := strings.Split(readLine(), " ")
	n, _ := strconv.Atoi(stdin[0])
	y, _ := strconv.Atoi(stdin[1])

	ptn := []int{-1, -1, -1}

LOOP:
	for a := 0; a <= n; a++ {
		for b := 0; b <= n-a; b++ {
			c := n - (a + b)
			if c < 0 {
				continue
			}
			total := a*10000 + b*5000 + c*1000
			if total == y {
				ptn[0] = a
				ptn[1] = b
				ptn[2] = c
				break LOOP
			}
		}
	}

	fmt.Printf("%d", ptn[0])
	fmt.Printf("%s", " ")
	fmt.Printf("%d", ptn[1])
	fmt.Printf("%s", " ")
	fmt.Printf("%d\n", ptn[2])
}
