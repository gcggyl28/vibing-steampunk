package steampunk

import (
	"math"
	"testing"
)

func TestNewBoiler(t *testing.T) {
	b := NewBoiler(150.0, 20.0)
	if b.MaxPressure != 150.0 {
		t.Errorf("expected MaxPressure 150.0, got %.1f", b.MaxPressure)
	}
	if b.Volume != 20.0 {
		t.Errorf("expected Volume 20.0, got %.1f", b.Volume)
	}
	if b.CurrentPressure != 0 {
		t.Errorf("expected initial pressure 0, got %.1f", b.CurrentPressure)
	}
}

func TestNewBoilerDefaults(t *testing.T) {
	b := NewBoiler(-1, 0)
	if b.MaxPressure != 100.0 {
		t.Errorf("expected default MaxPressure 100.0, got %.1f", b.MaxPressure)
	}
	if b.Volume != 10.0 {
		t.Errorf("expected default Volume 10.0, got %.1f", b.Volume)
	}
}

func TestBoilerHeat(t *testing.T) {
	b := NewBoiler(200.0, 10.0)
	b.Heat(120.0) // brings temp to 140°C
	if b.CurrentPressure <= 0 {
		t.Errorf("expected positive pressure after heating above 100°C, got %.1f", b.CurrentPressure)
	}
}

func TestBoilerMaxPressure(t *testing.T) {
	b := NewBoiler(50.0, 10.0)
	b.Heat(500.0) // extreme heat
	if b.CurrentPressure > b.MaxPressure {
		t.Errorf("pressure %.1f exceeded max %.1f", b.CurrentPressure, b.MaxPressure)
	}
}

func TestBoilerIsOperational(t *testing.T) {
	b := NewBoiler(150.0, 10.0)
	if b.IsOperational() {
		t.Error("cold boiler should not be operational")
	}
	b.Heat(130.0)
	if !b.IsOperational() {
		t.Error("heated boiler should be operational")
	}
}

func TestBoilerVent(t *testing.T) {
	b := NewBoiler(150.0, 10.0)
	b.Heat(130.0)
	pressureBefore := b.CurrentPressure
	b.Vent(5.0)
	if b.CurrentPressure >= pressureBefore {
		t.Error("pressure should decrease after venting")
	}
}

func TestBoilerTorque(t *testing.T) {
	b := NewBoiler(150.0, 16.0)
	b.Heat(130.0)
	torque := b.Torque()
	expected := b.CurrentPressure * math.Sqrt(b.Volume) * b.Efficiency
	if math.Abs(torque-expected) > 0.001 {
		t.Errorf("expected torque %.4f, got %.4f", expected, torque)
	}
}

func TestBoilerString(t *testing.T) {
	b := NewBoiler(100.0, 10.0)
	s := b.String()
	if len(s) == 0 {
		t.Error("expected non-empty string representation")
	}
}
