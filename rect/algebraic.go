package rect

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/mtfelian/mtfchess/base"
)

// FromLetter returns x coord from the given letter
func FromLetter(letter rune) int { return int(unicode.ToLower(letter) - 'a' + 1) }

// FromAlgebraic makes (x,y) coords from algebraic form.
// coord is case-insensitive
func FromAlgebraic(coord string) (base.ICoord, error) {
	coord = strings.ToLower(coord)
	re := regexp.MustCompile(`^([a-z])(\d{1,2})$`)
	if !re.MatchString(coord) {
		return nil, fmt.Errorf("wrong coord format: %s", coord)
	}

	parts := re.FindStringSubmatch(coord)
	if len(parts) != 3 {
		return nil, fmt.Errorf("wrong coord format: %s", coord)
	}

	x := FromLetter([]rune(parts[1])[0])
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	return Coord{x, y}, nil
}
