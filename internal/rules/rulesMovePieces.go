package rules

import "ChessEngineEOF/internal/model"

func LogicOfMovePieces(b model.Board, p model.PiecesMove) {

}
func CanMove(CanMove bool, MoveWhite bool, b model.Board, p model.PiecesMove) (*model.Board, bool) {
	if CanMove {

		if MoveWhite {
			b.Board[p.MHX][p.MHY] = b.Board[p.MFX][p.MFY]
			b.Board[p.MFX][p.MFY] = 0
			b.White = false
			return &b, true
		} else {
			b.Board[p.MHX][p.MHY] = b.Board[p.MFX][p.MFY]
			b.Board[p.MFX][p.MFY] = 0
			b.White = true
			return &b, true
		}
	} else {
		return &b, false
	}
}
