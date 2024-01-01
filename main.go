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

var seq = intSeq()

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sudoku <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]

	dboard := board(filename)

	/*
		dboard = [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
	*/

	// expand
	expandedDBoard := make([][]int, len(dboard)+2)
	expandedDBoardVars := make([][]int, len(dboard)+2)
	for i := range expandedDBoard {
		expandedDBoard[i] = make([]int, len(dboard[0])+2)
		expandedDBoardVars[i] = make([]int, len(dboard[0])+2)
	}

	for i, row := range dboard {
		for j, value := range row {
			expandedDBoard[i+1][j+1] = value
			expandedDBoardVars[i+1][j+1] = seq()

		}
	}

	// ==== for later check//
	file, err := os.Create("tvars")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i, row := range expandedDBoardVars {
		line := ""
		for j, v := range row {
			if i == 0 || j == 0 || i == len(expandedDBoard)-1 || j == len(row)-1 {
				continue
			}
			line += strconv.Itoa(v) + " "
		}
		_, err := file.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	// ==== //

	// add -1 to new cells for expansion
	for i, v := range expandedDBoard {
		for j := range v {
			if i == 0 || j == 0 || i == len(expandedDBoard)-1 || j == len(v)-1 {
				expandedDBoard[i][j] = -1
				t := seq()
				expandedDBoardVars[i][j] = t
				clauses = append(clauses, Clause{-t})
			}
		}
	}

	for i, v := range expandedDBoard {
		for j, vv := range v {
			if vv == -1 {
				continue
			}

			// vars around target cell
			tVars := []int{}
			for ii := -1; ii <= 1; ii++ {
				for jj := -1; jj <= 1; jj++ {
					tVars = append(tVars, expandedDBoardVars[i+ii][j+jj])
				}
			}
			// Determine
			if vv == 0 {
				for _, v := range tVars {
					clauses = append(clauses, Clause{-v})
				}
				continue
			}
			// Determine
			if vv == 9 {
				for _, v := range tVars {
					clauses = append(clauses, Clause{v})
				}
				continue
			}

			if vv == 8 {
				c := []int{}
				for _, v := range tVars {
					c = append(c, -v)
				}
				clauses = append(clauses, c)
			}

			// true isNot Positive
			for k := vv; k < 8; k++ {
				comb(k+1, 9-(k+1), false, tVars)
			}

			if vv == 1 {
				c := []int{}
				for _, v := range tVars {
					c = append(c, v)
				}
				clauses = append(clauses, c)
			}

			// false is Postive
			for k := 9 - vv; k < 8; k++ {
				comb(k+1, 9-(k+1), true, tVars)
			}
		}
	}

	cnf := clausesToString(clauses, seq()-1)

	// Save the CNF to a text file
	fOut := "r_cnf.txt"
	if err := os.WriteFile(fOut, []byte(cnf), 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("CNF file generated successfully:", fOut)
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

	// for _, row := range sudoku {
	// 	fmt.Println(row)
	// }
	return sudoku
}

func first(n int) uint {
	return (1 << n) - 1
}

// n1+n2 C n1
func comb(n1, n2 int, isPostive bool, tVars []int /*9*/) {
	var (
		j, k                                int
		x, s                                uint
		smallest, ripple, newSmallest, ones uint
	)

	m := make([]int, n1+1)

	x = first(n1)
	for (x & ^first(n1+n2)) == 0 {
		s = x
		k = 1
		for j = 1; j <= n1+n2; j++ {
			if s&1 != 0 {
				m[k] = j
				k++
			}
			s >>= 1
		}
		c := Clause{}
		for k = 1; k <= n1; k++ {
			if isPostive {
				// m[k] = 1 ~ max + 1
				c = append(c, tVars[m[k]-1])
			} else {
				c = append(c, -tVars[m[k]-1])
			}
			// fmt.Printf(" %2d", m[k])
		}
		clauses = append(clauses, c)

		smallest = x & -x
		ripple = x + smallest
		newSmallest = ripple & -ripple
		ones = ((newSmallest / smallest) >> 1) - 1
		x = ripple | ones
	}
}

func clausesToString(clauses []Clause, varCount int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("p cnf %d %d\n", varCount, len(clauses)))
	for _, clause := range clauses {
		for _, lit := range clause {
			sb.WriteString(fmt.Sprintf("%d ", lit))
		}
		sb.WriteString("0\n")
	}
	return sb.String()
}
