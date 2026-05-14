package steampunk

import (
	"testing"
)

func makeTestGovernorValve(t *testing.T) *Valve {
	t.Helper()
	v, err := NewValve("governor-valve", 100.0)
	if err != nil {
		t.Fatalf("failed to create valve: %v", err)
	}
	_ = v.Open()
	return v
}

func TestNewGovernor(t *testing.T) {
	v := makeTestGovernorValve(t)
	g, err := NewGovernor(1200.0, 0.5, v)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if g.TargetRPM() != 1200.0 {
		t.Errorf("expected target RPM 1200, got %.1f", g.TargetRPM())
	}
}

func TestNewGovernorInvalidRPM(t *testing.T) {
	v := makeTestGovernorValve(t)
	_, err := NewGovernor(-100.0, 0.5, v)
	if err == nil {
		t.Error("expected error for negative target RPM")
	}
}

func TestNewGovernorInvalidSensitivity(t *testing.T) {
	v := makeTestGovernorValve(t)
	_, err := NewGovernor(1200.0, 1.5, v)
	if err == nil {
		t.Error("expected error for sensitivity > 1")
	}
	_, err = NewGovernor(1200.0, 0.0, v)
	if err == nil {
		t.Error("expected error for sensitivity == 0")
	}
}

func TestNewGovernorNilValve(t *testing.T) {
	_, err := NewGovernor(1200.0, 0.5, nil)
	if err == nil {
		t.Error("expected error for nil valve")
	}
}

func TestGovernorActivate(t *testing.T) {
	v := makeTestGovernorValve(t)
	g, _ := NewGovernor(1200.0, 0.5, v)
	if g.IsActive() {
		t.Error("governor should be inactive by default")
	}
	g.Activate()
	if !g.IsActive() {
		t.Error("governor should be active after Activate()")
	}
	g.Deactivate()
	if g.IsActive() {
		t.Error("governor should be inactive after Deactivate()")
	}
}

func TestGovernorUpdateInactive(t *testing.T) {
	v := makeTestGovernorValve(t)
	_ = v.SetPosition(0.7)
	g, _ := NewGovernor(1200.0, 0.5, v)

	pos := g.Update(900.0)
	if pos != 0.7 {
		t.Errorf("inactive governor should not change valve position, got %.2f", pos)
	}
}

func TestGovernorUpdateReducesFlowWhenFast(t *testing.T) {
	v := makeTestGovernorValve(t)
	_ = v.SetPosition(0.8)
	g, _ := NewGovernor(1000.0, 0.5, v)
	g.Activate()

	// Engine running too fast — governor should close valve slightly
	pos := g.Update(1200.0)
	if pos >= 0.8 {
		t.Errorf("expected valve to close when RPM exceeds target, got position %.2f", pos)
	}
}

func TestGovernorUpdateIncreasesFlowWhenSlow(t *testing.T) {
	v := makeTestGovernorValve(t)
	_ = v.SetPosition(0.5)
	g, _ := NewGovernor(1000.0, 0.5, v)
	g.Activate()

	// Engine running too slow — governor should open valve slightly
	pos := g.Update(800.0)
	if pos <= 0.5 {
		t.Errorf("expected valve to open when RPM is below target, got position %.2f", pos)
	}
}

func TestGovernorString(t *testing.T) {
	v := makeTestGovernorValve(t)
	g, _ := NewGovernor(1200.0, 0.5, v)
	g.Activate()
	s := g.String()
	if s == "" {
		t.Error("String() should return a non-empty description")
	}
}
