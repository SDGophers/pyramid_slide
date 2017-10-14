package main

import (
	"fmt"
	"math"
)

var rows1 = [][]int{
	{3},
	{7, 4},
	{2, 4, 6},
	{8, 5, 9, 8},
}

var med_arr = [][]int{
	{75},
	{95, 64},
	{17, 47, 82},
	{18, 35, 87, 10},
	{20, 4, 82, 47, 65},
	{19, 1, 23, 75, 3, 34},
	{88, 2, 77, 73, 7, 63, 67},
	{99, 65, 4, 28, 6, 16, 70, 92},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
}

func traverseDirctions(takeLeft int, p [][]int) int {
	sum := p[0][0]

	column := 0

	for _, v := range p[1:] {
		if takeLeft%2 == 1 {
			sum += v[column]
		} else {
			column++
			sum += v[column]
		}
		takeLeft >>= 1
	}

	return sum
}

func traverseDirctionsAndPrint(takeLeft int, p [][]int) int {
	sum := p[0][0]

	column := 0

	for _, v := range p[1:] {
		if takeLeft%2 == 1 {
			sum += v[column]
			fmt.Println(v[column])
		} else {
			column++
			sum += v[column]
			fmt.Println(v[column])
		}
		//this can only handle bitSize number of rows
		takeLeft >>= 1
	}

	return sum
}

func bruteForce(p [][]int) int {
	numRows := len(p)

	possibilities := int(math.Pow(2, float64(numRows)))

	aSum := int(1e10)
	for i := 0; i < possibilities; i++ {
		var s int
		if s = traverseDirctions(i, p); s < aSum {
			aSum = s
		}
	}

	return aSum
}

func traverse(p [][]int) int {
	sum := p[0][0]

	column := 0

	for _, row := range p[1:] {
		if row[column] < row[column+1] {
			sum += row[column]
		} else {
			column++
			sum += row[column]
		}
	}

	return sum
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func reverseTraverse(p [][]int) int {
	for i := len(p) - 2; i >= 0; i-- {
		for j := range p[i] {
			p[i][j] += min(p[i+1][j], p[i+1][j+1])
		}
	}

	return p[0][0]
}
