package steampunk

import (
	"math"
	"testing"
)

func makeTestCrankshaft(t *testing.T) (*Crankshaft, *Piston) {
	t.Helper()
	c, err := NewCrankshaft("CS1", 0.05)
	if err != nil {
		t.Fatalf("NewCrankshaft: %v", err)
	}
	p, err := NewPiston("P1", 10.0, 15.0)
	if err != nil {
		t.Fatalf("NewPiston: %v", err)
	}
	c.AttachPiston(p)
	return c, p
}

func TestNewCrankshaft(t *testing.T) {
	c, _ := makeTestCrankshaft(t)
	if c.CrankRadius != 0.05 {
		t.Errorf("expected radius 0.05, got %.4f", c.CrankRadius)
	}
	if len(c.Pistons) != 1 {
		t.Errorf("expected 1 piston, got %d", len(c.Pistons))
	}
}

func TestNewCrankshaftInvalidRadius(t *testing.T) {
	_, err := NewCrankshaft("CS1", 0)
	if err == nil {
		t.Error("expected error for zero radius")
	}
}

func TestCrankshaftTorqueZeroWithoutPressure(t *testing.T) {
	c, _ := makeTestCrankshaft(t)
	c.Angle = math.Pi / 4
	if c.Torque() != 0 {
		t.Errorf("torque should be zero without applied pressure, got %.4f", c.Torque())
	}
}

func TestCrankshaftTorqueWithPressure(t *testing.T) {
	c, p := makeTestCrankshaft(t)
	p.ApplyPressure(8.0)
	c.Angle = math.Pi / 2 // maximum sine contribution
	torque := c.Torque()
	if torque <= 0 {
		t.Errorf("expected positive torque, got %.4f", torque)
	}
}

func TestCrankshaftRotate(t *testing.T) {
	c, p := makeTestCrankshaft(t)
	p.ApplyPressure(5.0)
	initialAngle := c.Angle
	c.Rotate(math.Pi / 6)
	if c.Angle == initialAngle {
		t.Error("angle should change after Rotate")
	}
	if c.RPM <= 0 {
		t.Error("RPM should be positive after rotation")
	}
}

func TestCrankshaftRotateWraps(t *testing.T) {
	c, _ := makeTestCrankshaft(t)
	c.Rotate(3 * math.Pi) // more than one full revolution
	if c.Angle < 0 || c.Angle >= 2*math.Pi {
		t.Errorf("angle out of range: %.4f", c.Angle)
	}
}

func TestCrankshaftString(t *testing.T) {
	c, _ := makeTestCrankshaft(t)
	if s := c.String(); len(s) == 0 {
		t.Error("String() should not be empty")
	}
}
