package steampunk

import "fmt"

// ValveState represents the current state of a valve.
type ValveState int

const (
	ValveClosed ValveState = iota
	ValveOpen
	ValvePartial
)

// Valve controls steam flow between components.
type Valve struct {
	name     string
	state    ValveState
	openPct  float64 // 0.0 to 1.0
	maxFlow  float64 // maximum flow rate in kg/s
}

// NewValve creates a new Valve with the given name and maximum flow rate.
func NewValve(name string, maxFlow float64) *Valve {
	if maxFlow <= 0 {
		maxFlow = 1.0
	}
	return &Valve{
		name:    name,
		state:   ValveClosed,
		openPct: 0.0,
		maxFlow: maxFlow,
	}
}

// Open fully opens the valve.
func (v *Valve) Open() {
	v.state = ValveOpen
	v.openPct = 1.0
}

// Close fully closes the valve.
func (v *Valve) Close() {
	v.state = ValveClosed
	v.openPct = 0.0
}

// SetPosition sets the valve opening percentage (0.0 to 1.0).
func (v *Valve) SetPosition(pct float64) {
	if pct < 0 {
		pct = 0
	}
	if pct > 1 {
		pct = 1
	}
	v.openPct = pct
	switch {
	case pct == 0:
		v.state = ValveClosed
	case pct == 1:
		v.state = ValveOpen
	default:
		v.state = ValvePartial
	}
}

// FlowRate returns the current steam flow rate in kg/s.
func (v *Valve) FlowRate() float64 {
	return v.maxFlow * v.openPct
}

// State returns the current valve state.
func (v *Valve) State() ValveState {
	return v.state
}

// String returns a human-readable description of the valve.
func (v *Valve) String() string {
	stateStr := map[ValveState]string{
		ValveClosed:  "CLOSED",
		ValveOpen:    "OPEN",
		ValvePartial: "PARTIAL",
	}[v.state]
	return fmt.Sprintf("Valve(%s)[%s %.0f%% flow=%.2fkg/s]",
		v.name, stateStr, v.openPct*100, v.FlowRate())
}
