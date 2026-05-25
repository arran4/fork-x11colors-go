package x11colors

import (
	"image/color"

	"testing"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		name     Name
		expected string
	}{
		{"Alice Blue", "alice-blue"},
		{"Purple (X11)", "purple-x11"},
	}

	for _, test := range tests {
		t.Run(string(test.name), func(t *testing.T) {
			if got := test.name.Slugify(); got != test.expected {
				t.Errorf("Slugify() = %q, want %q", got, test.expected)
			}
		})
	}
}

func TestGetByName(t *testing.T) {
	c, found := GetByName("Alice Blue")
	if !found {
		t.Errorf("Expected to find Alice Blue")
	}
	if c.Name != "Alice Blue" {
		t.Errorf("Expected Alice Blue, got %s", c.Name)
	}

	_, found = GetByName("Not A Color")
	if found {
		t.Errorf("Expected not to find Not A Color")
	}
}

func TestRandom(t *testing.T) {
	c := Random()
	if c.Name == "" {
		t.Errorf("Expected a color, got empty name")
	}
}

func TestRandomSeeded(t *testing.T) {
	c := RandomSeeded()
	if c.Name == "" {
		t.Errorf("Expected a color, got empty name")
	}
}

func TestGetClosest(t *testing.T) {
	tests := []struct {
		input    color.RGBA
		expected string
		name string
	}{
		{color.RGBA{R: 255, G: 255, B: 255, A: 255}, "gray100", "White"},
		{color.RGBA{R: 0, G: 0, B: 0, A: 255}, "Black", "Black"},
		{color.RGBA{R: 250, G: 0, B: 0, A: 255}, "Red", "Red"},
		{color.RGBA{R: 0, G: 250, B: 0, A: 255}, "green", "Green"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			closest := GetClosest(test.input)
			if closest.Name.String() != test.expected {
				t.Errorf("GetClosest() = %q, want %q", closest.Name, test.expected)
			}
		})
	}
}

func TestGetClosestEmpty(t *testing.T) {
	// Empty colors array
	oldColors := colors
	colors = []X11Color{}

	closest := GetClosest(color.RGBA{255, 255, 255, 255})
	if closest.Name != "" {
		t.Errorf("Expected empty name, got %s", closest.Name)
	}

	colors = oldColors
}
