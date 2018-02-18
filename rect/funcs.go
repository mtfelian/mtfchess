package rect

import (
	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	. "github.com/mtfelian/utils"
)

const (
	NoPawnLongMove       = 0 // disable pawn long move
	StandardPawnLongMove = 1 // pawn from starting rank can go 1 vert. cell further
)

const (
	NoMovesToDraw         = 0  // disable N moves draw rule
	Standard50MovesToDraw = 50 // 50 moves draw rule
)

const (
	NoXFoldRepetitionDraw       = 0 // disable X-fold repetition draw rule
	Standard3FoldRepetitionDraw = 3 // 3-fold repetition draw rule
)

// StandardChessBoardSettings returns a set of settings for standard chess
func StandardChessBoardSettings() *base.Settings {
	return &base.Settings{
		PawnLongMoveModifier:   StandardPawnLongMove,
		AllowedPromotions:      StandardAllowedPromotions(),
		PromotionConditionFunc: StandardPromotionConditionFunc,
		CastlingsFunc:          StandardCastlingFunc,
		EnPassantFunc:          StandardEnPassantFunc,
		MoveOrder:              true,
		MovesToDraw:            Standard50MovesToDraw,
		PositionsToDraw:        Standard3FoldRepetitionDraw,
	}
}

// testBoardSettings returns a set of settings for tests
func testBoardSettings() *base.Settings {
	return &base.Settings{
		PawnLongMoveModifier:   NoPawnLongMove,
		AllowedPromotions:      StandardAllowedPromotions(),
		PromotionConditionFunc: StandardPromotionConditionFunc,
		CastlingsFunc:          NoCastlingFunc,
		EnPassantFunc:          NoEnPassantFunc,
		MoveOrder:              false,
		MovesToDraw:            NoMovesToDraw,
		PositionsToDraw:        NoXFoldRepetitionDraw,
	}
}

// NoEnPassantFunc always disables en passant capturing
func NoEnPassantFunc(_ base.IBoard, _ base.IPiece) base.ICoord { return nil }

// StandardEnPassantFunc enables en passant capturing like in standard chess, returns coords to capture
// piece is a capturing piece
func StandardEnPassantFunc(board base.IBoard, piece base.IPiece) base.ICoord {
	if piece.Name() != base.PawnName {
		return nil
	}

	// pieceAt is a coord of a piece which can be captured en passant
	pieceAt := board.CanCaptureEnPassantAt()
	if pieceAt == nil {
		return nil
	}

	pX, epAtX := piece.Coord().(Coord).X, pieceAt.(Coord).X
	if epAtX != pX+1 && epAtX != pX-1 { // here we check X
		return nil
	}

	longMove := board.Settings().PawnLongMoveModifier // here is checked also piece Y coord
	if longMove == 0 {
		return nil
	}

	bh := board.Dim().(Coord).Y
	// maps colour of a capturing piece to Y of the piece to capture
	minYs := map[Colour]int{White: bh - 2 - longMove, Black: 3}
	maxYs := map[Colour]int{White: bh - 3, Black: 2 + longMove}

	step := 1
	if piece.Colour() == Black {
		step *= -1
	}

	minY, maxY := minYs[piece.Colour()], maxYs[piece.Colour()]
	for y := minY; y >= minY && y <= maxY; y = y + step { // here we check Y
		if y == piece.Coord().(Coord).Y+step {
			return Coord{epAtX, y}
		}
	}

	return nil
}

// StandardAllowedPromotions returns allowed pawn promotions pieces names list for standard chess
func StandardAllowedPromotions() []string {
	return []string{base.KnightName, base.BishopName, base.RookName, base.QueenName}
}

// StandardPromotionConditionFunc is a pawn promotion condition for standard chess
func StandardPromotionConditionFunc(board base.IBoard, piece base.IPiece, dst base.ICoord, to base.IPiece) bool {
	bh, fromY, dstY := board.(*Board).height, piece.Coord().(Coord).Y, dst.(Coord).Y
	return piece.Name() == base.PawnName && // only pawn can be promoted
		to.Colour() == piece.Colour() && // only to self-colored
		SliceContains(to.Name(), board.(*Board).Settings().AllowedPromotions) && // to piece from list
		(piece.Colour() == White && fromY == bh-1 && dstY == bh || // for white from pre-last horizontal to the last
			piece.Colour() == Black && fromY == 2 && dstY == 1) // for black from 2nd horizontal to the 1st
}

