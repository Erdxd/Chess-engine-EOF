package rules

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func LogicOfMovePieces(b *model.Board, p *model.PiecesMove) (bool, bool) {
	var CanMove bool
	var WhiteMove bool
	switch p.Pieces {
	case pieces.Pawn, pieces.PawnE:
		CanMove, WhiteMove = PawnMove(*b, *p)
		return CanMove, WhiteMove
	case pieces.Knight, pieces.KnightE:
		CanMove, WhiteMove = KnightMove(*b, *p)
		return CanMove, WhiteMove
	case pieces.Bishop, pieces.BishopE:
		CanMove, WhiteMove = BishopMove(*b, *p)
		return CanMove, WhiteMove
	case pieces.Rook, pieces.RookE:
		CanMove, WhiteMove = RookMove(*b, *p)
		return CanMove, WhiteMove
	case pieces.Queen, pieces.QueenE:
		CanMove, WhiteMove = QueenMove(*b, *p)
		return CanMove, WhiteMove
	case pieces.King, pieces.KingE:
		CanMove, WhiteMove = KingMove(*b, *p)
		return CanMove, WhiteMove
	}

	return CanMove, WhiteMove
}
func ApplyMove(CanMove bool, MoveWhite bool, b *model.Board, p *model.PiecesMove) (*model.Board, bool) {
	if CanMove {

		if MoveWhite {
			b.Board[p.MHX][p.MHY] = b.Board[p.MFX][p.MFY]
			b.Board[p.MFX][p.MFY] = 0
			b.White = false
			return b, true
		} else {
			b.Board[p.MHX][p.MHY] = b.Board[p.MFX][p.MFY]
			b.Board[p.MFX][p.MFY] = 0
			b.White = true
			return b, true
		}
	} else {
		return b, false
	}
}
