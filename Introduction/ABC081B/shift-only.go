/*
ABC081B - Shift only
https://atcoder.jp/contests/abs/tasks/abc081_b
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

var cnt int

func getDividingCount(n int, nums []int) int {
	halfNums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		num := nums[i]

		if num%2 != 0 {
			return cnt
		}

		halfNums = append(halfNums, num/2)
	}
	cnt++
	return getDividingCount(n, halfNums)
}

func main() {
	n, _ := strconv.Atoi(readLine())
	numStrs := strings.Split(readLine(), " ")

	nums := make([]int, 0, n)
	for _, s := range numStrs {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}

	cnt := getDividingCount(n, nums)
	fmt.Printf("%d\n", cnt)
}
