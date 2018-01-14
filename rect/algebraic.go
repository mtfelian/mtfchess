package rect

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
)

const (
	longAlgebraic = iota
)

// algebraicNotation implementation for INotation
type algebraicNotation struct {
	Coord base.ICoord
	mode  int
}

// NewLongAlgebraicNotation returns new long algebraic notation
func NewLongAlgebraicNotation() *algebraicNotation { return &algebraicNotation{mode: longAlgebraic} }

// FromLetter returns x coord from the given letter
func FromLetter(letter rune) int { return int(unicode.ToLower(letter) - 'a' + 1) }

// SetCoords sets notation coord to
func (n *algebraicNotation) SetCoord(to base.ICoord) base.INotation {
	n.Coord = to
	return n
}

// EncodeMove on board with piece to dst coord
func (n *algebraicNotation) EncodeMove(board base.IBoard, piece base.IPiece, dst base.ICoord) string {
	anFrom := NewLongAlgebraicNotation().SetCoord(piece.Coord())
	anTo := NewLongAlgebraicNotation().SetCoord(dst)
	delimiter := "-"
	if board.Piece(dst) != nil {
		delimiter = "x"
	}

	fig := piece.Capital()
	if piece.Name() == base.PawnName {
		fig = ""
	}

	projection := board.Project(piece, dst)
	projection.SetSideToMove(projection.SideToMove().Invert())

	check := ""
	if projection.InCheckMate(projection.SideToMove()) {
		check = "#"
		return fmt.Sprintf("%s%s%s%s%s", fig, anFrom.EncodeCoord(), delimiter, anTo.EncodeCoord(), check)
	}

	if projection.InCheck(projection.SideToMove()) {
		check = "+"
	}

	return fmt.Sprintf("%s%s%s%s%s", fig, anFrom.EncodeCoord(), delimiter, anTo.EncodeCoord(), check)
}

// EncodeCastling on board for sideToMove
func (n *algebraicNotation) EncodeCastling(board base.IBoard, sideToMove Colour, i int) string {
	switch {
	case board.RookCanCastle(sideToMove, 0):
		return "O-O-O"
	case board.RookCanCastle(sideToMove, 1):
		return "O-O"
	}
	return ""
}

// DecodeCoord coord string (case-insensitive) to (x,y) coords
func (n *algebraicNotation) DecodeCoord(coord string) error {
	coord = strings.ToLower(coord)
	re := regexp.MustCompile(`^([a-z])(\d{1,2})$`)
	if !re.MatchString(coord) {
		return fmt.Errorf("wrong coord format: %s", coord)
	}

	parts := re.FindStringSubmatch(coord)
	if len(parts) != 3 {
		return fmt.Errorf("wrong coord format: %s", coord)
	}

	x := FromLetter([]rune(parts[1])[0])
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return err
	}

	n.Coord = Coord{x, y}
	return nil
}

// EncodeCoord n.Coord as string
func (n *algebraicNotation) EncodeCoord() string {
	if n.Coord == nil {
		return ""
	}
	c := n.Coord.(Coord)
	return fmt.Sprintf("%s%d", string('a'-1+c.X), c.Y)
}
