package rect

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/mtfelian/mtfchess/base"
)

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

	x := int(parts[1][0] - 'a' + 1)
	y, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	return Coord{x, y}, nil
}
