/*
ABC087B - Coins
https://atcoder.jp/contests/abs/tasks/abc087_b
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

func calcPtn(ptn500, ptn100, ptn50, target int) int {
	var matchPtn int

	if target < 50 {
		return 0
	} else if target < 100 {
		ptn500 = 0
		ptn100 = 0
	} else if target < 500 {
		ptn500 = 0
	}

	for i := 0; i <= ptn50; i++ {
		for j := 0; j <= ptn100; j++ {
			for k := 0; k <= ptn500; k++ {
				amount := 50*i + 100*j + 500*k
				if amount == target {
					matchPtn++
				}
			}
		}
	}
	return matchPtn
}

func main() {
	n500, _ := strconv.Atoi(readLine())
	n100, _ := strconv.Atoi(readLine())
	n50, _ := strconv.Atoi(readLine())
	targetAmount, _ := strconv.Atoi(readLine())

	numOfPtn := calcPtn(n500, n100, n50, targetAmount)
	fmt.Printf("%d\n", numOfPtn)
}
