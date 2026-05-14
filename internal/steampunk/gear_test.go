package steampunk_test

import (
	"math"
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestNewGear(t *testing.T) {
	tests := []struct {
		name      string
		size      steampunk.GearSize
		expTeeth  int
	}{
		{"small gear", steampunk.GearSmall, 8},
		{"medium gear", steampunk.GearMedium, 16},
		{"large gear", steampunk.GearLarge, 32},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := steampunk.NewGear("g1", tt.size, 60.0)
			if g.Teeth != tt.expTeeth {
				t.Errorf("expected %d teeth, got %d", tt.expTeeth, g.Teeth)
			}
			if g.Material != "brass" {
				t.Errorf("expected material brass, got %s", g.Material)
			}
		})
	}
}

func TestGearMeshWith(t *testing.T) {
	driver := steampunk.NewGear("driver", steampunk.GearLarge, 60.0) // 32 teeth, 60 RPM
	driven := steampunk.NewGear("driven", steampunk.GearSmall, 0.0)  // 8 teeth

	outRPM := driver.MeshWith(driven)
	expected := 240.0 // 60 * 32 / 8

	if math.Abs(outRPM-expected) > 1e-9 {
		t.Errorf("expected output RPM %.2f, got %.2f", expected, outRPM)
	}
}

func TestGearMeshWithZeroTeeth(t *testing.T) {
	driver := steampunk.NewGear("driver", steampunk.GearMedium, 100.0)
	broken := &steampunk.Gear{ID: "broken", Teeth: 0}

	outRPM := driver.MeshWith(broken)
	if outRPM != 0 {
		t.Errorf("expected 0 RPM for zero-teeth gear, got %.2f", outRPM)
	}
}

func TestGearAngularVelocity(t *testing.T) {
	g := steampunk.NewGear("g", steampunk.GearMedium, 60.0)
	expected := 2 * math.Pi // 60 RPM => 2π rad/s

	av := g.AngularVelocity()
	if math.Abs(av-expected) > 1e-9 {
		t.Errorf("expected angular velocity %.6f, got %.6f", expected, av)
	}
}

func TestGearString(t *testing.T) {
	g := steampunk.NewGear("cog1", steampunk.GearSmall, 30.0)
	s := g.String()
	if s == "" {
		t.Error("expected non-empty string representation")
	}
}
