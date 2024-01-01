package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Implementing Connect 4 in GoLang")
	var board [8][8]int
	// pass the arguments as references because the functions now take params by reference
	initBoard(&board)

	// addCoin(&board, 3, 1)
	// addCoin(&board, 4, 2)
	// addCoin(&board, 3, 1)
	// addCoin(&board, 4, 2)
	// addCoin(&board, 3, 1)
	// addCoin(&board, 4, 2)
	// addCoin(&board, 3, 1)

	// addCoin(&board, 3, 1)
	// addCoin(&board, 3, 2)
	// addCoin(&board, 2, 1)
	// addCoin(&board, 2, 2)
	// addCoin(&board, 1, 1)
	// addCoin(&board, 1, 2)
	// addCoin(&board, 4, 1)

	// addCoin(&board, 0, 1)
	// addCoin(&board, 1, 2)
	// addCoin(&board, 1, 1)
	// addCoin(&board, 2, 2)
	// addCoin(&board, 2, 1)
	// addCoin(&board, 3, 2)
	// addCoin(&board, 2, 1)
	// addCoin(&board, 3, 2)
	// addCoin(&board, 4, 1)
	// addCoin(&board, 3, 2)
	// addCoin(&board, 3, 1)

	// addCoin(&board, 7, 1)
	// addCoin(&board, 6, 2)
	// addCoin(&board, 6, 1)
	// addCoin(&board, 5, 2)
	// addCoin(&board, 5, 1)
	// addCoin(&board, 4, 2)
	// addCoin(&board, 5, 1)
	// addCoin(&board, 4, 2)
	// addCoin(&board, 3, 1)
	// addCoin(&board, 4, 2)
	// addCoin(&board, 4, 1)

	for {
		displayBoard(&board)
		fmt.Println("Enter the column to drop ")
		var colToDrop int
		fmt.Scanln(&colToDrop)
		addCoin(&board, colToDrop, 1)
		winner := checkSolved(&board)
		if winner != -1 {
			fmt.Println("Player", winner, "wins!")
			break
		}

		// Check for a tie
		if isBoardFull(&board) {
			fmt.Println("It's a tie!")
			break
		}

		turn := 2
		playGame(&board, turn, 5)
		displayBoard(&board)

		winner = checkSolved(&board)
		if winner != -1 {
			fmt.Println("Player", winner, "wins!")
			break
		}

		// Check for a tie
		if isBoardFull(&board) {
			fmt.Println("It's a tie!")
			break
		}
	}

}
func compVsComp(board *[8][8]int, d int) {
	for turn := 1; turn <= 2; turn = 3 - turn {
		playGame(board, turn, d)
		displayBoard(board)

		winner := checkSolved(board)
		if winner != -1 {
			fmt.Println("Player", winner, "wins!")
			break
		}

		// Check for a tie
		if isBoardFull(board) {
			fmt.Println("It's a tie!")
			break
		}
	}
}
func isBoardFull(board *[8][8]int) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 0 {
				return false
			}
		}
	}
	return true

}

func isDroppable(board *[8][8]int, col int) bool {

	if board[0][col] != 0 {
		return false
	} else {
		return true
	}

}

func playGame(board *[8][8]int, turn int, d int) {
	var bestScore int = math.MinInt
	var bestMove int = -1

	for col := 0; col < 8; col++ {
		if isDroppable(board, col) {
			addCoin(board, col, turn)
			score := minimax(board, d, false, turn, math.MinInt, math.MaxInt)
			removeCoin(board, col)

			if score > bestScore {
				bestScore = score
				bestMove = col
			}
		}
	}

	if bestMove != -1 {
		addCoin(board, bestMove, turn)
	}
}

func minimax(board *[8][8]int, depth int, maximizingPlayer bool, turn int, alpha int, beta int) int {
	if depth == 0 || checkSolved(board) != -1 {
		return curScore(board, turn)
	}

	if maximizingPlayer {
		maxEval := math.MinInt
		for col := 0; col < 8; col++ {
			if isDroppable(board, col) {
				addCoin(board, col, turn)
				eval := minimax(board, depth-1, false, turn, alpha, beta)
				removeCoin(board, col)

				maxEval = int(math.Max(float64(maxEval), float64(eval)))
				alpha = int(math.Max(float64(alpha), float64(eval)))

				if beta <= alpha {
					break
				}
			}
		}
		return maxEval
	} else {
		minEval := math.MaxInt
		for col := 0; col < 8; col++ {
			if isDroppable(board, col) {
				addCoin(board, col, 3-turn)
				eval := minimax(board, depth-1, true, turn, alpha, beta)
				removeCoin(board, col)

				minEval = int(math.Min(float64(minEval), float64(eval)))
				beta = int(math.Min(float64(beta), float64(eval)))

				if beta <= alpha {
					break
				}
			}
		}
		return minEval
	}
}

func removeCoin(board *[8][8]int, col int) {
	for i := 0; i < 8; i++ {
		if board[i][col] != 0 {
			board[i][col] = 0
			break
		}
	}
}

func curScore(board *[8][8]int, turn int) int {
	score := 0
	// horizontal one coin away from win layouts

	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j+1] == board[i][j+2] && board[i][j+1] != 0 {
				if board[i][j] == board[i][j+1] || board[i][j+2] == board[i][j+3] {
					if board[i][j+1] == turn {
						score += 10
					} else {
						score -= 20
					}
				}
			}
		}
	}

	b_trans := transpose(board)

	for i := 0; i < 8; i++ {
		for j := 0; j < 5; j++ {
			if b_trans[i][j+1] == b_trans[i][j+2] && b_trans[i][j+1] != 0 {
				if b_trans[i][j] == b_trans[i][j+1] || b_trans[i][j+2] == b_trans[i][j+3] {
					if b_trans[i][j+1] == turn {
						score += 10
					} else {
						score -= 20
					}
				}
			}
		}
	}

	return score

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

	// displayBoard(board)
	fmt.Println("The current board score is ", curScore(board, turn))
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

	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		if board[i][j] == board[i+1][j+1] && board[i+1][j+1] == board[i+2][j+2] && board[i+2][j+2] == board[i+3][j+3] && board[i+2][j+2] != 0 {
	// 			return board[i+2][j+2]
	// 		}
	// 	}
	// }

	// b_trans = transpose(board)
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 5; j++ {
	// 		if b_trans[i][j] == b_trans[i+1][j+1] && b_trans[i+1][j+1] == b_trans[i+2][j+2] && b_trans[i+2][j+2] == b_trans[i+3][j+3] && b_trans[i+2][j+2] != 0 {
	// 			return b_trans[i+2][j+2]
	// 		}
	// 	}
	// }

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == board[i+1][j+1] && board[i+1][j+1] == board[i+2][j+2] && board[i+2][j+2] == board[i+3][j+3] && board[i+2][j+2] != 0 {
				return board[i+2][j+2]
			}
		}
	}

	// checking diagonally (from top-right to bottom-left)
	b_trans = transpose(board)
	for i := 0; i < 5; i++ {
		for j := 3; j < 8; j++ {
			if b_trans[i][j] == b_trans[i+1][j-1] && b_trans[i+1][j-1] == b_trans[i+2][j-2] && b_trans[i+2][j-2] == b_trans[i+3][j-3] && b_trans[i+2][j-2] != 0 {
				return b_trans[i+2][j-2]
			}
		}
	}

	return -1

}
