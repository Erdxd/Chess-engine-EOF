package main

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/board"
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/position"
	"ChessEngineEOF/internal/rules"
	"ChessEngineEOF/search"
	"fmt"
)

func main() {
	deepth := 5
	fmt.Println("Welcome to Chess Engine EOF")
	boardNew := board.NewBoard()
	BoardIinit := board.InitPieces(boardNew)
	BoardFork1 := position.Fork1(BoardIinit)
	PlayGame(BoardFork1, deepth)

}
func PlayGame(b *model.BoardP, deepth int) {
	b.White = true
	Count := 0

	for Count < 30 {
		if b.White {
			fmt.Println("Move white")
			fmt.Println("Move:", Count+1)
		} else {
			fmt.Println("Move black")
			fmt.Println("Move:", Count+1)

		}
		Moves := movegen.GenerateMoves(b)
		BestMove := search.TheBestMove(b, deepth)
		if len(Moves) == 0 {
			break
		}
		b, _, _, _ = rules.ApplyMove(true, b.White, b, &BestMove)
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
