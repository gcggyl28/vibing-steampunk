package steampunk

import (
	"testing"
)

func makeTestSteamTrap(t *testing.T) *SteamTrap {
	t.Helper()
	trap, err := NewSteamTrap(0.05, 120.0, 0.85)
	if err != nil {
		t.Fatalf("failed to create steam trap: %v", err)
	}
	return trap
}

func TestNewSteamTrap(t *testing.T) {
	trap, err := NewSteamTrap(0.05, 120.0, 0.85)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if trap == nil {
		t.Fatal("expected non-nil steam trap")
	}
}

func TestNewSteamTrapInvalidOrificeSize(t *testing.T) {
	_, err := NewSteamTrap(0.0, 120.0, 0.85)
	if err == nil {
		t.Fatal("expected error for zero orifice size")
	}

	_, err = NewSteamTrap(-0.01, 120.0, 0.85)
	if err == nil {
		t.Fatal("expected error for negative orifice size")
	}
}

func TestNewSteamTrapInvalidSaturationTemp(t *testing.T) {
	_, err := NewSteamTrap(0.05, 0.0, 0.85)
	if err == nil {
		t.Fatal("expected error for zero saturation temperature")
	}

	_, err = NewSteamTrap(0.05, -10.0, 0.85)
	if err == nil {
		t.Fatal("expected error for negative saturation temperature")
	}
}

func TestNewSteamTrapInvalidEfficiency(t *testing.T) {
	_, err := NewSteamTrap(0.05, 120.0, 0.0)
	if err == nil {
		t.Fatal("expected error for zero efficiency")
	}

	_, err = NewSteamTrap(0.05, 120.0, 1.1)
	if err == nil {
		t.Fatal("expected error for efficiency > 1")
	}
}

func TestSteamTrapInitiallyClosed(t *testing.T) {
	trap := makeTestSteamTrap(t)
	if trap.IsOpen() {
		t.Error("expected steam trap to be initially closed")
	}
}

func TestSteamTrapOpenOnCondensate(t *testing.T) {
	trap := makeTestSteamTrap(t)
	// Temperature below saturation triggers condensate discharge
	trap.Update(80.0, 3.0) // 80°C < 120°C saturation
	if !trap.IsOpen() {
		t.Error("expected steam trap to open when condensate present")
	}
}

func TestSteamTrapClosedOnSteam(t *testing.T) {
	trap := makeTestSteamTrap(t)
	// Temperature at or above saturation — steam present, trap stays closed
	trap.Update(125.0, 3.0)
	if trap.IsOpen() {
		t.Error("expected steam trap to remain closed when steam is present")
	}
}

func TestSteamTrapDischargeRate(t *testing.T) {
	trap := makeTestSteamTrap(t)
	trap.Update(80.0, 3.0)

	rate := trap.DischargeRate()
	if rate <= 0 {
		t.Errorf("expected positive discharge rate when open, got %f", rate)
	}
}

func TestSteamTrapDischargeRateZeroWhenClosed(t *testing.T) {
	trap := makeTestSteamTrap(t)
	trap.Update(130.0, 3.0) // above saturation, trap closed

	rate := trap.DischargeRate()
	if rate != 0 {
		t.Errorf("expected zero discharge rate when closed, got %f", rate)
	}
}

func TestSteamTrapCondensateRemoved(t *testing.T) {
	trap := makeTestSteamTrap(t)
	trap.Update(80.0, 3.0)

	removed := trap.CondensateRemoved()
	if removed <= 0 {
		t.Errorf("expected positive condensate removed, got %f", removed)
	}
}

func TestSteamTrapString(t *testing.T) {
	trap := makeTestSteamTrap(t)
	s := trap.String()
	if s == "" {
		t.Error("expected non-empty string representation")
	}
}
