package main

import (
	"ChessEngineEOF/internal/model"
	"fmt"
)

func main() {
	fmt.Println("Welcome to Chess Engine EOF")
	fmt.Println(model.NewBoard().Board)
}
