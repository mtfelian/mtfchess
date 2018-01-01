package rect

import (
	"fmt"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/mtfelian/utils"
)

// NoPawnLongMoveFunc always disables pawn long forward move
func NoPawnLongMoveFunc(_ base.IBoard, _ base.IPiece) int { return 0 }

// StandardLongMoveFunc is a condition for pawn long forward move for standard chess
func StandardLongMoveFunc(board base.IBoard, piece base.IPiece) int {
	bh := board.Dim().(Coord).Y
	if (piece.Colour() == White && piece.Coord().(Coord).Y == 2) ||
		(piece.Colour() == Black && piece.Coord().(Coord).Y == bh-1) {
		return 1
	}
	return 0
}

// NoEnPassantFunc always disables en passant capturing
func NoEnPassantFunc(_ base.IBoard, _ base.IPiece) base.ICoords { return nil }

// StandardEnPassantFunc enables en passant capturing like in standard chess, returns coords to capture
func StandardEnPassantFunc(board base.IBoard, piece base.IPiece) base.ICoords {
	if piece.Name() != "pawn" {
		return nil
	}

	epData := board.CanCaptureEnPassant()
	if epData == nil {
		return nil
	}

	pX, epAtX := piece.Coord().(Coord).X, epData.From.(Coord).X
	if epAtX != pX+1 && epAtX != pX-1 {
		return nil
	}

	longMove := board.Settings().PawnLongMoveFunc(board, epData.Piece) // here is checked also piece Y coord
	if longMove == 0 {
		return nil
	}

	bh := board.Dim().(Coord).Y
	minYs := map[Colour]int{White: bh - 2 - longMove, Black: 3}
	maxYs := map[Colour]int{White: bh - 3, Black: 3 + longMove}

	step := 1
	if piece.Colour() == Black {
		step *= -1
	}

	res := NewCoords([]base.ICoord{})
	minY, maxY := minYs[piece.Colour()], maxYs[piece.Colour()]
	for y := minY; y >= minY && y <= maxY; y = y + step {
		res.Add(Coord{epAtX, y})
	}

	return res
}

// StandardAllowedPromotions returns allowed pawn promotions pieces names list for standard chess
func StandardAllowedPromotions() []string { return []string{"knight", "bishop", "rook", "queen"} }

// StandardPromotionConditionFunc is a pawn promotion condition for standard chess
func StandardPromotionConditionFunc(board base.IBoard, piece base.IPiece, dst base.ICoord, to base.IPiece) bool {
	bh, fromY, dstY := board.(*Board).height, piece.Coord().(Coord).Y, dst.(Coord).Y
	return piece.Name() == "pawn" && // only pawn can be promoted
		to.Colour() == piece.Colour() && // only to self-colored
		SliceContains(to.Name(), board.(*Board).Settings().AllowedPromotions) && // to piece from list
		(piece.Colour() == White && fromY == bh-1 && dstY == bh || // for white from pre-last horizontal to the last
			piece.Colour() == Black && fromY == 2 && dstY == 1) // for black from 2nd horizontal to the 1st
}

// standardCastling returns castling data for standard chess for given colour on a given board
// set aSide to true to return a-side castling, otherwise to return h-side castling
func standardCastling(board base.IBoard, colour Colour, aSide bool) base.Castling {
	res, bh := base.Castling{Enabled: false}, board.Dim().(Coord).Y

	king := board.King(colour)
	if king == nil || king.WasMoved() {
		return res
	}

	if board.InCheck(colour) {
		return res
	}

	kC := king.Coord().(Coord)
	rooks := board.FindPieces(base.PieceFilter{
		Names:   []string{"rook"},
		Colours: []Colour{colour},
		Condition: func(r base.IPiece) bool {
			rC, y := r.Coord().(Coord), map[Colour]int{White: 1, Black: bh}
			return ((rC.X < kC.X && aSide) || (rC.X > kC.X && !aSide)) && rC.Y == y[colour] && kC.Y == y[colour]
		},
	})

	if rooks == nil || len(rooks) != 1 || rooks[0].WasMoved() {
		return res
	}

	// king and rook destination coordinates after castling
	kDstX, rDstX, xStep := 7, 6, 1
	if aSide {
		kDstX, rDstX, xStep = 3, 4, -1
	}
	kingDstCoord := map[Colour]Coord{White: {kDstX, 1}, Black: {kDstX, bh}}
	rookDstCoord := map[Colour]Coord{White: {rDstX, 1}, Black: {rDstX, bh}}

	// checking that king's path from source cell to destination cell is not attacked and free of pieces
	attacked := board.FindAttackedCellsBy(base.PieceFilter{Colours: []Colour{colour.Invert()}})
	//fmt.Println(colour.Invert(), attacked)
	for i := xStep; kC.X+i != kingDstCoord[colour].X+xStep; i += xStep {
		shouldBeFree := Coord{kC.X + i, kC.Y}
		_ = fmt.Sprint("")
		//fmt.Println(colour, i, shouldBeFree, attacked.Contains(shouldBeFree), board.Piece(shouldBeFree) != nil)
		if attacked.Contains(shouldBeFree) || board.Piece(shouldBeFree) != nil {
			//fmt.Println("!! xStep=", xStep, i, kC.X)
			//fmt.Println(">>", colour, kingDstCoord[colour].X)
			return res
		}
	}

	return base.Castling{
		Piece:   [2]base.IPiece{king, rooks[0]},
		To:      [2]base.ICoord{kingDstCoord[colour], rookDstCoord[colour]},
		Enabled: true,
	}
}

// StandardCastlingFunc is a castling func for standard chess
func StandardCastlingFunc(board base.IBoard, colour Colour) base.Castlings {
	castlings := base.Castlings{standardCastling(board, colour, true), standardCastling(board, colour, false)}
	res := base.Castlings{}
	for i := range castlings {
		if castlings[i].Enabled {
			res = append(res, castlings[i])
		}
	}
	return res
}

// NoCastlingFunc is a castling func which disables castling
func NoCastlingFunc(_ base.IBoard, _ Colour) base.Castlings { return base.Castlings{} }
