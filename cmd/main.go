package main

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/board"
	"ChessEngineEOF/internal/evaluate"
	"ChessEngineEOF/search"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Chess Engine EOF")
	boardNew := board.NewBoard()
	BoardIinit := board.InitPieces(boardNew)
	fmt.Println(BoardIinit)
	moves := movegen.GenerateMoves(BoardIinit)
	fmt.Println(moves)
	s := evaluate.EvaluateBoard(BoardIinit)
	fmt.Println(s)
	deepth := 1
	score := search.MiniMax(deepth, BoardIinit, true)
	fmt.Println(deepth, score)
}
