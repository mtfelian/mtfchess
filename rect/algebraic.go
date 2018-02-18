package rect

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/mtfelian/mtfchess/base"
)

const (
	longAlgebraic = iota
)

const (
	moveDelimiter    = "-"
	captureDelimiter = "x"
)

const (
	noPostfix        = ""
	checkPostfix     = "+"
	checkmatePostfix = "#"
)

const (
	aSideCastling = "O-O-O"
	zSideCastling = "O-O"
)

var (
	castlingRegexp           = regexp.MustCompile(`(?i)^(O-O(?:-O)?)[+#]?$`)
	longAlgebraicCoordRegexp = regexp.MustCompile(`^([a-z])(\d{1,2})$`)
	longAlgebraicMoveRegexp  = regexp.MustCompile(`^([a-z]\d{1,2})[-x]([a-z]\d{1,2})[+#]?$`)
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

// ToLetter returns letter from the given x coord
func ToLetter(x int) rune { return 'a' - 1 + rune(x) }

// SetCoords sets notation coord to
func (n *algebraicNotation) SetCoord(to base.ICoord) base.INotation {
	n.Coord = to
	return n
}

// DecodeMove returns a func that tries to make a decoded move (or castling) on a board
func (n *algebraicNotation) DecodeMove(board base.IBoard, move string) (func() bool, error) {

	// move is a castling
	re := castlingRegexp.Copy()
	if re.MatchString(move) {
		move = strings.ToUpper(move)

		parts := re.FindStringSubmatch(move)
		if len(parts) != 2 {
			return nil, fmt.Errorf("wrong casling move format: %s", move)
		}

		castlings := board.Castlings(board.SideToMove())
		castlingStrings := []string{aSideCastling, zSideCastling}
		for i := range castlings {
			if parts[1] == castlingStrings[castlings[i].I] {
				return func() bool { return board.MakeCastling(castlings[i]) }, nil
			}
		}
		return nil, fmt.Errorf("this castling move is not available")
	}

	// move is not a castling

	move, re = strings.ToLower(move), longAlgebraicMoveRegexp.Copy()
	if !re.MatchString(move) {
		return nil, fmt.Errorf("wrong move format: %s", move)
	}

	parts := re.FindStringSubmatch(move)
	if len(parts) != 3 {
		return nil, fmt.Errorf("wrong move format: %s", move)
	}

	if err := n.DecodeCoord(parts[1]); err != nil {
		return nil, err
	}
	fromCoord := n.Coord.Copy()
	if err := n.DecodeCoord(parts[2]); err != nil {
		return nil, err
	}
	toCoord := n.Coord.Copy()

	return func() bool { return board.MakeMove(toCoord, board.Piece(fromCoord)) }, nil
}

// EncodeMove on board with piece to dst coord
func (n *algebraicNotation) EncodeMove(board base.IBoard, piece base.IPiece, dst base.ICoord) string {
	anFrom := NewLongAlgebraicNotation().SetCoord(piece.Coord())
	anTo := NewLongAlgebraicNotation().SetCoord(dst)
	delimiter := moveDelimiter
	if board.Piece(dst) != nil {
		delimiter = captureDelimiter
	}

	fig := string(piece.Capital())
	if piece.Name() == base.PawnName {
		fig = ""
	}

	projection := board.Project(piece, dst)
	projection.SetSideToMove(projection.SideToMove().Invert())

	check := noPostfix
	if projection.InCheckmate(projection.SideToMove()) {
		check = checkmatePostfix
		return fig + anFrom.EncodeCoord() + delimiter + anTo.EncodeCoord() + check
	}

	if projection.InCheck(projection.SideToMove()) {
		check = checkPostfix
	}

	return fig + anFrom.EncodeCoord() + delimiter + anTo.EncodeCoord() + check
}

// EncodeCastling on board
func (n *algebraicNotation) EncodeCastling(i int) string {
	if i == 0 {
		return aSideCastling
	}
	return zSideCastling
}

// DecodeCoord coord string (case-insensitive) to (x,y) coords
func (n *algebraicNotation) DecodeCoord(coord string) error {
	coord = strings.ToLower(coord)
	re := longAlgebraicCoordRegexp.Copy()
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
	return fmt.Sprintf("%c%d", ToLetter(c.X), c.Y)
}
