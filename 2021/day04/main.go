package main

import "fmt"

func main() {
	initBoards()
	lastwinBoard := bingo{}
	lastwinCall := 0
	for _, call := range calls {
		for i, board := range inputs {
			if board.won {
				continue
			}
			for x, line := range board.board {
				for y := range line {
					if call == board.board[x][y] {
						inputs[i].unmarked -= call
						if board.markSum[0][x] == 4 || board.markSum[1][y] == 4 {
							lastwinBoard = inputs[i]
							lastwinCall = call
							inputs[i].won = true
						}
						inputs[i].markSum[0][x]++
						inputs[i].markSum[1][y]++
					}
				}
			}
		}
	}
	callWinner(lastwinBoard, lastwinCall)
}

func initBoards() {
	for i, board := range inputs {
		sum := 0
		for x, line := range board.board {
			for y := range line {
				sum += board.board[x][y]
			}
		}
		inputs[i].unmarked = sum
		inputs[i].markSum = [2][5]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}
	}
	// fmt.Printf("%v", inputs[0])
}

func callWinner(b bingo, c int) {
	fmt.Println("win")
	fmt.Printf("%v\n", b)
	fmt.Printf("%d\n", b.unmarked*c)
}
