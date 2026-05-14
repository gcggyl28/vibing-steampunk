package steampunk

import (
	"testing"
)

func TestNewCondenser(t *testing.T) {
	c, err := NewCondenser(10.0, 0.8)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("expected non-nil condenser")
	}
}

func TestNewCondenserInvalidArea(t *testing.T) {
	_, err := NewCondenser(-1.0, 0.8)
	if err == nil {
		t.Fatal("expected error for negative cooling area")
	}
}

func TestNewCondenserInvalidEfficiency(t *testing.T) {
	_, err := NewCondenser(10.0, 1.5)
	if err == nil {
		t.Fatal("expected error for efficiency > 1")
	}
	_, err = NewCondenser(10.0, -0.1)
	if err == nil {
		t.Fatal("expected error for negative efficiency")
	}
}

func TestCondenserReceiveSteam(t *testing.T) {
	c, _ := NewCondenser(10.0, 0.8)
	c.ReceiveSteam(5.0)
	if c.inletPressure != 5.0 {
		t.Errorf("expected inlet pressure 5.0, got %.2f", c.inletPressure)
	}
}

func TestCondenserOutletTemperature(t *testing.T) {
	c, _ := NewCondenser(2.0, 0.5)
	temp := c.OutletTemperature()
	// base=100, reduction=2*0.5*5=5, so temp=95
	if temp != 95.0 {
		t.Errorf("expected outlet temp 95.0, got %.2f", temp)
	}
}

func TestCondenserOutletTemperatureFloor(t *testing.T) {
	// Large area + high efficiency should floor at ambient (25°C)
	c, _ := NewCondenser(100.0, 1.0)
	temp := c.OutletTemperature()
	if temp < 25.0 {
		t.Errorf("outlet temp should not drop below ambient 25°C, got %.2f", temp)
	}
}

func TestCondenserCondensateFlow(t *testing.T) {
	c, _ := NewCondenser(10.0, 0.8)
	if c.CondensateFlow() != 0 {
		t.Error("expected zero flow with no steam")
	}
	c.ReceiveSteam(4.0)
	flow := c.CondensateFlow()
	if flow <= 0 {
		t.Errorf("expected positive condensate flow, got %.4f", flow)
	}
}

func TestCondenserIsEffective(t *testing.T) {
	c, _ := NewCondenser(10.0, 0.8)
	if c.IsEffective() {
		t.Error("should not be effective without steam")
	}
	c.ReceiveSteam(3.0)
	if !c.IsEffective() {
		t.Error("should be effective with steam and efficiency >= 0.5")
	}
}

func TestCondenserString(t *testing.T) {
	c, _ := NewCondenser(10.0, 0.8)
	s := c.String()
	if s == "" {
		t.Error("expected non-empty string representation")
	}
}
