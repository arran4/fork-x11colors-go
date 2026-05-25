package x11colors

import (
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
