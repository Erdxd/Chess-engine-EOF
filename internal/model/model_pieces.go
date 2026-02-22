package model

import "math"

const (
	Empty  = 0
	Pawn   = 100
	Bishop = 300
	Knight = 310
	Rook   = 500
	Queen  = 900
	King   = 2000000

	PawnE   = -100
	BishopE = -300
	KnightE = -310
	RookE   = -500
	QueenE  = -900
	KingE   = -2000000
)

type Board struct {
	Board  [10][10]int
	White  bool
	Move2W bool
	Move2B bool
}

func NewBoard() *Board {
	b := &Board{}
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
func InitPieces(b *Board) *Board {
	RowHeavypieces := []int{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	RowHeavypiecesE := []int{RookE, KnightE, BishopE, QueenE, KingE, BishopE, KnightE, RookE}

	for i := 1; i <= 8; i++ {
		b.Board[2][i] = Pawn
		b.Board[7][i] = PawnE
	}
	for i, v := range RowHeavypieces {
		b.Board[1][i+1] = v
	}
	for i, v := range RowHeavypiecesE {
		b.Board[8][i+1] = v
	}

	return b
}
func MoveMaker(b *Board, mFX, mFY, mHX, mHY int, pieces int) (*Board, bool) {
	var MoveWhite bool
	var CanMove bool

	switch {
	case pieces == Pawn || pieces == PawnE:

		if b.White == true {
			if b.Board[mHX][mHY] > 0 || b.Board[mHX][mHY] == 99 {
				CanMove = false
			} else {
				if mHY-mFY < 2 && mHY-mFY > 0 && Pawn != b.Board[mFX][2] && mFX == mHX && b.Board[mFX][mHY] == 0 {
					MoveWhite = true
					CanMove = true

				} else if Pawn == b.Board[mFX][2] && mHY-mFY > 1 && mHY-mFY < 3 && b.Board[mFX][mHY] == 0 && mFX == mHX && b.Board[mFX][mHY-1] == 0 {
					MoveWhite = true
					CanMove = true
					b.Move2W = true
				} else if b.Board[mHX][mHY] < 0 && mHY-mFY == 1 && math.Abs(float64(mHX)-float64(mFX)) == 1 {
					MoveWhite = true
					CanMove = true
				} else if b.Move2B == true && b.Board[mFX][5] == Pawn && mHY-mFY < 2 && math.Abs(float64(mHX)-float64(mFX)) == 1 && mHY-mFY > 0 {
					MoveWhite = true
					CanMove = true
					b.Board[mHX][mHY-1] = Empty
				} else {

					CanMove = false
				}

			}
			if CanMove {
				b.Move2B = false
			}

		} else {
			if b.Board[mHX][mHY] < 0 || b.Board[mHX][mHY] == 99 {
				CanMove = false
			} else {
				if mFY-mHY < 2 && mFY-mHY > 0 && PawnE != b.Board[mFX][7] && mFX == mHX && b.Board[mFX][mHY] == 0 {
					MoveWhite = false
					CanMove = true
				} else if PawnE == b.Board[mFX][7] && mFY-mHY > 1 && mFY-mHY < 3 && b.Board[mFX][mHY] == 0 && mFX == mHX && b.Board[mFX][mHY+1] == 0 {
					MoveWhite = false
					CanMove = true
					b.Move2B = true
				} else if b.Board[mHX][mHY] > 0 && mFY-mHY == 1 && math.Abs(float64(mFX)-float64(mHX)) == 1 {
					MoveWhite = false
					CanMove = true
				} else if b.Move2W == true && b.Board[mFX][4] == PawnE && mFY-mHY < 2 && math.Abs(float64(mFX)-float64(mHX)) == 1 && mFY-mHY > 0 {
					b.Board[mHX][mHY+1] = Empty
					MoveWhite = false
					CanMove = true

				} else {
					CanMove = false
				}

			}
			if CanMove {
				b.Move2W = false
			}

		}
	case pieces == Knight || pieces == KnightE:
		dy := math.Abs(float64(mHY) - float64(mFY))
		dx := math.Abs(float64(mHX) - float64(mFX))
		if (dx == 1 && dy == 2) || (dx == 2 && dy == 1) {
			if b.White {

				if b.Board[mHX][mHY] == 99 || b.Board[mHX][mHY] > 0 {
					CanMove = false
				} else {
					CanMove = true
					MoveWhite = true
				}

				b.White = false
			} else {
				if b.Board[mHX][mHY] == 99 || b.Board[mHX][mHY] < 0 {
					CanMove = false
				} else {
					CanMove = true
					MoveWhite = false
				}

				b.White = true
			}
		}
	}
	if CanMove {

		if MoveWhite {
			b.Board[mHX][mHY] = b.Board[mFX][mFY]
			b.Board[mFX][mFY] = 0
			b.White = false
			return b, true
		} else {
			b.Board[mHX][mHY] = b.Board[mFX][mFY]
			b.Board[mFX][mFY] = 0
			b.White = true
			return b, true
		}
	} else {
		return b, false
	}

}
