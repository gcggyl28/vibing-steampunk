package steampunk

import (
	"errors"
	"fmt"
)

// Pipeline connects a Boiler to an Engine via a set of Valves.
type Pipeline struct {
	boiler  *Boiler
	engine  *Engine
	valves  []*Valve
	pressureLoss float64 // pressure drop across the pipeline (bar)
}

// NewPipeline creates a Pipeline connecting the given boiler and engine.
func NewPipeline(boiler *Boiler, engine *Engine) *Pipeline {
	return &Pipeline{
		boiler:       boiler,
		engine:       engine,
		pressureLoss: 0.5,
	}
}

// AddValve appends a valve to the pipeline.
func (p *Pipeline) AddValve(v *Valve) {
	p.valves = append(p.valves, v)
}

// IsFlowing returns true when all valves are at least partially open.
func (p *Pipeline) IsFlowing() bool {
	if len(p.valves) == 0 {
		return false
	}
	for _, v := range p.valves {
		if v.State() == ValveClosed {
			return false
		}
	}
	return true
}

// EffectivePressure returns the pressure delivered to the engine after losses.
func (p *Pipeline) EffectivePressure() float64 {
	if !p.IsFlowing() {
		return 0
	}
	// Minimum valve opening limits flow
	minOpen := 1.0
	for _, v := range p.valves {
		if v.openPct < minOpen {
			minOpen = v.openPct
		}
	}
	pressure := p.boiler.Pressure() - p.pressureLoss
	if pressure < 0 {
		pressure = 0
	}
	return pressure * minOpen
}

// Engage opens all valves and attempts to start the engine.
func (p *Pipeline) Engage() error {
	if len(p.valves) == 0 {
		return errors.New("pipeline: no valves configured")
	}
	for _, v := range p.valves {
		v.Open()
	}
	if err := p.engine.Start(); err != nil {
		for _, v := range p.valves {
			v.Close()
		}
		return fmt.Errorf("pipeline: engage failed: %w", err)
	}
	return nil
}

// Disengage closes all valves and stops the engine.
func (p *Pipeline) Disengage() {
	for _, v := range p.valves {
		v.Close()
	}
	p.engine.Stop()
}

// String returns a summary of the pipeline state.
func (p *Pipeline) String() string {
	return fmt.Sprintf("Pipeline[flowing=%v pressure=%.2fbar valves=%d]",
		p.IsFlowing(), p.EffectivePressure(), len(p.valves))
}
