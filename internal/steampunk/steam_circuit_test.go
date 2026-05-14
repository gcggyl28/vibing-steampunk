package steampunk

import "testing"

func makeTestCircuit(t *testing.T) (*SteamCircuit, *Boiler, *Pipeline, *Condenser) {
	t.Helper()
	b, err := NewBoiler(100.0, 10.0)
	if err != nil {
		t.Fatalf("NewBoiler: %v", err)
	}
	b.Heat(200.0)
	p, err := NewPipeline("main", 10.0)
	if err != nil {
		t.Fatalf("NewPipeline: %v", err)
	}
	c, err := NewCondenser(8.0, 0.75)
	if err != nil {
		t.Fatalf("NewCondenser: %v", err)
	}
	sc, err := NewSteamCircuit(b, p, c)
	if err != nil {
		t.Fatalf("NewSteamCircuit: %v", err)
	}
	return sc, b, p, c
}

func TestNewSteamCircuit(t *testing.T) {
	sc, _, _, _ := makeTestCircuit(t)
	if sc == nil {
		t.Fatal("expected non-nil SteamCircuit")
	}
}

func TestNewSteamCircuitNilComponents(t *testing.T) {
	b, _ := NewBoiler(100.0, 10.0)
	p, _ := NewPipeline("main", 10.0)
	c, _ := NewCondenser(8.0, 0.75)

	if _, err := NewSteamCircuit(nil, p, c); err == nil {
		t.Error("expected error for nil boiler")
	}
	if _, err := NewSteamCircuit(b, nil, c); err == nil {
		t.Error("expected error for nil pipeline")
	}
	if _, err := NewSteamCircuit(b, p, nil); err == nil {
		t.Error("expected error for nil condenser")
	}
}

func TestSteamCircuitStart(t *testing.T) {
	sc, _, _, _ := makeTestCircuit(t)
	if err := sc.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	if !sc.IsRunning() {
		t.Error("circuit should be running after Start")
	}
}

func TestSteamCircuitStartColdBoiler(t *testing.T) {
	b, _ := NewBoiler(100.0, 10.0) // cold boiler, not heated
	p, _ := NewPipeline("main", 10.0)
	c, _ := NewCondenser(8.0, 0.75)
	sc, _ := NewSteamCircuit(b, p, c)
	if err := sc.Start(); err == nil {
		t.Error("expected error starting circuit with cold boiler")
	}
}

func TestSteamCircuitStop(t *testing.T) {
	sc, _, _, _ := makeTestCircuit(t)
	_ = sc.Start()
	sc.Stop()
	if sc.IsRunning() {
		t.Error("circuit should not be running after Stop")
	}
}

func TestSteamCircuitTick(t *testing.T) {
	sc, _, _, _ := makeTestCircuit(t)
	_ = sc.Start()
	sc.Tick()
	if sc.CondensateReturn() <= 0 {
		t.Error("expected positive condensate return after tick")
	}
}

func TestSteamCircuitTickNotRunning(t *testing.T) {
	sc, _, _, _ := makeTestCircuit(t)
	sc.Tick() // should be a no-op
	if sc.CondensateReturn() != 0 {
		t.Error("expected zero condensate return when not running")
	}
}

func TestSteamCircuitString(t *testing.T) {
	sc, _, _, _ := makeTestCircuit(t)
	if sc.String() == "" {
		t.Error("expected non-empty string")
	}
}
