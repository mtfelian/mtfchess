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
func parseEP(line string, sideToMove Colour, board *rect.Board) (base.ICoord, error) {
	if line == "-" {
		return nil, nil
	}

	epCoord, err := rect.FromAlgebraic(line)
	if err != nil {
		return nil, err
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
			return coord, nil
		}
	}

	return nil, fmt.Errorf("piece which can be EP-captured not found on board")
}

// parseCastling line about allowed castlings
// this func changes board parameter
func parseCastling(line string, board *rect.Board) error {
	bC := board.Dim().(rect.Coord)
	// findRook finds rook of colour.
	// Set border to true to find rook nearest board border, otherwise set it to false.
	findRook := func(colour Colour, border, leftMost bool) *piece.Rook {
		king := board.King(colour)
		if king == nil {
			panic("king is not set while parseCastling() in XFEN")
		}

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

		isBorder := make([]bool, len(rooks))
		for i := range rooks {
			isBorder[i] = true
			c := rooks[i].Coord().(rect.Coord)
			if leftMost {
				for x := c.X - 1; x >= 1; x-- {
					pieceAt := board.Piece(rect.Coord{x, c.Y})
					if pieceAt != nil && (pieceAt.Name() == "rook" || pieceAt.Name() == "king") {
						isBorder[i] = false
					}
				}
			} else {
				for x := c.X + 1; x <= bC.X; x++ {
					pieceAt := board.Piece(rect.Coord{x, c.Y})
					if pieceAt != nil && (pieceAt.Name() == "rook" || pieceAt.Name() == "king") {
						isBorder[i] = false
					}
				}
			}
			if isBorder[i] == border {
				return rooks[i].(*piece.Rook)
			}
		}

		return rooks[0].(*piece.Rook)
	}

	if strings.Contains(line, "K") {
		r := findRook(White, true)
		if r == nil {
			panic("wrong FEN, K-castling specified, but rook not found")
		}
		board.SetRookInitialCoords(White, 1, r.Coord())
	}
	if strings.Contains(line, "k") {
		r := findRook(Black, true)
		if r == nil {
			panic("wrong FEN, k-castling specified, but rook not found")
		}
		board.SetRookInitialCoords(Black, 1, r.Coord())
	}
	if strings.Contains(line, "Q") {
		r := findRook(White, true)
		if r == nil {
			panic("wrong FEN, Q-castling specified, but rook not found")
		}
		board.SetRookInitialCoords(Black, 0, r.Coord())
	}
	if strings.Contains(line, "q") {
		r := findRook(Black, true)
		if r == nil {
			panic("wrong FEN, q-castling specified, but rook not found")
		}
		board.SetRookInitialCoords(Black, 0, r.Coord())
	}
}

// StandardXFEN represents a parsed standard FEN data
type StandardXFEN struct {
	Board              base.IBoard     // position
	SideToMove         Colour          // side to move
	RookCoords         base.RookCoords // rook coords
	EnPassantCaptureAt base.ICoord
	// HalfMovesCount is a number of halfmoves since the last capture or pawn advance, to detect
	// 3-fold repetition or 50 moves draw rule
	HalfMovesCount uint
	MoveNumber     uint // moves counter
}

// NewFromStandardXFEN creates new chess board with pieces from standard FEN
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

	ep, err := parseEP(xfenParts[3], sideToMove, b)
	if err != nil {
		return nil, err
	}

	s := StandardXFEN{
		Board:              b,
		EnPassantCaptureAt: ep,
		SideToMove:         sideToMove,
		HalfMovesCount:     halfMovesCount,
		MoveNumber:         moveNumber,
	}

	// todo: castling, side to move not implemented in board yet, tests on it
	// no need to return s, simply write all needed data to board
	_ = s
	//todo xfenParts[2], allowedCastling

	return b, nil
}
