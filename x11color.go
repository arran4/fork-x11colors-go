package x11colors

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// ToString returns a standard hex string representation of the color (e.g., "#F0F8FF").
func (c X11Color) ToString() string {
	return fmt.Sprintf("#%02X%02X%02X", c.RGBA.R, c.RGBA.G, c.RGBA.B)
}

// String returns string value for color name
func (n Name) String() string {
	return string(n)
}

// Slugify returns url slugs string value for color name
func (n Name) Slugify() string {
	return strings.ToLower(
		strings.Replace(
			strings.Replace(
				strings.Replace(n.String(), ")", "", -1), "(", "", -1,
			),
			" ", "-", -1,
		),
	)
}

// Random returns random color
func Random() X11Color {
	return colors[rand.Intn(len(colors))]
}

// RandomSeeded initialises generator with time-based seed and returns random color.
// Deprecated: In modern Go, the global random generator is automatically seeded.
// This function simply calls Random().
func RandomSeeded() X11Color {
	return Random()
}

// GetByName returns X11Color by its name if found
func GetByName(name string) (x11color X11Color, found bool) {
	x11color, found = names[name]
	return
}

// FromString attempts to parse a string into an X11Color.
// It tries to match by exact name, case-insensitive name, slug, or hex code (e.g., "#F0F8FF" or "#FFF").
func FromString(s string) (X11Color, bool) {
	// Try exact match
	if c, ok := names[s]; ok {
		return c, true
	}

	// Try parsing as hex code
	if strings.HasPrefix(s, "#") {
		hex := s[1:]
		if len(hex) == 3 {
			// Expand short hex e.g. #ABC to #AABBCC
			hex = string([]byte{hex[0], hex[0], hex[1], hex[1], hex[2], hex[2]})
		}
		if len(hex) == 6 {
			val, err := strconv.ParseUint(hex, 16, 32)
			if err == nil {
				r := uint8(val >> 16)
				g := uint8((val >> 8) & 0xFF)
				b := uint8(val & 0xFF)

				// Check if this hex matches any X11 color perfectly
				for _, c := range colors {
					if c.RGBA.R == r && c.RGBA.G == g && c.RGBA.B == b {
						return c, true
					}
				}

				// Basic lightness calculation to set a sensible CaptionBlack value
				// Luma = 0.2126 * R + 0.7152 * G + 0.0722 * B
				luma := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)
				captionBlack := luma > 128

				return X11Color{
					Name:         Name(s),
					RGBA:         color.RGBA{R: r, G: g, B: b, A: 0xFF},
					CaptionBlack: captionBlack,
				}, true
			}
		}
	}

	// Try case-insensitive name match or slug match
	sLower := strings.ToLower(s)
	// We can iterate through names to find case-insensitive match since it's already generated
	for nameStr, c := range names {
		if strings.ToLower(nameStr) == sLower || c.Name.Slugify() == sLower {
			return c, true
		}
	}

	return X11Color{}, false
}

func sqDiff(c1, c2 uint8) float64 {
	d := float64(c1) - float64(c2)
	return d * d
}

// GetClosest returns the X11Color that is closest to the provided color.RGBA
// It uses simple Euclidean distance in RGB color space.
func GetClosest(c color.RGBA) X11Color {
	if len(colors) == 0 {
		return X11Color{}
	}
	closest := colors[0]
	minDist := math.MaxFloat64

	for _, x11c := range colors {
		dist := sqDiff(c.R, x11c.RGBA.R) + sqDiff(c.G, x11c.RGBA.G) + sqDiff(c.B, x11c.RGBA.B)
		if dist < minDist {
			minDist = dist
			closest = x11c
		}
	}
	return closest
}
