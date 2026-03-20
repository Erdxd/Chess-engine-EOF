package rules

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func IsMate(b *model.BoardP) bool {

	if IsCheck(b) != true {
		return false
	}

	AllMoves := movegen.GenerateMoves(b)

	for _, move := range AllMoves {
		testBoard := &model.BoardP{}
		*testBoard = *b
		testBoard, Can, _, _ := ApplyMove(true, b.White, testBoard, &move)
		if !Can {
			continue
		}
		if !IsCheck(testBoard) {

			return false

		}

	}

	return true
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
	bcheck := b.White
	b.White = !b.White

	moves := movegen.GenerateMoves(b)
	for _, move := range moves {
		if move.MHX == kingX && move.MHY == kingY {
			return true
		}
	}
	b.White = bcheck
	return false
}
