package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

const (
	inputFileName  = "input.csv"
	outputFileName = "output.csv"
)

func main() {
	// step1: csvファイルを読み込み
	inputFile, err := os.OpenFile(inputFileName, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}
	defer inputFile.Close()

	// step2: ファイルデータから構造体を生成
	var users []User
	if err := gocsv.Unmarshal(inputFile, &users); err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}

	// step3: 出力ファイルを生成
	outputFile, err := os.OpenFile(outputFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}
	defer outputFile.Close()

	// step4: 出力ファイルに書き込み
	if err := gocsv.MarshalFile(users, outputFile); err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}

	return
}

type User struct {
	Id    string `csv:"id"`
	Name  string `csv:"name"`
	Age   int    `csv:"age"`
	Hobby string `csv:"hobby"`
}
