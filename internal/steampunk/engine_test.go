package steampunk

import (
	"testing"
)

func makeTestEngine(t *testing.T) *Engine {
	t.Helper()
	boiler := NewBoiler(150.0, 10.0)
	boiler.Heat(130.0) // bring to operational pressure

	g1, err := NewGear(20, 1.0)
	if err != nil {
		t.Fatalf("failed to create gear: %v", err)
	}
	g2, err := NewGear(40, 1.0)
	if err != nil {
		t.Fatalf("failed to create gear: %v", err)
	}
	gt := NewGearTrain()
	gt.AddGear(g1)
	gt.AddGear(g2)

	return NewEngine(boiler, gt)
}

func TestNewEngine(t *testing.T) {
	e := makeTestEngine(t)
	if e.Running {
		t.Error("new engine should not be running")
	}
}

func TestEngineStart(t *testing.T) {
	e := makeTestEngine(t)
	if err := e.Start(); err != nil {
		t.Errorf("expected engine to start, got error: %v", err)
	}
	if !e.Running {
		t.Error("engine should be running after Start()")
	}
}

func TestEngineStartColdBoiler(t *testing.T) {
	boiler := NewBoiler(150.0, 10.0) // cold boiler, no heat
	gt := NewGearTrain()
	e := NewEngine(boiler, gt)
	if err := e.Start(); err == nil {
		t.Error("expected error starting engine with cold boiler")
	}
}

func TestEngineStop(t *testing.T) {
	e := makeTestEngine(t)
	_ = e.Start()
	e.Stop()
	if e.Running {
		t.Error("engine should not be running after Stop()")
	}
}

func TestEngineOutputTorqueWhenStopped(t *testing.T) {
	e := makeTestEngine(t)
	if torque := e.OutputTorque(); torque != 0 {
		t.Errorf("expected 0 torque when stopped, got %.2f", torque)
	}
}

func TestEngineOutputTorqueWhenRunning(t *testing.T) {
	e := makeTestEngine(t)
	_ = e.Start()
	if torque := e.OutputTorque(); torque <= 0 {
		t.Errorf("expected positive torque when running, got %.2f", torque)
	}
}

func TestEngineOutputRPM(t *testing.T) {
	e := makeTestEngine(t)
	_ = e.Start()
	rpm := e.OutputRPM(100.0)
	// gear ratio is 40/20 = 2.0, so output RPM should be 200
	if rpm != 200.0 {
		t.Errorf("expected output RPM 200.0, got %.1f", rpm)
	}
}

func TestEngineString(t *testing.T) {
	e := makeTestEngine(t)
	s := e.String()
	if len(s) == 0 {
		t.Error("expected non-empty string representation")
	}
}
