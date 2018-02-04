package rect

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
)

// XFEN is an X-FEN string
type XFEN string

// getPosLineTokens parses line as runes into string tokens
// it should be done especially for board with at least one of rect dimensions
// greater then 9 (in this case token may consist of one or two runes)
func getPosLineTokens(line string) []string {
	tokens := []string{}
	firstDigit := true
	for _, rune := range line {
		if !unicode.IsDigit(rune) {
			tokens = append(tokens, string(rune))
			firstDigit = true
			continue
		}

		// if unicode.IsDigit(rune)
		if firstDigit {
			tokens = append(tokens, string(rune))
			firstDigit = false
			continue
		}

		// if unicode.IsDigit(rune) && !firstDigit
		tokens[len(tokens)-1] += string(rune)
	}
	return tokens
}

// parseBoardWidth parses one line of posLines and returns a board width, it can be any of line
// (they should tokens in same int value) due to all horizontals (rows) have the same length == board width
func parseBoardWidth(line string) int {
	w := 0
	for _, token := range getPosLineTokens(line) {
		i, err := strconv.Atoi(token)
		if err == nil { // token is a number
			w += i
			continue
		}
		w++
	}
	return w
}

// parsePosLines parses lines containing FEN position parts (between '/' splitters) into pieces on a board
// this func changes board parameter
func parsePosLines(lines []string, board *Board) error {
	for y, line := range lines {
		x := 1
		for _, token := range getPosLineTokens(line) {
			i, err := strconv.Atoi(token)
			if err == nil { // token is a number
				x += i
				continue
			}

			coord, runeToken := Coord{x, board.Dim().(Coord).Y - y}, []rune(token)[0]

			colour := White
			if unicode.IsLower(runeToken) {
				colour = Black
			}

			f, exists := map[rune]func(Colour) base.IPiece{
				'p': NewPawn, 'n': NewKnight, 'b': NewBishop, 'r': NewRook,
				'q': NewQueen, 'a': NewArchbishop, 'c': NewChancellor, 'k': NewKing,
			}[unicode.ToLower(runeToken)]
			if !exists {
				return fmt.Errorf("invalid piece token: %s", token)
			}
			board.PlacePiece(coord, f(colour))

			// marking pieces moved as long as possible to detect it
			bh := board.Dim().(Coord).Y
			switch {
			case runeToken == 'p' && coord.Y != bh-1:
				board.Piece(coord).MarkMoved()
			case runeToken == 'P' && coord.Y != 2:
				board.Piece(coord).MarkMoved()
			default:
				if unicode.IsUpper(runeToken) && coord.Y != 1 || unicode.IsLower(runeToken) && coord.Y != bh {
					board.Piece(coord).MarkMoved()
				}
			}
			x++
		}
	}
	return nil
}

// parseSideToMove parses line into side to move colour
// this func changes board parameter
func parseSideToMove(line string, board *Board) error {
	if len(line) != 1 {
		return fmt.Errorf("invalid side to move: %s", line)
	}
	switch []rune(strings.ToLower(line))[0] {
	case 'w':
		board.SetSideToMove(White)
	case 'b':
		board.SetSideToMove(Black)
	default:
		return fmt.Errorf("invalid side to move token: %s", line)
	}
	return nil
}

// parseEP parses line into dst coords of a piece which can be EP-captured in board, with specified sideToMove
// this func changes board parameter
func parseEP(line string, sideToMove Colour, board *Board) error {
	if line == "-" {
		return nil
	}

	ep := NewLongAlgebraicNotation()
	if err := ep.DecodeCoord(line); err != nil {
		return err
	}

	epPieceColour, bh := sideToMove.Invert(), board.Dim().(Coord).Y
	step, limY := 1, bh-1
	if epPieceColour == Black {
		step, limY = -1, 2
	}

	// FEN has 'EP capture dst cell' coords while the board keeps 'piece to capture' coords

	epCoordX := ep.Coord.(Coord).X
	for y := ep.Coord.(Coord).Y + step; y != limY; y = y + step {
		coord := Coord{epCoordX, y}
		p := board.Piece(coord)
		if p != nil && p.Name() == base.PawnName {
			board.SetCanCaptureEnPassantAt(coord)
			return nil
		}
	}

	return fmt.Errorf("piece which can be EP-captured not found on board")
}

// parseCastling parses line about allowed castlings
// this func changes board parameter
func parseCastling(line string, board *Board) error {
	if line == "-" {
		return nil
	}

	bC := board.Dim().(Coord)
	// findRook finds rook of colour.
	// Set i to 0 for aSide rook finding, set i to 1 for zSide rook finding
	// Set outer to true to find outer rook (closer to board border), otherwise inner rook (closer to king)
	findRook := func(colour Colour, i int, outer bool) *Rook {
		rooks := board.FindPieces(base.PieceFilter{
			Names:   []string{base.RookName},
			Colours: []Colour{colour},
			Condition: func(r base.IPiece) bool {
				rC, y := r.Coord().(Coord), map[Colour]int{White: 1, Black: bC.Y}
				return rC.Y == y[colour]
			},
		})
		if rooks == nil || len(rooks) == 0 {
			return nil
		}

		if len(rooks) == 1 {
			return rooks[0].(*Rook)
		}

		if len(rooks) > 2 {
			panic("found more then two rooks on starting horizontal")
		}

		minRookX, maxRookX := bC.X+1, 0
		var aRook, zRook base.IPiece
		for j := range rooks {
			c := rooks[j].Coord().(Coord)
			if c.X < minRookX {
				minRookX, aRook = c.X, rooks[j]
			}
			if c.X > maxRookX {
				maxRookX, zRook = c.X, rooks[j]
			}
		}
		if outer && i == 1 || !outer && i == 0 {
			return zRook.(*Rook)
		}
		// if (outer && i == 0) || (!outer && i == 1)
		return aRook.(*Rook)
	}

	for _, token := range []rune(line) {
		colour := White
		if unicode.IsLower(token) {
			colour = Black
		}
		king := board.King(colour)
		if king == nil {
			return fmt.Errorf("king is not set while parseCastling() in XFEN")
		}
		outer, i, kC := strings.Contains("KkQq", string(token)), 0, king.Coord().(Coord)
		if strings.Contains("Kk", string(token)) || (!outer && FromLetter(token) > kC.X) {
			i = 1
		}
		r := findRook(colour, i, outer)
		if r == nil {
			return fmt.Errorf("wrong FEN, %s-castling specified, but rook not found", string(token))
		}
		board.SetRookInitialCoords(colour, i, r.Coord())
	}
	return nil
}

