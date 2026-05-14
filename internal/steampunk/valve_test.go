package steampunk

import (
	"testing"
)

func TestNewValve(t *testing.T) {
	v := NewValve("main", 5.0)
	if v.name != "main" {
		t.Errorf("expected name 'main', got %q", v.name)
	}
	if v.maxFlow != 5.0 {
		t.Errorf("expected maxFlow 5.0, got %f", v.maxFlow)
	}
	if v.state != ValveClosed {
		t.Errorf("expected new valve to be closed")
	}
}

func TestNewValveDefaultMaxFlow(t *testing.T) {
	v := NewValve("test", 0)
	if v.maxFlow != 1.0 {
		t.Errorf("expected default maxFlow 1.0, got %f", v.maxFlow)
	}
}

func TestValveOpen(t *testing.T) {
	v := NewValve("test", 2.0)
	v.Open()
	if v.state != ValveOpen {
		t.Errorf("expected valve to be open")
	}
	if v.FlowRate() != 2.0 {
		t.Errorf("expected flow rate 2.0, got %f", v.FlowRate())
	}
}

func TestValveClose(t *testing.T) {
	v := NewValve("test", 2.0)
	v.Open()
	v.Close()
	if v.state != ValveClosed {
		t.Errorf("expected valve to be closed")
	}
	if v.FlowRate() != 0 {
		t.Errorf("expected flow rate 0, got %f", v.FlowRate())
	}
}

func TestValveSetPosition(t *testing.T) {
	v := NewValve("test", 4.0)
	v.SetPosition(0.5)
	if v.state != ValvePartial {
		t.Errorf("expected partial state")
	}
	if v.FlowRate() != 2.0 {
		t.Errorf("expected flow rate 2.0, got %f", v.FlowRate())
	}
}

func TestValveSetPositionClamped(t *testing.T) {
	v := NewValve("test", 1.0)
	v.SetPosition(1.5)
	if v.openPct != 1.0 {
		t.Errorf("expected openPct clamped to 1.0, got %f", v.openPct)
	}
	v.SetPosition(-0.5)
	if v.openPct != 0.0 {
		t.Errorf("expected openPct clamped to 0.0, got %f", v.openPct)
	}
}

func TestValveString(t *testing.T) {
	v := NewValve("main", 3.0)
	v.Open()
	s := v.String()
	if s == "" {
		t.Errorf("expected non-empty string")
	}
}
