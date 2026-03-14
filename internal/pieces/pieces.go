package pieces

import "ChessEngineEOF/internal/model"

const (
	Empty  = 0
	Pawn   = 100
	Bishop = 300
	Knight = 310
	Rook   = 500
	Queen  = 900
	King   = 2000000

	PawnE    = -100
	BishopE  = -300
	KnightE  = -310
	RookE    = -500
	QueenE   = -900
	KingE    = -2000000
	Mate     = 1000000000000
	Infinity = Mate + 1
)

func MyPiece(b model.BoardP, piece int) bool {
	if piece == 0 {
		return false
	}
	if b.White {
		return piece > 0
	} else {
		return piece < 0
	}

}
