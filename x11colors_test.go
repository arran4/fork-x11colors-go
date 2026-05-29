package x11colors

import (
	"image/color"
	"testing"
)

func TestToString(t *testing.T) {
	c := X11Color{
		Name: "Alice Blue",
		RGBA: color.RGBA{0xF0, 0xF8, 0xFF, 0xFF},
	}
	if str := c.ToString(); str != "#F0F8FF" {
		t.Errorf("expected #F0F8FF, got %s", str)
	}

	c2 := X11Color{
		Name: "Black",
		RGBA: color.RGBA{0x00, 0x00, 0x00, 0xFF},
	}
	if str := c2.ToString(); str != "#000000" {
		t.Errorf("expected #000000, got %s", str)
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
		expectName  string
	}{
		// Exact Match
		{"Alice Blue", false, "Alice Blue"},
		{"Black", false, "Black"},

		// Case-insensitive match
		{"alice blue", false, "Alice Blue"},
		{"BLACK", false, "Black"},
		{"WhItE", false, "White"},

		// Slug match
		{"alice-blue", false, "Alice Blue"},

		// Hex match perfect matches X11 color
		{"#F0F8FF", false, "Alice Blue"},
		{"#f0f8ff", false, "Alice Blue"}, // lowercase hex
		{"#000", false, "Black"}, // short hex

		// Hex match arbitrary
		{"#123456", false, "#123456"},
		{"#abc", false, "#abc"},

		// Errors
		{"invalid", true, ""},
		{"#1234", true, ""}, // bad hex length
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			c, err := FromString(tt.input)
			if tt.expectError {
				if err == nil {
					t.Errorf("expected error for input %q, got none", tt.input)
				}
				return
			}
			if err != nil {
				t.Errorf("did not expect error for input %q, got: %v", tt.input, err)
				return
			}
			if c.Name.String() != tt.expectName {
				t.Errorf("expected name %q, got %q", tt.expectName, c.Name)
			}
		})
	}
}
