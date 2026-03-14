package movegen

import "ChessEngineEOF/internal/model"

func QueenMoveGen(b *model.BoardP, x, y int) []model.PiecesMove {
	Moves := []model.PiecesMove{}
	Moves = append(Moves, BishopMoveGen(b, x, y)...)
	Moves = append(Moves, RookMoveGen(b, x, y)...)
	return Moves
}
