package movegen

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
	"ChessEngineEOF/internal/rules"
)

func GenerateMoves(b *model.Board) []model.PiecesMove {

	moves := []model.PiecesMove{}
	for x := 1; x <= 8; x++ {
		for y := 1; y <= 8; y++ {
			piece := b.Board[x][y]
			if piece == 0 || !pieces.MyPiece(*b, piece) {
				continue
			}
			switch piece {
			case pieces.Pawn, pieces.PawnE:
				MovesPawn := PawnMoveGenerate(b, x, y)

				moves = append(moves, MovesPawn...)

			case pieces.Knight, pieces.KnightE:
				CanMove, WhiteMove = rules.KnightMove(b, &model.PiecesMove{})
				return
			case pieces.Bishop, pieces.BishopE:
				movesBishop := BishopMoveGen(b, x, y)
				moves = append(moves, movesBishop...)
			case pieces.Rook, pieces.RookE:
				MovesRook := RookMoveGen(b, x, y)
				moves = append(moves, MovesRook...)
			case pieces.Queen, pieces.QueenE:
				MovesQueen := QueenMoveGen(b, x, y)
				moves = append(moves, MovesQueen...)
			case pieces.King, pieces.KingE:
				CanMove, WhiteMove = rules.KingMove(b, &model.PiecesMove{})
			}

		}
	}
}
