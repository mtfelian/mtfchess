package xfen

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/mtfelian/utils"
)

// parsePosLines parses lines containing FEN position parts (between '/' splitters) into pieces on a board
// this func changes board parameter
func parsePosLines(lines []string, board *rect.Board) error {
	for y, line := range lines {
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

		x := 1
		for _, token := range tokens {
			i, err := strconv.Atoi(token)
			if err == nil { // token is a number
				x += i
				continue
			}

			c := rect.Coord{x, board.Dim().(rect.Coord).Y - y}
			runeToken := []rune(token)[0]
			switch runeToken {
			case 'P':
				board.PlacePiece(c, piece.NewPawn(White))
			case 'N':
				board.PlacePiece(c, piece.NewKnight(White))
			case 'B':
				board.PlacePiece(c, piece.NewBishop(White))
			case 'R':
				board.PlacePiece(c, piece.NewRook(White))
			case 'Q':
				board.PlacePiece(c, piece.NewQueen(White))
			case 'A': // todo board.PlacePiece(c, piece.NewArchbishop(White))
			case 'C': // todo board.PlacePiece(c, piece.NewChancellor(White))
			case 'K':
				board.PlacePiece(c, piece.NewKing(White))
				board.SetKing(White, board.Piece(c))
			case 'p':
				board.PlacePiece(c, piece.NewPawn(Black))
			case 'n':
				board.PlacePiece(c, piece.NewKnight(Black))
			case 'b':
				board.PlacePiece(c, piece.NewBishop(Black))
			case 'r':
				board.PlacePiece(c, piece.NewRook(Black))
			case 'q':
				board.PlacePiece(c, piece.NewQueen(Black))
			case 'a': // todo board.PlacePiece(c, piece.NewArchbishop(Black))
			case 'c': // todo board.PlacePiece(c, piece.NewChancellor(Black))
			case 'k':
				board.PlacePiece(c, piece.NewKing(Black))
				board.SetKing(Black, board.Piece(c))
			default:
				return fmt.Errorf("invalid piece token: %s", token)
			}
			x++

			// marking pieces moved as long as possible to detect it
			bh := board.Dim().(rect.Coord).Y
			switch runeToken {
			case 'p':
				if c.Y != bh-1 {
					board.Piece(c).MarkMoved()
				}
			case 'P':
				if c.Y != 2 {
					board.Piece(c).MarkMoved()
				}
			default:
				if unicode.IsUpper(runeToken) && c.Y != 1 {
					board.Piece(c).MarkMoved()
				}
				if unicode.IsLower(runeToken) && c.Y != bh {
					board.Piece(c).MarkMoved()
				}
			}

		}
	}
	return nil
}

// parseSideToMove parses line into side to move colour
func parseSideToMove(line string) Colour {
	return map[string]Colour{"w": White, "b": Black}[strings.ToLower(line)]
}

// parseEP parses line into dst coords of a piece which can be EP-captured in board, with specified sideToMove
// this func changes board parameter
func parseEP(line string, sideToMove Colour, board *rect.Board) error {
	if line == "-" {
		return nil
	}

	epCoord, err := rect.FromAlgebraic(line)
	if err != nil {
		return err
	}

	epPieceColour, bh := sideToMove.Invert(), board.Dim().(rect.Coord).Y
	step, limY := 1, bh-1
	if epPieceColour == Black {
		step, limY = -1, 2
	}

	// FEN has 'EP capture dst cell' coords while the board keeps 'piece to capture' coords

	epCoordX := epCoord.(rect.Coord).X
	for y := epCoord.(rect.Coord).Y + step; y != limY; y = y + step {
		coord := rect.Coord{epCoordX, y}
		p := board.Piece(coord)
		if p != nil && p.Name() == "pawn" {
			board.SetCanCaptureEnPassantAt(coord)
			return nil
		}
	}

	return fmt.Errorf("piece which can be EP-captured not found on board")
}

// parseCastling parses line about allowed castlings
// this func changes board parameter
func parseCastling(line string, board *rect.Board) error {
	bC := board.Dim().(rect.Coord)
	// findRook finds rook of colour.
	// Set i to 0 for aSide rook finding, set i to 1 for zSide rook finding
	// Set outer to true to find outer rook (closer to board border), otherwise inner rook (closer to king)
	findRook := func(colour Colour, i int, outer bool) *piece.Rook {
		rooks := board.FindPieces(base.PieceFilter{
			Names:   []string{"rook"},
			Colours: []Colour{colour},
			Condition: func(r base.IPiece) bool {
				rC, y := r.Coord().(rect.Coord), map[Colour]int{White: 1, Black: bC.Y}
				return rC.Y == y[colour]
			},
		})
		if rooks == nil || len(rooks) == 0 {
			return nil
		}

		if len(rooks) == 1 {
			return rooks[0].(*piece.Rook)
		}

		if len(rooks) > 2 {
			panic("found more then two rooks on starting horizontal")
		}

		minRookX, maxRookX := bC.X+1, 0
		var aRook, zRook base.IPiece
		for j := range rooks {
			c := rooks[j].Coord().(rect.Coord)
			if c.X < minRookX {
				minRookX, aRook = c.X, rooks[j]
			}
			if c.X > maxRookX {
				maxRookX, zRook = c.X, rooks[j]
			}
		}
		if outer && i == 1 || !outer && i == 0 {
			return zRook.(*piece.Rook)
		}
		// if (outer && i == 0) || (!outer && i == 1)
		return aRook.(*piece.Rook)
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
		outer := strings.Contains(string(token), "KkQq")
		kC := king.Coord().(rect.Coord)
		i := 0
		if strings.Contains(string(token), "Kk") || (!outer && rect.FromLetter(token) > kC.X) {
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

// NewFromStandardXFEN creates new chess board with pieces from standard XFEN
func NewFromStandardXFEN(fen string) (*rect.Board, error) {
	xfenParts := strings.Split(fen, " ")
	if len(xfenParts) != 6 {
		return nil, fmt.Errorf("invalid xfen length")
	}

	halfMovesCount, err := StringToUint(xfenParts[4])
	if err != nil {
		return nil, err
	}

	moveNumber, err := StringToUint(xfenParts[5])
	if err != nil {
		return nil, err
	}

	posLines := strings.Split(xfenParts[0], "/")
	bh := len(posLines)
	if bh < 3 {
		return nil, fmt.Errorf("bh is too small")
	}
	bw := len(posLines[0])
	if bw < 3 {
		return nil, fmt.Errorf("bw is too small")
	}

	b := rect.NewEmptyBoard(bw, bh, rect.StandardChessBoardSettings())

	if err := parsePosLines(posLines, b); err != nil {
		return nil, err
	}

	sideToMove := parseSideToMove(xfenParts[1])

	if err := parseEP(xfenParts[3], sideToMove, b); err != nil {
		return nil, err
	}

	if err := parseCastling(xfenParts[2], b); err != nil {
		return nil, err
	}

	_, _, _ = sideToMove, halfMovesCount, moveNumber

	// todo: side to move not implemented in board yet, tests on it, especially on castling parsing
	return b, nil
}
