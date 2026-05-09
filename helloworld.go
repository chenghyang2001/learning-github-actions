package main

import (
	"fmt"
	"time"
)

// 印出問候語與當前時間後結束，作為 GitHub Actions CI 流程的最小驗證目標
func main() {
	fmt.Println("Hello, Peter Yang!")
	fmt.Println("現在時間：", time.Now().Format(time.DateTime))
}
