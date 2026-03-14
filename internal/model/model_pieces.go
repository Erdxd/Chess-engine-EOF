package model

type BoardP struct {
	Board  [10][10]int
	White  bool
	Move2W bool
	Move2B bool
}
type PiecesMove struct {
	MFX   int
	MFY   int
	MHX   int
	MHY   int
	Piece int
}
