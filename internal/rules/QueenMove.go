package rules

import (
	"ChessEngineEOF/internal/model"
	"math"
)

func QueenMove(b *model.Board, p *model.PiecesMove) (bool, bool) {
	CanMove := true
	MoveWhite := false
	if b.White {
		if b.Board[p.MHX][p.MHY] == 99 || b.Board[p.MHX][p.MHY] > 0 {
			CanMove = false
			MoveWhite = true
		} else {
			if math.Abs(float64(p.MHX)-float64(p.MFX)) == math.Abs(float64(p.MHY)-float64(p.MFY)) {

				var stepY int
				var stepX int

				ClearPath := true
				// вверх налево
				if p.MHY > p.MFY && p.MFX > p.MHX {
					stepY = 1
					stepX = -1

				} else if p.MHY > p.MFY && p.MHX > p.MFX { // вверх направо
					stepX = 1
					stepY = 1
				} else if p.MHY < p.MFY && p.MHX < p.MFX { // вниз влево
					stepX = -1
					stepY = -1
				} else if p.MFY > p.MHY && p.MFX < p.MHX { //вниз направо
					stepY = -1
					stepX = 1
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

			} else if math.Abs(float64(p.MHX)-float64(p.MFX)) != math.Abs(float64(p.MHY)-float64(p.MFY)) {
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
		}

	} else {
		if b.Board[p.MHX][p.MHY] == 99 || b.Board[p.MHX][p.MHY] < 0 {
			CanMove = false
			MoveWhite = false
		}
		if math.Abs(float64(p.MHX)-float64(p.MFX)) == math.Abs(float64(p.MHY)-float64(p.MFY)) {

			var stepY int
			var stepX int

			ClearPath := true
			// вверх налево
			if p.MHY > p.MFY && p.MFX > p.MHX {
				stepY = 1
				stepX = -1

			} else if p.MHY > p.MFY && p.MHX > p.MFX { // вверх направо
				stepX = 1
				stepY = 1
			} else if p.MHY < p.MFY && p.MHX < p.MFX { // вниз влево
				stepX = -1
				stepY = -1
			} else if p.MFY > p.MHY && p.MFX < p.MHX { //вниз направо
				stepY = -1
				stepX = 1
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

		} else if math.Abs(float64(p.MHX)-float64(p.MFX)) != math.Abs(float64(p.MHY)-float64(p.MFY)) {
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
		}
	}
	return CanMove, MoveWhite
}
