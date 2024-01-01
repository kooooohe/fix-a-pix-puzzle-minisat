package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Clause []int

var clauses []Clause

func intSeq() func() int {
	c := 0
	return func() int {
		c++
		return c
	}
}

var seq intSeq

func main() {
	// 拡張版と通常版を保持
	// 拡張版のみに変数を追加、ただわかりやすいように通常版も拡張と同じにする
	// sliceを拡張するためのコピー元を持つときに全てfalseとしておく。
		if len(os.Args) < 2 {
		fmt.Println("Usage: sudoku <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	// baseClauses := generateClauses()

	dboard := board(filename)

	// expand 
    expandedDBoard := make([][]int, len(dboard)+2)
    expandedDBoardVars := make([][]int, len(dboard)+2)
    for i := range expandedDBoard {
        expandedDBoard[i] = make([]int, len(dboard[0])+2)
        expandedDBoardVars[i] = make([]int, len(dboard[0])+2)
	//TODOここに全てfalseが入るようにする+ -1 を入れる
    }

    // 元のスライスの要素を新しいスライスにコピー
    for i, row := range dboard {
        for j, value := range row {
            expandedDBoard[i+1][j+1] = value
	    expandedDBoardVars[i+1][j+i] = seq()
        }
    }

    // 端をマイナス1にする、varsのはじをマイナスにする
    for i,v := range expandedDBoard {
	    for j,vv := range v {
	    }
    }
	fmt.Println(len(dboard[0]))
	// fmt.Println(len(dboard[2]))
	// fmt.Println(len(dboard[3]))
	// fmt.Println(len(dboard[16]))

	/*
 // 3x3のスライスを定義
    original := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }

    // 5x5の新しいスライスを初期化
    expanded := make([][]int, 5)
    for i := range expanded {
        expanded[i] = make([]int, 5)
	//TODOここに全てfalseが入るようにする+ -1 を入れる
    }

    // 元のスライスの要素を新しいスライスにコピー
    for i, row := range original {
        for j, value := range row {
            expanded[i+1][j+1] = value
        }
    }

    // 結果を表示
    fmt.Println("Original Slice:")
    for _, row := range original {
        fmt.Println(row)
    }

    fmt.Println("\nExpanded Slice:")
    for _, row := range expanded {
        fmt.Println(row)
    }
	*/

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

