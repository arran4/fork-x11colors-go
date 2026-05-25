package x11colors

import (
	"image/color"
	"math"
	"math/rand"
	"strings"
)

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
