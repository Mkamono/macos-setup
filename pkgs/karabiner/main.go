package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"karabiner/core"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: karabiner <command> [arguments]")
		fmt.Println("Commands:")
		fmt.Println("  format <input.json>    - JSONファイルを読み込んで整形する")
		fmt.Println("  generate               - 設定からJSONを生成する")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "format":
		formatCmd := flag.NewFlagSet("format", flag.ExitOnError)
		formatCmd.Parse(os.Args[2:])
		if formatCmd.NArg() < 1 {
			fmt.Println("入力ファイルを指定してください")
			os.Exit(1)
		}
		formatJSON(formatCmd.Arg(0))

	case "generate":
		generateCmd := flag.NewFlagSet("generate", flag.ExitOnError)
		generateCmd.Parse(os.Args[2:])
		generateConfig()

	default:
		fmt.Printf("不明なコマンド: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func formatJSON(inputFile string) {
	// JSONファイルを読み込んで整形する
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("ファイルを開けません: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// JSONを読み込む
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("ファイルの読み込みに失敗: %v\n", err)
		os.Exit(1)
	}

	// 一時的な interface{} にデコード
	var data interface{}
	if err := json.Unmarshal(content, &data); err != nil {
		fmt.Printf("JSONのパースに失敗: %v\n", err)
		os.Exit(1)
	}

	// 整形してエンコード
	formatted, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("JSONの整形に失敗: %v\n", err)
		os.Exit(1)
	}

	// ファイルに書き込む
	if err := os.WriteFile(inputFile, formatted, 0644); err != nil {
		fmt.Printf("ファイルの書き込みに失敗: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("ファイルを整形しました: %s\n", inputFile)

}

func generateConfig() {
	// kCon := NewConfig()
	kCon := core.KarabinerConfig{}
	file, err := json.MarshalIndent(kCon, "", "  ")
	if err != nil {
		fmt.Printf("設定の生成に失敗: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("generated.json", file, 0644); err != nil {
		fmt.Printf("ファイルの書き込みに失敗: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("設定ファイルを生成しました: generated.json")
}
