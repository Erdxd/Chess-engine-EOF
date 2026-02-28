package board

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func NewBoard() *model.Board {
	b := &model.Board{}
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
func InitPieces(b *model.Board) *model.Board {
	RowHeavypieces := []int{pieces.Rook, pieces.Knight, pieces.Bishop, pieces.Queen, pieces.King, pieces.Bishop, pieces.Knight, pieces.Rook}
	RowHeavypiecesE := []int{pieces.RookE, pieces.KnightE, pieces.BishopE, pieces.QueenE, pieces.KingE, pieces.BishopE, pieces.KnightE, pieces.RookE}

	for i := 1; i <= 8; i++ {
		b.Board[2][i] = pieces.Pawn
		b.Board[7][i] = pieces.PawnE
	}
	for i, v := range RowHeavypieces {
		b.Board[1][i+1] = v
	}
	for i, v := range RowHeavypiecesE {
		b.Board[8][i+1] = v
	}

	return b
}
