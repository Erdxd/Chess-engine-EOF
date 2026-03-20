package rules

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func IsMate(b *model.BoardP) bool {
	var mate bool

	return mate
}
func IsCheck(b *model.BoardP) bool {
	var kingX, kingY int

	for x := 0; x <= 8; x++ {
		for y := 0; y <= 8; y++ {
			if b.Board[x][y] == pieces.King && b.White {
				kingX, kingY = x, y
			} else if b.Board[x][y] == pieces.KingE && !b.White {
				kingX, kingY = x, y
			}

		}
	}
	b.White = !b.White
	moves := movegen.GenerateMoves(b)
	for _, move := range moves {
		if move.MHX == kingX && move.MHY == kingY {
			return true
		}
	}
	return false
}
