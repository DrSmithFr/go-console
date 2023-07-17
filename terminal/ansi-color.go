package terminal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type AnsiColorMode string

const (
	/*
	 * Classical 4-bit Ansi colors, including 8 classical colors and 8 bright color. Output syntax is "ESC[${foreGroundColorcode};${backGroundColorcode}m"
	 * Must be compatible with all terminals and it's the minimal version supported.
	 */
	Ansi4 AnsiColorMode = "Ansi4"

	/*
	 * 8-bit Ansi colors (240 different colors + 16 duplicate color codes, ensuring backward compatibility).
	 * Output syntax is: "ESC[38;5;${foreGroundColorcode};48;5;${backGroundColorcode}m"
	 * Should be compatible with most terminals.
	 */
	Ansi8 AnsiColorMode = "Ansi8"

	/*
	 * 24-bit Ansi colors (RGB).
	 * Output syntax is: "ESC[38;2;${foreGroundColorcodeRed};${foreGroundColorcodeGreen};${foreGroundColorcodeBlue};48;2;${backGroundColorcodeRed};${backGroundColorcodeGreen};${backGroundColorcodeBlue}m"
	 * May be compatible with many modern terminals.
	 */
	Ansi24 AnsiColorMode = "Ansi24"
)

func ConvertFromHexToAnsiColorCode(hexColor string, mode AnsiColorMode) string {
	hex := strings.TrimPrefix(hexColor, "#")

	if len(hex) == 3 {
		// #RGB
		// #RGB => #RRGGBB
		hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
	}

	if len(hex) != 6 {
		panic(fmt.Sprintf("Invalid color: #%s", hex))
	}

	color, err := strconv.ParseInt(hex, 16, 0)

	if err != nil {
		panic(fmt.Sprintf("Cannot parse color: #%s", hex))
	}

	r := (color >> 16) & 255
	g := (color >> 8) & 255
	b := color & 255

	switch mode {
	case Ansi4:
		return fmt.Sprintf("%d", convertFromRGB(r, g, b, mode))
	case Ansi8:
		return fmt.Sprintf("8;5;%d", convertFromRGB(r, g, b, mode))
	case Ansi24:
		return fmt.Sprintf("8;2;%d;%d;%d", r, g, b)
	default:
		panic("Invalid AnsiColorMode")
	}
}

func convertFromRGB(r int64, g int64, b int64, mode AnsiColorMode) int64 {
	switch mode {
	case Ansi4:
		return degradeHexColorToAnsi4(r, g, b)
	case Ansi8:
		return degradeHexColorToAnsi8(r, g, b)
	default:
		panic(fmt.Sprintf("RGB cannot be converted to %s", mode))
	}
}

func degradeHexColorToAnsi4(r int64, g int64, b int64) int64 {
	r = int64(math.Round(float64(r) / 255.0))
	g = int64(math.Round(float64(g) / 255.0))
	b = int64(math.Round(float64(b) / 255.0))

	return b<<2 | g<<1 | r
}

func degradeHexColorToAnsi8(r int64, g int64, b int64) int64 {
	if r == g && g == b {
		if r < 8 {
			return 16
		}

		if r > 248 {
			return 231
		}

		return int64(math.Round(((float64(r)-8.0)/247.0)*24.0)) + 232
	} else {
		return 16 +
			36*int64(math.Round((float64(r)/255.0)*5.0)) +
			6*int64(math.Round((float64(g)/255.0)*5.0)) +
			int64(math.Round((float64(b)/255.0)*5.0))
	}
}
