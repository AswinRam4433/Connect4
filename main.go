package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Implementing Connect 4 in GoLang")
	var board [8][8]int
	initBoard(&board)
	addCoin(&board, 3, 1)
	addCoin(&board, 4, 2)
	addCoin(&board, 3, 1)
	addCoin(&board, 4, 2)
	addCoin(&board, 3, 1)
	addCoin(&board, 4, 2)
	addCoin(&board, 3, 1)
	// pass the arguments as references because the functions now take params by reference

}

func displayBoard(board *[8][8]int) {
	fmt.Println("\nPrinting the board ")
	for i := 0; i < 8; i++ {
		fmt.Println(board[i])

	}
}
func addCoin(board *[8][8]int, col int, turn int) {
	// pass the args by reference to cause the states to persist
	// col specifies the column into which we have to drop the coin. Ranges from 0 to 7
	// turn specifies which player is going to play. Alternates between 1 and 2

	// fmt.Println("In the add coin func")

	if board[0][col] != 0 {
		fmt.Println("Invalid Move")
	}
	for i := 7; i >= 0; i-- {
		// fmt.Println("The value of board i col is ", board[i][col])
		if board[i][col] == 0 {
			board[i][col] = turn
			fmt.Println("Added the coin")
			break
		}

	}
	checkValidBoard(board)

	displayBoard(board)
	ans := checkSolved(board)
	if ans != -1 {
		fmt.Println("We have a winner  ", ans)
	}
}
func checkValidBoard(board *[8][8]int) {
	fmt.Println("Checking if board is valid ")
	var a int = 0
	var b int = 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 1 {
				a++
			} else if board[i][j] == 2 {
				b++
			}

		}
	}
	if math.Abs(float64(a)-float64(b)) > 1 {
		panic("The board is inconsistent")
	}

}
func initBoard(board *[8][8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board[i][j] = 0
		}
	}
	displayBoard(board)
}

func transpose(a *[8][8]int) [][]int {
	newArr := make([][]int, len(*a))
	for i := 0; i < len(*a); i++ {
		for j := 0; j < len((*a)[0]); j++ {
			newArr[j] = append(newArr[j], (*a)[i][j])
		}
	}
	return newArr
}

func checkSolved(board *[8][8]int) int {
	fmt.Println("Checking for winner: ")

	// checking row wise
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == board[i][j+1] && board[i][j+1] == board[i][j+2] && board[i][j+2] == board[i][j+3] && board[i][j+1] != 0 {
				return board[i][j]
			}
		}
	}

	// checking column wise
	b_trans := transpose(board)
	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			if b_trans[i][j] == b_trans[i][j+1] && b_trans[i][j+1] == b_trans[i][j+2] && b_trans[i][j+2] == b_trans[i][j+3] && b_trans[i][j+2] != 0 {
				return b_trans[i][j]
			}
		}
	}

	return -1

}
