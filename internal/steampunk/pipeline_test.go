package steampunk

import "testing"

func makeTestPipeline(t *testing.T) (*Pipeline, *Boiler, *Engine) {
	t.Helper()
	b := NewBoiler(10.0, 200.0)
	b.Heat(180.0)
	e := NewEngine(b)
	p := NewPipeline(b, e)
	return p, b, e
}

func TestNewPipeline(t *testing.T) {
	p, _, _ := makeTestPipeline(t)
	if p.boiler == nil || p.engine == nil {
		t.Fatal("expected non-nil boiler and engine")
	}
	if p.IsFlowing() {
		t.Error("new pipeline with no valves should not be flowing")
	}
}

func TestPipelineAddValve(t *testing.T) {
	p, _, _ := makeTestPipeline(t)
	v := NewValve("v1", 3.0)
	p.AddValve(v)
	if len(p.valves) != 1 {
		t.Errorf("expected 1 valve, got %d", len(p.valves))
	}
}

func TestPipelineIsFlowing(t *testing.T) {
	p, _, _ := makeTestPipeline(t)
	v := NewValve("v1", 3.0)
	p.AddValve(v)
	if p.IsFlowing() {
		t.Error("closed valve should prevent flow")
	}
	v.Open()
	if !p.IsFlowing() {
		t.Error("open valve should allow flow")
	}
}

func TestPipelineEffectivePressure(t *testing.T) {
	p, b, _ := makeTestPipeline(t)
	v := NewValve("v1", 3.0)
	p.AddValve(v)
	v.Open()
	expected := b.Pressure() - p.pressureLoss
	if got := p.EffectivePressure(); got != expected {
		t.Errorf("expected %.2f, got %.2f", expected, got)
	}
}

func TestPipelineEngageDisengage(t *testing.T) {
	p, _, _ := makeTestPipeline(t)
	v := NewValve("v1", 3.0)
	p.AddValve(v)
	if err := p.Engage(); err != nil {
		t.Fatalf("Engage failed: %v", err)
	}
	if !p.IsFlowing() {
		t.Error("pipeline should be flowing after Engage")
	}
	p.Disengage()
	if p.IsFlowing() {
		t.Error("pipeline should not be flowing after Disengage")
	}
}

func TestPipelineEngageNoValves(t *testing.T) {
	p, _, _ := makeTestPipeline(t)
	if err := p.Engage(); err == nil {
		t.Error("expected error when engaging pipeline with no valves")
	}
}

func TestPipelineString(t *testing.T) {
	p, _, _ := makeTestPipeline(t)
	s := p.String()
	if s == "" {
		t.Error("expected non-empty string")
	}
}
