package main

import "fmt"

//globals
var arrows [30][30]string
var numbers [30][30]int

func main() {

	fmt.Println("Welcome to the longest common subsequence program!")
	fmt.Printf("Input your first string: ")
	var firstString string

	//get the user input
	fmt.Scan(&firstString)
	fmt.Println("")

	fmt.Printf("Input your second string: ")
	var secondString string

	//get the user input
	fmt.Scan(&secondString)
	fmt.Println("")

	//call lcs length to construct tables
	lcsLength(firstString, secondString)

	//print the lcs
	fmt.Printf("LCS: ")

	printLCS(firstString, len(firstString), len(secondString))

}

//method to find the longest common subsequence between 2 strings
func lcsLength(first, second string) {

	//get the length of the strings
	m := len(first)
	n := len(second)

	//make first column and row filled with 0's
	for i := 1; i <= m; i++ {
		numbers[i][0] = 0
	}
	for i := 0; i <= n; i++ {
		numbers[0][i] = 0
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {

			if first[i-1] == second[j-1] {

				//change the number
				numbers[i][j] = (numbers[i-1][j-1] + 1)

				//assign the arrow
				arrows[i][j] = "diagonal (up-left)"

			} else if numbers[i-1][j] >= numbers[i][j-1] {

				//set the number
				numbers[i][j] = numbers[i-1][j]

				//assign the arrow
				arrows[i][j] = "up"

			} else {

				//set the number
				numbers[i][j] = numbers[i][j-1]

				//assign the arrow
				arrows[i][j] = "left"

			}

		}
	}
}

func printLCS(x string, i int, j int) {

	if i == 0 || j == 0 {

		return
	}

	if arrows[i][j] == "diagonal (up-left)" {
		printLCS(x, i-1, j-1)

		fmt.Printf(string(x[i-1]))

	} else if arrows[i][j] == "up" {

		printLCS(x, i-1, j)

	} else {
		printLCS(x, i, j-1)
	}

}
