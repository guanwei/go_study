// dup1 输出标准输入中出现次数大于1的行，前面是次数
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "end" {
			break
		}
		counts[input.Text()]++
	}
	// 注意：忽略input.Err()中可能的错误
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}