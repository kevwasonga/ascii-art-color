package utils

import (
	"errors"
	"fmt"
)

func RGBToANSI(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}



func FindColorANSI(color string) (string, error) {
	rgb, exists := colorMap[color]
	if !exists {
		return "", errors.New("color not available")
	}
	return RGBToANSI(rgb[0], rgb[1], rgb[2]), nil
}
