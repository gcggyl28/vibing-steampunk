package steampunk

import "fmt"

// SteamCircuit models a closed-loop steam system connecting a Boiler,
// Pipeline, and Condenser. It tracks circuit state and condensate return.
type SteamCircuit struct {
	boiler    *Boiler
	pipeline  *Pipeline
	condenser *Condenser
	running   bool
}

// NewSteamCircuit assembles a SteamCircuit from its components.
// All components must be non-nil.
func NewSteamCircuit(b *Boiler, p *Pipeline, c *Condenser) (*SteamCircuit, error) {
	if b == nil {
		return nil, fmt.Errorf("boiler must not be nil")
	}
	if p == nil {
		return nil, fmt.Errorf("pipeline must not be nil")
	}
	if c == nil {
		return nil, fmt.Errorf("condenser must not be nil")
	}
	return &SteamCircuit{boiler: b, pipeline: p, condenser: c}, nil
}

// Start activates the steam circuit if the boiler is operational.
func (sc *SteamCircuit) Start() error {
	if !sc.boiler.IsOperational() {
		return fmt.Errorf("cannot start circuit: boiler is not operational")
	}
	sc.running = true
	return nil
}

// Stop shuts down the steam circuit.
func (sc *SteamCircuit) Stop() {
	sc.running = false
}

// IsRunning reports whether the circuit is active.
func (sc *SteamCircuit) IsRunning() bool {
	return sc.running
}

// Tick advances the circuit by one simulation step, routing steam
// from the boiler through the pipeline into the condenser.
func (sc *SteamCircuit) Tick() {
	if !sc.running {
		return
	}
	pressure := sc.pipeline.EffectivePressure(sc.boiler.Pressure())
	sc.condenser.ReceiveSteam(pressure)
}

// CondensateReturn returns the condensate flow rate available to
// feed back into the boiler (kg/s).
func (sc *SteamCircuit) CondensateReturn() float64 {
	if !sc.running {
		return 0
	}
	return sc.condenser.CondensateFlow()
}

// String returns a summary of the steam circuit state.
func (sc *SteamCircuit) String() string {
	status := "stopped"
	if sc.running {
		status = "running"
	}
	return fmt.Sprintf("SteamCircuit(%s boiler=%.1fbar condensate=%.3fkg/s)",
		status, sc.boiler.Pressure(), sc.CondensateReturn())
}
