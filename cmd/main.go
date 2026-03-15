package main

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/board"
	"ChessEngineEOF/internal/evaluate"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Chess Engine EOF")
	boardNew := board.NewBoard()
	BoardIinit := board.InitPieces(boardNew)
	fmt.Println(BoardIinit)
	moves := movegen.GenerateMoves(BoardIinit)
	fmt.Println(moves)
	White, Black := evaluate.EvaluateBoard(boardNew)
	fmt.Println(White)
	fmt.Println(Black)

}
