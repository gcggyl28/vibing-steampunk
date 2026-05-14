package steampunk

import (
	"fmt"
	"math"
)

// GearSize represents the size category of a steampunk gear.
type GearSize int

const (
	GearSmall  GearSize = iota // 8 teeth
	GearMedium                 // 16 teeth
	GearLarge                  // 32 teeth
)

// Gear represents a steampunk mechanical gear component.
type Gear struct {
	ID       string
	Size     GearSize
	Teeth    int
	RPM      float64
	Material string
}

// NewGear creates a new Gear with the given parameters.
func NewGear(id string, size GearSize, rpm float64) *Gear {
	teeth := teethForSize(size)
	return &Gear{
		ID:       id,
		Size:     size,
		Teeth:    teeth,
		RPM:      rpm,
		Material: "brass",
	}
}

// teethForSize returns the number of teeth for a given gear size.
func teethForSize(size GearSize) int {
	switch size {
	case GearSmall:
		return 8
	case GearMedium:
		return 16
	case GearLarge:
		return 32
	default:
		return 16
	}
}

// MeshWith calculates the output RPM when this gear meshes with another.
func (g *Gear) MeshWith(other *Gear) float64 {
	if other.Teeth == 0 {
		return 0
	}
	return g.RPM * float64(g.Teeth) / float64(other.Teeth)
}

// AngularVelocity returns the angular velocity in radians per second.
func (g *Gear) AngularVelocity() float64 {
	return g.RPM * 2 * math.Pi / 60
}

// String returns a human-readable representation of the gear.
func (g *Gear) String() string {
	return fmt.Sprintf("Gear{id=%s, size=%d, teeth=%d, rpm=%.2f, material=%s}",
		g.ID, g.Size, g.Teeth, g.RPM, g.Material)
}
