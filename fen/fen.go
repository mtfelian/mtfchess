package fen

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mtfelian/mtfchess/base"
	. "github.com/mtfelian/mtfchess/colour"
	"github.com/mtfelian/mtfchess/piece"
	"github.com/mtfelian/mtfchess/rect"
	. "github.com/mtfelian/utils"
	"strconv"
)

// StandardFEN represents a parsed standard FEN data
type StandardFEN struct {
	Board      base.IBoard              // position
	SideToMove Colour                   // side to move
	Castlings  map[Colour]base.Castling // allowed castlings
	EnPassant  base.EPCapture
	// HalfMovesCount is a number of halfmoves since the last capture or pawn advance, to detect
	// 3-fold repetition or 50 moves draw rule
	HalfMovesCount uint
	MoveNumber     uint // moves counter
}

// NewFromStandardFEN creates new chess board with pieces from standard FEN
func NewFromStandardFEN(fen string) (*rect.Board, error) {
	fenParts := strings.Split(fen, " ")
	if len(fenParts) != 6 {
		return nil, fmt.Errorf("invalid fen length")
	}

	sPos, sSideToMove, sAllowedCastling := fenParts[0], fenParts[1], fenParts[2]
	sEnPassant, sHalfMovesCount, sMoveNumber := fenParts[3], fenParts[4], fenParts[5]

	halfMovesCount, err := StringToUint(sHalfMovesCount)
	if err != nil {
		return nil, err
	}

	moveNumber, err := StringToUint(sMoveNumber)
	if err != nil {
		return nil, err
	}

	posLines := strings.Split(sPos, "/")
	bh := len(posLines)
	if bh < 3 {
		return nil, fmt.Errorf("bh is too small")
	}
	bw := len(posLines[0])
	if bw < 3 {
		return nil, fmt.Errorf("bw is too small")
	}

	b := rect.NewEmptyBoard(bw, bh, rect.StandardChessBoardSettings())

	parsePosLines := func(lines []string, board *rect.Board) error {
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
				if err == nil { // token is a numbr
					x += i
					continue
				}

				c := rect.Coord{x, bh - y}
				switch token {
				case "P":
					board.PlacePiece(c, piece.NewPawn(White))
				case "N":
					board.PlacePiece(c, piece.NewKnight(White))
				case "B":
					board.PlacePiece(c, piece.NewBishop(White))
				case "R":
					board.PlacePiece(c, piece.NewRook(White))
				case "Q":
					board.PlacePiece(c, piece.NewQueen(White))
				case "A": // todo board.PlacePiece(c, piece.NewArchbishop(White))
				case "C": // todo board.PlacePiece(c, piece.NewChancellor(White))
				case "K":
					board.PlacePiece(c, piece.NewKing(White))
				case "p":
					board.PlacePiece(c, piece.NewPawn(Black))
				case "n":
					board.PlacePiece(c, piece.NewKnight(Black))
				case "b":
					board.PlacePiece(c, piece.NewBishop(Black))
				case "r":
					board.PlacePiece(c, piece.NewRook(Black))
				case "q":
					board.PlacePiece(c, piece.NewQueen(Black))
				case "a": // todo board.PlacePiece(c, piece.NewArchbishop(Black))
				case "c": // todo board.PlacePiece(c, piece.NewChancellor(Black))
				case "k":
					board.PlacePiece(c, piece.NewKing(Black))
				default:
					return fmt.Errorf("invalid piece token: %s", token)
				}
				x++
			}
		}

		return nil
	}

	if err := parsePosLines(posLines, b); err != nil {
		return nil, err
	}

	s := StandardFEN{
		Board:          b,
		SideToMove:     map[string]Colour{"w": White, "b": Black}[strings.ToLower(sSideToMove)],
		HalfMovesCount: halfMovesCount,
		MoveNumber:     moveNumber,
	}

	// todo: castling, enpassant parsing, side to move not implemented in board yet, tests on it
	// no need to return s, simply write all needed data to board

	_ = s
	_ = sAllowedCastling
	_ = sEnPassant

	return b, nil
}
