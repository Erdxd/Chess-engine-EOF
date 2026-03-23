package main

import (
	"ChessEngineEOF/internal/board"
	"ChessEngineEOF/internal/evaluate"
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
	"ChessEngineEOF/internal/rules"
	"ChessEngineEOF/search"
	"fmt"
)

func main() {
	deepth := 2
	fmt.Println("Welcome to Chess Engine EOF")
	boardNew := board.NewBoard()
	BoardIinit := board.InitPieces(boardNew)
	//BoardMateInOne1 := position.MateInOne1(BoardIinit)
	PlayGame(BoardIinit, deepth)

}
func PlayGame(b *model.BoardP, deepth int) {
	var winner string
	PrintBoard(b)

	Count := 0

	for Count < 5 {
		if b.White {
			fmt.Println("Move white")
			fmt.Println("Move:", Count+1)
		} else {
			fmt.Println("Move black")
			fmt.Println("Move:", Count+1)

		}

		BestMove, score := search.TheBestMove(b, deepth)
		if score == pieces.Mate {
			fmt.Println("White is winner")
		} else if score == -pieces.Mate {
			fmt.Println("Black is winner")
		}
		b, _, _, _ = rules.ApplyMove(true, b.White, b, &BestMove)
		if evaluate.EvaluateBoard(b) == pieces.Mate || evaluate.EvaluateBoard(b) == -pieces.Mate {
			if b.White {
				winner = "White"
			} else {
				winner = "Black"
			}
			fmt.Println("End of the game:", winner)
		}
		PrintBoard(b)
		Count++
	}
}
func PrintBoard(b *model.BoardP) {
	fmt.Println("  a  b  c  d  e  f  g  h")
	for y := 8; y >= 1; y-- {
		fmt.Printf("%d ", y)
		for x := 1; x <= 8; x++ {
			val := b.Board[x][y]
			switch val {
			case 0:
				fmt.Print(".  ")
			case 99:
				fmt.Print("#  ")
			case 100:
				fmt.Print("P  ") // White pawn
			case -100:
				fmt.Print("p  ") // Black pawn
			case 310:
				fmt.Print("N  ") // White knight
			case -310:
				fmt.Print("n  ") // Black knight
			case 300:
				fmt.Print("B  ") // White bishop
			case -300:
				fmt.Print("b  ") // Black bishop
			case 500:
				fmt.Print("R  ") // White rook
			case -500:
				fmt.Print("r  ") // Black rook
			case 900:
				fmt.Print("Q  ") // White queen
			case -900:
				fmt.Print("q  ") // Black queen
			case 2000000:
				fmt.Print("K  ") // White king
			case -2000000:
				fmt.Print("k  ") // Black king
			default:
				fmt.Printf("?  ") // Что-то странное
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
