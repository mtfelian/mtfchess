package base

import (
	"fmt"

	. "github.com/mtfelian/mtfchess/colour"
)

// reason is an outcome reason
type reason int

const (
	notCompleted reason = iota
	checkmate
	timeOver
	stalemate
	drawByAgreement
	drawBy3FoldRepetition
	drawBy50MovesRule
	drawByNotSufficientMaterial
)

// Outcome is a game outcome
type Outcome struct {
	Winner Colour
	Reason reason
}

// Equals returns true if o equals to
func (o Outcome) Equals(to Outcome) bool { return o.Winner == to.Winner && o.Reason == to.Reason }

// String makes Outcome to implement fmt.Stringer
func (o Outcome) String() string {
	switch o.Reason {
	case notCompleted:
		return "Game is in progress"
	case checkmate:
		return fmt.Sprintf("%s won by checkmate", o.Winner.Name())
	case timeOver:
		return fmt.Sprintf("%s lost by time over", o.Winner.Invert().Name())
	case stalemate:
		return "Stalemate"
	case drawByAgreement:
		return "Draw by agreement"
	case drawBy3FoldRepetition:
		return "Draw by 3-fold repetition"
	case drawBy50MovesRule:
		return "Draw by 50 moves rule"
	case drawByNotSufficientMaterial:
		return "Draw by no sufficient material"
	}
	return ""
}

// NewOutcomeNotCompleted returns an incomplete game outcome
func NewOutcomeNotCompleted() Outcome { return Outcome{Winner: Transparent, Reason: notCompleted} }

// NewStalemate returns an outcome for stalemate
func NewStalemate() Outcome { return Outcome{Winner: Transparent, Reason: stalemate} }

// NewCheckmate returns an outcome for checkmate by winner
func NewCheckmate(winner Colour) Outcome { return Outcome{Winner: winner, Reason: checkmate} }

// NewTimeOver returns an outcome for time over for side
func NewTimeOver(side Colour) Outcome { return Outcome{Winner: side.Invert(), Reason: timeOver} }

// NewDrawByAgreement returns an outcome for draw by agreement
func NewDrawByAgreement() Outcome { return Outcome{Winner: Transparent, Reason: drawByAgreement} }

// NewDrawBy3FoldRepetition returns an outcome for draw by 3-fold repetition
func NewDrawBy3FoldRepetition() Outcome {
	return Outcome{Winner: Transparent, Reason: drawBy3FoldRepetition}
}

// NewDrawBy50MovesRule returns an outcome for draw by 50 moves rule
func NewDrawBy50MovesRule() Outcome { return Outcome{Winner: Transparent, Reason: drawBy50MovesRule} }

// NewDrawByNotSufficientMaterial returns an outcome for draw due to not sufficient material
func NewDrawByNotSufficientMaterial() Outcome {
	return Outcome{Winner: Transparent, Reason: drawByNotSufficientMaterial}
}
