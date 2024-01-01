package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 拡張版と通常版を保持
	// 拡張版のみに変数を追加、ただわかりやすいように通常版も拡張と同じにする
	// sliceを拡張するためのコピー元を持つときに全てfalseとしておく。
}



func board(fName string) [][]int {
	var sudoku [][]int
	file, err := os.Open(fName)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")

		var row []int
		for _, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Printf("Error converting string to int: %s\n", err)
				os.Exit(1)
			}
			row = append(row, num)
		}
		sudoku = append(sudoku, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	for _, row := range sudoku {
		fmt.Println(row)
	}
	return sudoku
}

