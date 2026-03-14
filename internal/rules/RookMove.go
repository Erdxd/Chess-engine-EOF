package rules

import "ChessEngineEOF/internal/model"

func RookMove(b *model.BoardP, p *model.PiecesMove) (bool, bool) {
	CanMove := false
	MoveWhite := false
	if b.White {
		if b.Board[p.MHX][p.MHY] == 99 || b.Board[p.MHX][p.MHY] > 0 {
			CanMove = false
			MoveWhite = true
		} else {
			var stepY int
			var stepX int

			ClearPath := true
			if p.MHY > p.MFY && p.MFX == p.MHX {
				stepY = 1
			} else if p.MHY == p.MFY && p.MHX > p.MFX {
				stepX = 1
			} else if p.MHY == p.MFY && p.MFX > p.MHX {
				stepX = -1
			} else if p.MHY < p.MFY && p.MFX == p.MHX {
				stepY = -1
			}
			currX := p.MFX + stepX
			currY := p.MFY + stepY
			for currX != p.MHX || currY != p.MHY {
				if b.Board[currX][currY] != 0 {
					ClearPath = false
					break
				}
				currX += stepX
				currY += stepY
			}
			if b.Board[p.MHX][p.MHY] < 0 && ClearPath == true {

				CanMove = true
				MoveWhite = true
			} else if ClearPath == true {
				CanMove = true
				MoveWhite = true
			}

		}
		b.White = false
	} else {
		if b.Board[p.MHX][p.MHY] == 99 || b.Board[p.MHX][p.MHY] < 0 {
			CanMove = false
			MoveWhite = false
		} else {
			var stepY int
			var stepX int

			ClearPath := true
			if p.MHY > p.MFY && p.MFX == p.MHX {
				stepY = 1
			} else if p.MHY == p.MFY && p.MHX > p.MFX {
				stepX = 1
			} else if p.MHY == p.MFY && p.MFX > p.MHX {
				stepX = -1
			} else if p.MHY < p.MFY && p.MFX == p.MHX {
				stepY = -1
			}
			currX := p.MFX + stepX
			currY := p.MFY + stepY
			for currX != p.MHX || currY != p.MHY {
				if b.Board[currX][currY] != 0 {
					ClearPath = false
					break
				}
				currX += stepX
				currY += stepY
			}
			if b.Board[p.MHX][p.MHY] < 0 && ClearPath == true {

				CanMove = true
				MoveWhite = false
			} else if ClearPath == true {
				CanMove = true
				MoveWhite = false
			}
			b.White = true
		}

	}
	return CanMove, MoveWhite
}
