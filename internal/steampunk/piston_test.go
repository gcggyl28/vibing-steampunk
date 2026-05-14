package steampunk

import (
	"math"
	"testing"
)

func TestNewPiston(t *testing.T) {
	p, err := NewPiston("P1", 10.0, 15.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if p.Bore != 10.0 || p.Stroke != 15.0 {
		t.Errorf("unexpected dimensions bore=%.1f stroke=%.1f", p.Bore, p.Stroke)
	}
	if p.State != PistonIdle {
		t.Errorf("new piston should be idle")
	}
}

func TestNewPistonInvalidBore(t *testing.T) {
	_, err := NewPiston("P1", 0, 15.0)
	if err == nil {
		t.Error("expected error for zero bore")
	}
}

func TestNewPistonInvalidStroke(t *testing.T) {
	_, err := NewPiston("P1", 10.0, -1)
	if err == nil {
		t.Error("expected error for negative stroke")
	}
}

func TestPistonArea(t *testing.T) {
	p, _ := NewPiston("P1", 10.0, 15.0)
	expected := math.Pi * 25.0 // π * r²
	got := p.Area()
	if math.Abs(got-expected) > 0.01 {
		t.Errorf("area: expected %.4f got %.4f", expected, got)
	}
}

func TestPistonApplyPressure(t *testing.T) {
	p, _ := NewPiston("P1", 10.0, 15.0)
	p.ApplyPressure(5.0)
	if p.Force <= 0 {
		t.Error("force should be positive after applying pressure")
	}
	if p.State != PistonStroke {
		t.Error("piston should transition to stroke state")
	}
}

func TestPistonApplyZeroPressure(t *testing.T) {
	p, _ := NewPiston("P1", 10.0, 15.0)
	p.ApplyPressure(0)
	if p.State != PistonIdle {
		t.Error("piston should remain idle with zero pressure")
	}
}

func TestPistonAdvanceCycle(t *testing.T) {
	p, _ := NewPiston("P1", 10.0, 15.0)
	p.ApplyPressure(5.0)

	cycles := 0
	for i := 0; i < 1000; i++ {
		if p.Advance(0.1) {
			cycles++
			break
		}
	}
	if cycles != 1 {
		t.Errorf("expected 1 cycle completion, got %d", cycles)
	}
}

func TestPistonString(t *testing.T) {
	p, _ := NewPiston("P1", 10.0, 15.0)
	s := p.String()
	if len(s) == 0 {
		t.Error("String() should not be empty")
	}
}
