package model

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
	Board [10][10]int
	White bool
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
func PawnMove(b *Board, mFX, mFY, mHX, mHY int) (*Board, bool) {

	if b.White == true {
		if b.Board[mHX][mHY] > 0 || b.Board[mHX][mHY] == 99 {
			return b, false
		} else {
			b.Board[mHX][mHY] = b.Board[mFX][mFY]
			b.Board[mFX][mFY] = 0
		}
		b.White = false
	} else {
		if b.Board[mHX][mHY] < 0 || b.Board[mHX][mHY] == 99 {
			return b, false
		} else {
			b.Board[mHX][mHY] = b.Board[mFX][mFY]
			b.Board[mFX][mFY] = 0
		}
		b.White = true
	}
	return b, true
}
