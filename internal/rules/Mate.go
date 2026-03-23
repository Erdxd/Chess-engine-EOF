package rules

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func IsMate(b *model.BoardP) bool {

	kingfound, check := IsCheck(b)
	if !kingfound {
		return true
	}
	if !check {
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
		kingfound, check := IsCheck(testBoard)
		if !kingfound {
			return true
		}
		if !check {
			return false
		}

	}

	return true
}
func IsCheck(b *model.BoardP) (bool, bool) {
	var kingX, kingY int
	kingfound := false
	for x := 0; x <= 8; x++ {
		for y := 0; y <= 8; y++ {
			if b.Board[x][y] == pieces.King && b.White {
				kingX, kingY = x, y
				kingfound = true
			} else if b.Board[x][y] == pieces.KingE && !b.White {
				kingX, kingY = x, y
				kingfound = true
			}

		}
	}
	bcheck := b.White
	b.White = !b.White

	moves := movegen.GenerateMoves(b)
	for _, move := range moves {
		if move.MHX == kingX && move.MHY == kingY {
			return kingfound, true
		}
	}
	b.White = bcheck
	return kingfound, false
}
