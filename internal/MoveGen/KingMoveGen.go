package movegen

import "ChessEngineEOF/internal/model"

func KingMoveGen(b *model.BoardP, x, y int) []model.PiecesMove {
	Moves := []model.PiecesMove{}
	if b.White {
		forminusminus(&Moves, x, y, b, 1)
		forminusplus(&Moves, x, y, b, 1)
		forplusminus(&Moves, x, y, b, 1)
		forplusplus(&Moves, x, y, b, 1)
		forUp(&Moves, x, y, b, 1)
		forRight(&Moves, x, y, b, 1)
		forLeft(&Moves, x, y, b, 1)
		forDown(&Moves, x, y, b, 1)
	} else {
		forminusminusB(&Moves, x, y, b, 1)
		forminusplusB(&Moves, x, y, b, 1)
		forplusminusB(&Moves, x, y, b, 1)
		forplusplusB(&Moves, x, y, b, 1)
		forUpB(&Moves, x, y, b, 1)
		forRightB(&Moves, x, y, b, 1)
		forLeftB(&Moves, x, y, b, 1)
		forDownB(&Moves, x, y, b, 1)
	}

	return Moves
}
