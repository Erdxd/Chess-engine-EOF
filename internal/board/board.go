package board

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func NewBoard() *model.BoardP {
	b := &model.BoardP{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.Board[i][j] = 99
		}

	}
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			b.Board[i][j] = 0
		}
	}
	return b

}
func InitPieces(b *model.BoardP) *model.BoardP {
	RowHeavypieces := []int{pieces.Rook, pieces.Knight, pieces.Bishop, pieces.Queen, pieces.King, pieces.Bishop, pieces.Knight, pieces.Rook}
	RowHeavypiecesE := []int{pieces.RookE, pieces.KnightE, pieces.BishopE, pieces.QueenE, pieces.KingE, pieces.BishopE, pieces.KnightE, pieces.RookE}

	for i := 1; i <= 8; i++ {
		b.Board[i][2] = pieces.Pawn
		b.Board[i][7] = pieces.PawnE
	}
	for i, v := range RowHeavypieces {
		b.Board[i+1][1] = v
	}
	for i, v := range RowHeavypiecesE {
		b.Board[i+1][8] = v
	}

	return b
}
