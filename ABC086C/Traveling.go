/*
ABC086C - Traveling
https://atcoder.jp/contests/abs/tasks/arc089_a
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n, t1, t2 int
	var x1, y1, x2, y2 float64
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&t1, &x1, &y1)
		ax := int(math.Abs(x1))
		ay := int(math.Abs(y1))
		diff := int(ax) + int(ay)

		// 偶奇が一致しなければ終了
		if t1%2 != diff%2 {
			fmt.Println("No")
			return
		}

		if int((x1-x2)+(y1-y2)) > t1-t2 {
			fmt.Println("No")
			return
		}
		t2, x2, y2 = t1, x1, y1
	}

	fmt.Println("Yes")
}
