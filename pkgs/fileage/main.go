package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// フラグの定義
	var seconds int64
	flag.Int64Var(&seconds, "s", 86400, "経過時間（秒単位）")
	flag.Int64Var(&seconds, "seconds", 86400, "経過時間（秒単位）")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("ファイルが指定された秒数以内に更新されているかを確認します")
		fmt.Println("使用方法: go run check_file_age.go [-s 秒数] <ファイルパス>")
		fmt.Println("例: go run check_file_age.go -s 3600 /path/to/file  # 1時間")
		os.Exit(2)
	}

	filePath := args[0]

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Printf("エラー: ファイル '%s' にアクセスできません: %v\n", filePath, err)
		os.Exit(2)
	}

	modTime := fileInfo.ModTime()
	duration := time.Since(modTime)

	if duration.Seconds() > float64(seconds) {
		fmt.Printf("ファイルは%d秒以上前に更新されています（経過時間: %.0f秒）\n",
			seconds, duration.Seconds())
		os.Exit(1)
	}

	fmt.Printf("ファイルは%d秒以内に更新されています（経過時間: %.0f秒）\n",
		seconds, duration.Seconds())
	os.Exit(0)
}
