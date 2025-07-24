package helpers

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

func HexToCell(hex string) (tcell.Color, error) {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	b, err := strconv.ParseInt(hex[4:6], 16, 64)

	if err != nil {
		return tcell.ColorDefault, err
	}

	return tcell.NewRGBColor(int32(r), int32(g), int32(b)), nil
}