// standardCastling returns castling data for standard chess for given colour on a given board
// parameter x is a rook coord
func standardCastling(board base.IBoard, colour Colour, rook base.IPiece) base.Castling {
	res, bh := base.Castling{Enabled: false}, board.Dim().(Coord).Y

	king := board.King(colour)
	if king == nil || king.WasMoved() || board.InCheck(colour) {
		return res
	}

	// detect, whether it aSide or zSide castling
	rooksCoords, n := board.RookInitialCoords(colour), -1
	if len(rooksCoords) != 2 {
		panic("rookCoords should have len 2")
	}
	rC := rook.Coord().(Coord)
	for i := range rooksCoords {
		if rooksCoords[i] == nil {
			continue
		}
		if rooksCoords[i].(Coord).X == rC.X {
			n = i
			break
		}
	}
	if n == -1 { // rook not found
		return res
	}

	// kDstX and rDstX is a king and rook destination X after castling
	kDstX, rDstX := 3, 4 // aSide castling
	if n == 1 {
		kDstX, rDstX = 7, 6 // zSide castling
	}

	// checking that king's path from source cell to destination cell is not attacked and free of pieces
	// except the same rook
	kC, step := king.Coord().(Coord), 1 // go right if king now is to the left from the dst
	if kDstX < kC.X {                   // go left if king now is to the right from the dst
		step = -1
	}
	attacked := board.FindAttackedCellsBy(base.PieceFilter{Colours: []Colour{colour.Invert()}})
	for i := step; kC.X+i != kDstX+step; i += step {
		shouldBeFree := Coord{kC.X + i, kC.Y}
		if attacked.Contains(shouldBeFree) ||
			(board.Piece(shouldBeFree) != nil && !board.Piece(shouldBeFree).Coord().Equals(rook.Coord())) {
			return res
		}
	}

	// checking that rook's path from source cell to destination cell is free of pieces except king
	step = 1          // go right if rook now is to the left from the dst
	if rDstX < rC.X { // go left if rook now is to the right from the dst
		step = -1
	}
	for i := step; rC.X+i != rDstX+step; i += step {
		shouldBeFree := Coord{rC.X + i, rC.Y}
		if board.Piece(shouldBeFree) != nil && !board.Piece(shouldBeFree).Coord().Equals(king.Coord()) {
			return res
		}
	}

	kingDstCoord := map[Colour]Coord{White: {kDstX, 1}, Black: {kDstX, bh}}
	rookDstCoord := map[Colour]Coord{White: {rDstX, 1}, Black: {rDstX, bh}}
	return base.Castling{
		Piece:   [2]base.IPiece{king, rook},
		To:      [2]base.ICoord{kingDstCoord[colour], rookDstCoord[colour]},
		I:       n,
		Enabled: true,
	}
}

// StandardCastlingFunc is a castling func for standard chess
func StandardCastlingFunc(board base.IBoard, colour Colour) base.Castlings {
	bh := board.Dim().(Coord).Y
	rooks := board.FindPieces(base.PieceFilter{
		Names:   []string{base.RookName},
		Colours: []Colour{colour},
		Condition: func(r base.IPiece) bool {
			rC, y := r.Coord().(Coord), map[Colour]int{White: 1, Black: bh}
			rCoords, boardAllowed := board.RookInitialCoords(colour), false
			for i := range rCoords {
				if rCoords[i] != nil && rCoords[i].Equals(r.Coord()) {
					boardAllowed = true
					break
				}
			}
			return rC.Y == y[colour] && boardAllowed && !r.WasMoved()
		},
	})

	castlings := base.Castlings{}
	for i := range rooks {
		castlings = append(castlings, standardCastling(board, colour, rooks[i]))
	}
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
