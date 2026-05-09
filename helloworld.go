package main

import (
	"fmt"
	"time"

	// 嵌入 IANA 時區資料庫，確保在 GitHub Actions（無 tzdata 的 Linux 環境）也能正常使用 Asia/Taipei
	_ "time/tzdata"
)

// 印出問候語與當前台北時間後結束，作為 GitHub Actions CI 流程的最小驗證目標
func main() {
	// 取得台北時區；若載入失敗則 fallback 至 UTC 並印出警告
	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		fmt.Println("警告：無法載入 Asia/Taipei 時區，改用 UTC：", err)
		loc = time.UTC
	}

	fmt.Println("Hello, Peter Yang!")
	fmt.Println("現在時間：", time.Now().In(loc).Format(time.DateTime))
}
