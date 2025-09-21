package theme

import (
	"strconv"
	"strings"

	"github.com/jupiterrider/purego-sdl3/sdl"
)

type ThemeColors struct {
	Background    sdl.Color // general background
	Surface       sdl.Color // card/panel background
	Primary       sdl.Color // main accent color
	Secondary     sdl.Color // secondary accent color
	Success       sdl.Color
	Warning       sdl.Color
	Error         sdl.Color
	TextPrimary   sdl.Color // main text
	TextSecondary sdl.Color // secondary text
	Border        sdl.Color
}

type ThemeSpacing struct {
	XS int
	S  int
	M  int
	L  int
	XL int
}

type ThemeTypography struct {
	FontFamily string
	FontSizeS  float32
	FontSizeM  float32
	FontSizeL  float32
}

type Theme struct {
	Colors     ThemeColors
	Spacing    ThemeSpacing
	Typography ThemeTypography
}

var DefaultTheme = Theme{
	Colors: ThemeColors{
		Background:    HexToColor("#181818"),
		Surface:       HexToColor("#505050"),
		Primary:       HexToColor("#1976D2"),
		Secondary:     HexToColor("#9C27B0"),
		Success:       HexToColor("#4CAF50"),
		Warning:       HexToColor("#FFC107"),
		Error:         HexToColor("#F44336"),
		TextPrimary:   HexToColor("#ffffff"),
		TextSecondary: HexToColor("#757575"),
		Border:        HexToColor("#E0E0E0"),
	},
	Spacing: ThemeSpacing{
		XS: 4,
		S:  8,
		M:  16,
		L:  24,
		XL: 32,
	},
	Typography: ThemeTypography{
		FontFamily: "Roboto-Light.ttf",
		FontSizeS:  20,
		FontSizeM:  40,
		FontSizeL:  60,
	},
}

// HexToColor converts "#RRGGBB" or "#RRGGBBAA" into sdl.Color.
// Alpha defaults to 255 if omitted.
func HexToColor(hex string) sdl.Color {
	hex = strings.TrimPrefix(hex, "#")

	parse := func(s string) uint8 {
		i, _ := strconv.ParseUint(s, 16, 8)
		return uint8(i)
	}

	switch len(hex) {
	case 6:
		return sdl.Color{
			R: parse(hex[0:2]),
			G: parse(hex[2:4]),
			B: parse(hex[4:6]),
			A: 255,
		}
	case 8:
		return sdl.Color{
			R: parse(hex[0:2]),
			G: parse(hex[2:4]),
			B: parse(hex[4:6]),
			A: parse(hex[6:8]),
		}
	default:
		return sdl.Color{R: 0, G: 0, B: 0, A: 255} // fallback if length is wrong
	}
}
