package steampunk

import (
	"testing"
)

func makeTestGauge(t *testing.T) *PressureGauge {
	t.Helper()
	g, err := NewPressureGauge("main", 0, 200, 0.8, 0.95)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return g
}

func TestNewPressureGauge(t *testing.T) {
	g := makeTestGauge(t)
	if g.Reading() != 0 {
		t.Errorf("expected initial reading 0, got %.2f", g.Reading())
	}
}

func TestNewPressureGaugeInvalidRange(t *testing.T) {
	_, err := NewPressureGauge("bad", 100, 50, 0.8, 0.95)
	if err == nil {
		t.Error("expected error for invalid pressure range")
	}
}

func TestNewPressureGaugeInvalidWarnRatio(t *testing.T) {
	_, err := NewPressureGauge("bad", 0, 200, 1.1, 0.95)
	if err == nil {
		t.Error("expected error for warnRatio >= 1")
	}
}

func TestNewPressureGaugeInvalidCritRatio(t *testing.T) {
	_, err := NewPressureGauge("bad", 0, 200, 0.8, 0.75)
	if err == nil {
		t.Error("expected error for critRatio <= warnRatio")
	}
}

func TestPressureGaugeNormalStatus(t *testing.T) {
	g := makeTestGauge(t)
	g.Update(100)
	if g.Status() != "NORMAL" {
		t.Errorf("expected NORMAL, got %s", g.Status())
	}
}

func TestPressureGaugeWarning(t *testing.T) {
	g := makeTestGauge(t)
	g.Update(165) // 200 * 0.8 = 160, so 165 > warn
	if !g.IsWarning() {
		t.Error("expected warning state")
	}
	if g.IsCritical() {
		t.Error("should not be critical yet")
	}
}

func TestPressureGaugeCritical(t *testing.T) {
	g := makeTestGauge(t)
	g.Update(192) // 200 * 0.95 = 190, so 192 > crit
	if !g.IsCritical() {
		t.Error("expected critical state")
	}
	if g.Status() != "CRITICAL" {
		t.Errorf("expected CRITICAL, got %s", g.Status())
	}
}

func TestPressureGaugeOverpressure(t *testing.T) {
	g := makeTestGauge(t)
	g.Update(210)
	if !g.IsOverpressure() {
		t.Error("expected overpressure state")
	}
	if g.Status() != "OVERPRESSURE" {
		t.Errorf("expected OVERPRESSURE, got %s", g.Status())
	}
}

func TestPressureGaugeString(t *testing.T) {
	g := makeTestGauge(t)
	g.Update(50)
	s := g.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}
