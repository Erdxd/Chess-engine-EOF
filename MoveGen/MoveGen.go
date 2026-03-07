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
			switch p.Piece {
			case pieces.Pawn, pieces.PawnE:
				CanMove, WhiteMove = rules.PawnMoveGenerate(b, &model.PiecesMove{})
				return
			case pieces.Knight, pieces.KnightE:
				CanMove, WhiteMove = rules.KnightMove(b, &model.PiecesMove{})
				return
			case pieces.Bishop, pieces.BishopE:
				CanMove, WhiteMove = rules.BishopMove(b, &model.PiecesMove{})
				return
			case pieces.Rook, pieces.RookE:
				CanMove, WhiteMove = rules.RookMove(b, &model.PiecesMove{})
				return
			case pieces.Queen, pieces.QueenE:
				CanMove, WhiteMove = rules.QueenMove(b, &model.PiecesMove{})
				return
			case pieces.King, pieces.KingE:
				CanMove, WhiteMove = rules.KingMove(b, &model.PiecesMove{})
				return
			}

		}
	}
}
