package rules

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func LogicOfMovePieces(b *model.BoardP, p *model.PiecesMove) (bool, bool) {
	var CanMove bool
	var WhiteMove bool
	switch p.Piece {
	case pieces.Pawn, pieces.PawnE:
		CanMove, WhiteMove = PawnMove(b, p)
		return CanMove, WhiteMove
	case pieces.Knight, pieces.KnightE:
		CanMove, WhiteMove = KnightMove(b, p)
		return CanMove, WhiteMove
	case pieces.Bishop, pieces.BishopE:
		CanMove, WhiteMove = BishopMove(b, p)
		return CanMove, WhiteMove
	case pieces.Rook, pieces.RookE:
		CanMove, WhiteMove = RookMove(b, p)
		return CanMove, WhiteMove
	case pieces.Queen, pieces.QueenE:
		CanMove, WhiteMove = QueenMove(b, p)
		return CanMove, WhiteMove
	case pieces.King, pieces.KingE:
		CanMove, WhiteMove = KingMove(b, p)
		return CanMove, WhiteMove
	}

	return CanMove, WhiteMove
}
func ApplyMove(CanMove bool, MoveWhite bool, b *model.BoardP, p *model.PiecesMove) (*model.BoardP, bool, int, model.PiecesMove) {
	if CanMove {
		was := b.Board[p.MHX][p.MHY]

		if MoveWhite {

			b.Board[p.MHX][p.MHY] = b.Board[p.MFX][p.MFY]
			b.Board[p.MFX][p.MFY] = 0
			b.White = false
			return b, true, was, model.PiecesMove{
				MFX:   p.MFX,
				MFY:   p.MFY,
				MHX:   p.MHX,
				MHY:   p.MHY,
				Piece: b.Board[p.MHX][p.MHY],
			}
		} else {
			b.Board[p.MHX][p.MHY] = b.Board[p.MFX][p.MFY]
			b.Board[p.MFX][p.MFY] = 0
			b.White = true
			return b, true, was, model.PiecesMove{
				MFX:   p.MFX,
				MFY:   p.MFY,
				MHX:   p.MHX,
				MHY:   p.MHY,
				Piece: b.Board[p.MHX][p.MHY],
			}
		}
	} else {
		return b, false, 0, model.PiecesMove{}
	}
}
func UndoMove(b *model.BoardP, piece int, p *model.PiecesMove) model.BoardP {

	b.Board[p.MFX][p.MFY] = p.Piece
	b.Board[p.MHX][p.MHY] = piece
	b.White = !b.White
	return *b
}