// PositionPart returns a position part of an X-FEN
func (s XFEN) PositionPart() string {
	return strings.Join(strings.Split(string(s), " ")[:4], " ")
}

// Board returns a new rectangular chess board position from standard X-FEN
func (s XFEN) Board() (base.IBoard, error) {
	xfenParts := strings.Split(string(s), " ")
	if len(xfenParts) != 6 {
		return nil, fmt.Errorf("invalid X-FEN length")
	}

	// xfenParts slice indexes:
	// 0 - position, 1 - side to move, 2 - castling rights, 3 - EP dst cell, 4 - half-moves counter, 5 - move number

	posLines := strings.Split(xfenParts[0], "/")
	bh := len(posLines)
	if bh < 3 {
		return nil, fmt.Errorf("board height is too small")
	}
	bw := parseBoardWidth(posLines[0])
	if bw < 3 {
		return nil, fmt.Errorf("board width is too small")
	}

	b := NewEmptyBoard(bw, bh, StandardChessBoardSettings())

	if err := parsePosLines(posLines, b); err != nil {
		return nil, err
	}

	if err := parseSideToMove(xfenParts[1], b); err != nil {
		return nil, err
	}

	if err := parseEP(xfenParts[3], b.SideToMove(), b); err != nil {
		return nil, err
	}

	if err := parseCastling(xfenParts[2], b); err != nil {
		return nil, err
	}

	halfMovesCount, err := strconv.Atoi(xfenParts[4])
	if err != nil {
		return nil, err
	}
	b.SetHalfMoveCount(halfMovesCount)
	moveNumber, err := strconv.Atoi(xfenParts[5])
	if err != nil {
		return nil, err
	}
	b.SetMoveNumber(moveNumber)
	b.ComputeOutcome()

	return b, nil
}

// NewXFEN converts rectangular board position to X-FEN
func NewXFEN(board *Board) XFEN {
	xfen := ""

	// converting position
	cells := board.Cells().(Cells)
	setCase := map[Colour]func(rune) rune{White: unicode.ToUpper, Black: unicode.ToLower}
	for y := range cells {
		empty := 0
		for x := range cells[y] {
			piece := cells[y][x].Piece()
			if piece == nil {
				empty++
				continue
			}
			if empty != 0 {
				xfen += strconv.Itoa(empty)
				empty = 0
			}
			xfen += string(setCase[piece.Colour()](piece.Capital()))
		}
		if empty != 0 {
			xfen += strconv.Itoa(empty)
		}
		xfen += "/"
	}
	xfen = xfen[:len(xfen)-1]

	sideToMove := board.SideToMove()

	// converting side to move
	xfen += " " + string(unicode.ToLower([]rune(sideToMove.Name())[0]))

	// converting castling flags
	bh, castlingFlags := board.Dim().(Coord).Y, ""
	for _, colour := range AllColours() {
		king := board.King(colour)
		if king == nil || king.WasMoved() {
			continue
		}
		rookInitialCoords := board.RookInitialCoords(colour)
		for i := len(rookInitialCoords) - 1; i >= 0; i-- {
			if rookInitialCoords[i] == nil {
				continue
			}
			r := board.Piece(rookInitialCoords[i])
			if r == nil || r.WasMoved() {
				continue
			}

			rC := r.Coord().(Coord)
			rooks := board.FindPieces(
				base.PieceFilter{
					Names:   []string{base.RookName},
					Colours: []Colour{colour},
					Condition: func(p base.IPiece) bool {
						pC, y := p.Coord().(Coord), map[Colour]int{White: 1, Black: bh}
						return pC.Y == y[colour] && (pC.X < rC.X && i == 0 || pC.X > rC.X && i == 1)
					},
				})
			castlingFlag := setCase[colour]([]rune{'q', 'k'}[i])
			if len(rooks) > 0 {
				castlingFlag = setCase[colour](ToLetter(rC.X))
			}
			castlingFlags += string(castlingFlag)
		}
	}
	if castlingFlags == "" {
		castlingFlags = "-"
	}
	xfen += " " + castlingFlags

	// converting en-passant capture coord
	canCaptureEP, epCaptureFEN := board.CanCaptureEnPassantAt(), "-"
	if canCaptureEP != nil {
		step := 1
		if sideToMove == Black {
			step = -1
		}
		epCoord, notation := canCaptureEP.(Coord), NewLongAlgebraicNotation()
		notation.SetCoord(Coord{epCoord.X, epCoord.Y + step})
		epCaptureFEN = notation.EncodeCoord()
	}
	xfen += " " + epCaptureFEN

	// converting counters
	xfen += fmt.Sprintf(" %d %d", board.HalfMoveCount(), board.MoveNumber())

	return XFEN(xfen)
}

// NewStandardChessStartingPosition returns XFEN for starting position for standard chess
func NewStandardChessStartingPosition() XFEN {
	return XFEN(`rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1`)
}
