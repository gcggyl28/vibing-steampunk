package steampunk

import (
	"testing"
)

func makeTestPressureMonitor(t *testing.T) (*PressureMonitor, *Boiler) {
	t.Helper()
	boiler, err := NewBoiler(5.0, 200.0)
	if err != nil {
		t.Fatalf("failed to create boiler: %v", err)
	}
	monitor, err := NewPressureMonitor(boiler, 4.5)
	if err != nil {
		t.Fatalf("failed to create pressure monitor: %v", err)
	}
	return monitor, boiler
}

func TestNewPressureMonitor(t *testing.T) {
	monitor, _ := makeTestPressureMonitor(t)
	if monitor == nil {
		t.Fatal("expected non-nil PressureMonitor")
	}
}

func TestNewPressureMonitorNilBoiler(t *testing.T) {
	_, err := NewPressureMonitor(nil, 4.5)
	if err == nil {
		t.Fatal("expected error for nil boiler, got nil")
	}
}

func TestNewPressureMonitorInvalidThreshold(t *testing.T) {
	boiler, err := NewBoiler(5.0, 200.0)
	if err != nil {
		t.Fatalf("failed to create boiler: %v", err)
	}
	_, err = NewPressureMonitor(boiler, 0.0)
	if err == nil {
		t.Fatal("expected error for zero threshold, got nil")
	}
	_, err = NewPressureMonitor(boiler, -1.0)
	if err == nil {
		t.Fatal("expected error for negative threshold, got nil")
	}
}

func TestPressureMonitorBelowThreshold(t *testing.T) {
	monitor, boiler := makeTestPressureMonitor(t)
	// Boiler starts cold with zero pressure
	if monitor.IsOverPressure() {
		t.Errorf("expected IsOverPressure=false when pressure=%.2f, threshold=%.2f",
			boiler.Pressure(), monitor.Threshold())
	}
}

func TestPressureMonitorAtThreshold(t *testing.T) {
	boiler, err := NewBoiler(5.0, 200.0)
	if err != nil {
		t.Fatalf("failed to create boiler: %v", err)
	}
	// Heat boiler to max pressure
	for i := 0; i < 1000; i++ {
		boiler.Heat(10.0)
	}
	monitor, err := NewPressureMonitor(boiler, 4.5)
	if err != nil {
		t.Fatalf("failed to create pressure monitor: %v", err)
	}
	if !monitor.IsOverPressure() {
		t.Errorf("expected IsOverPressure=true when pressure=%.2f exceeds threshold=%.2f",
			boiler.Pressure(), monitor.Threshold())
	}
}

func TestPressureMonitorThreshold(t *testing.T) {
	monitor, _ := makeTestPressureMonitor(t)
	if monitor.Threshold() != 4.5 {
		t.Errorf("expected threshold=4.5, got %.2f", monitor.Threshold())
	}
}

func TestPressureMonitorCurrentPressure(t *testing.T) {
	monitor, boiler := makeTestPressureMonitor(t)
	boiler.Heat(50.0)
	if monitor.CurrentPressure() != boiler.Pressure() {
		t.Errorf("expected monitor pressure=%.4f to match boiler pressure=%.4f",
			monitor.CurrentPressure(), boiler.Pressure())
	}
}

func TestPressureMonitorString(t *testing.T) {
	monitor, _ := makeTestPressureMonitor(t)
	s := monitor.String()
	if s == "" {
		t.Error("expected non-empty string representation")
	}
}
