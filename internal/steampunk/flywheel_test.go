package steampunk

import (
	"math"
	"testing"
)

func makeTestFlywheel(t *testing.T) *Flywheel {
	t.Helper()
	fw, err := NewFlywheel(0.5, 20.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return fw
}

func TestNewFlywheel(t *testing.T) {
	fw := makeTestFlywheel(t)
	if fw.radius != 0.5 {
		t.Errorf("expected radius 0.5, got %f", fw.radius)
	}
	if fw.mass != 20.0 {
		t.Errorf("expected mass 20.0, got %f", fw.mass)
	}
}

func TestNewFlywheelInvalidRadius(t *testing.T) {
	_, err := NewFlywheel(-1, 10)
	if err == nil {
		t.Error("expected error for negative radius")
	}
	_, err = NewFlywheel(0, 10)
	if err == nil {
		t.Error("expected error for zero radius")
	}
}

func TestNewFlywheelInvalidMass(t *testing.T) {
	_, err := NewFlywheel(0.5, -5)
	if err == nil {
		t.Error("expected error for negative mass")
	}
}

func TestFlywheelMomentOfInertia(t *testing.T) {
	fw := makeTestFlywheel(t) // r=0.5, m=20
	expected := 0.5 * 20.0 * 0.5 * 0.5 // 2.5
	if math.Abs(fw.MomentOfInertia()-expected) > 1e-9 {
		t.Errorf("expected I=%.4f, got %.4f", expected, fw.MomentOfInertia())
	}
}

func TestFlywheelKineticEnergyAtRest(t *testing.T) {
	fw := makeTestFlywheel(t)
	if fw.KineticEnergy() != 0 {
		t.Errorf("expected 0 KE at rest, got %f", fw.KineticEnergy())
	}
}

func TestFlywheelSetRPM(t *testing.T) {
	fw := makeTestFlywheel(t)
	fw.SetRPM(300)
	if fw.RPM() != 300 {
		t.Errorf("expected rpm=300, got %f", fw.RPM())
	}
	fw.SetRPM(-10)
	if fw.RPM() != 0 {
		t.Errorf("expected rpm clamped to 0, got %f", fw.RPM())
	}
}

func TestFlywheelApplyTorque(t *testing.T) {
	fw := makeTestFlywheel(t)
	// I = 2.5 kg·m², torque = 25 N·m, dt = 1s → α = 10 rad/s²
	// ω after 1s = 10 rad/s → rpm = 10*60/(2π) ≈ 95.49
	newRPM := fw.ApplyTorque(25, 1.0)
	expectedRPM := 10.0 * 60 / (2 * math.Pi)
	if math.Abs(newRPM-expectedRPM) > 0.01 {
		t.Errorf("expected rpm≈%.2f, got %.2f", expectedRPM, newRPM)
	}
}

func TestFlywheelApplyTorqueZeroDt(t *testing.T) {
	fw := makeTestFlywheel(t)
	fw.SetRPM(100)
	newRPM := fw.ApplyTorque(1000, 0)
	if newRPM != 100 {
		t.Errorf("expected rpm unchanged at 100, got %f", newRPM)
	}
}

func TestFlywheelString(t *testing.T) {
	fw := makeTestFlywheel(t)
	s := fw.String()
	if len(s) == 0 {
		t.Error("expected non-empty string")
	}
}
