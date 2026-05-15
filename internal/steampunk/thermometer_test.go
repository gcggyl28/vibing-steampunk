package steampunk

import (
	"strings"
	"testing"
)

func makeTestThermometer(t *testing.T) *Thermometer {
	t.Helper()
	th, err := NewThermometer("boiler-temp", 20.0, 200.0, 0.6, 0.85)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return th
}

func TestNewThermometer(t *testing.T) {
	th := makeTestThermometer(t)
	if th.Temperature() != 20.0 {
		t.Errorf("expected initial temp 20.0, got %.2f", th.Temperature())
	}
}

func TestNewThermometerInvalidRange(t *testing.T) {
	_, err := NewThermometer("bad", 100.0, 50.0, 0.6, 0.85)
	if err == nil {
		t.Error("expected error for minTemp >= maxTemp")
	}
}

func TestNewThermometerInvalidWarnRatio(t *testing.T) {
	_, err := NewThermometer("bad", 20.0, 200.0, 1.1, 0.85)
	if err == nil {
		t.Error("expected error for warnRatio out of range")
	}
}

func TestNewThermometerInvalidCritRatio(t *testing.T) {
	_, err := NewThermometer("bad", 20.0, 200.0, 0.6, 0.5)
	if err == nil {
		t.Error("expected error for critRatio <= warnRatio")
	}
}

func TestThermometerSetTemperature(t *testing.T) {
	th := makeTestThermometer(t)
	th.SetTemperature(100.0)
	if th.Temperature() != 100.0 {
		t.Errorf("expected 100.0, got %.2f", th.Temperature())
	}
}

func TestThermometerClampMin(t *testing.T) {
	th := makeTestThermometer(t)
	th.SetTemperature(-50.0)
	if th.Temperature() != 20.0 {
		t.Errorf("expected clamped to 20.0, got %.2f", th.Temperature())
	}
}

func TestThermometerClampMax(t *testing.T) {
	th := makeTestThermometer(t)
	th.SetTemperature(999.0)
	if th.Temperature() != 200.0 {
		t.Errorf("expected clamped to 200.0, got %.2f", th.Temperature())
	}
}

func TestThermometerWarning(t *testing.T) {
	th := makeTestThermometer(t)
	// warnTemp = 20 + (200-20)*0.6 = 128.0
	th.SetTemperature(130.0)
	if !th.IsWarning() {
		t.Error("expected warning at 130.0")
	}
	if th.IsCritical() {
		t.Error("expected not critical at 130.0")
	}
}

func TestThermometerCritical(t *testing.T) {
	th := makeTestThermometer(t)
	// critTemp = 20 + (200-20)*0.85 = 173.0
	th.SetTemperature(180.0)
	if !th.IsCritical() {
		t.Error("expected critical at 180.0")
	}
}

func TestThermometerStatusOK(t *testing.T) {
	th := makeTestThermometer(t)
	th.SetTemperature(50.0)
	if !strings.Contains(th.Status(), "OK") {
		t.Errorf("expected OK status, got: %s", th.Status())
	}
}

func TestThermometerString(t *testing.T) {
	th := makeTestThermometer(t)
	th.SetTemperature(50.0)
	if th.String() != th.Status() {
		t.Error("String() should match Status()")
	}
}
