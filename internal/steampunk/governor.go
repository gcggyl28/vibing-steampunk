package steampunk

import "fmt"

// Governor is a centrifugal governor that regulates engine speed
// by adjusting steam flow based on rotational velocity.
type Governor struct {
	targetRPM   float64
	currentRPM  float64
	sensitivity float64
	valve       *Valve
	active      bool
}

// NewGovernor creates a new Governor with the given target RPM and a
// control valve. sensitivity controls how aggressively the governor
// responds to speed deviations (0 < sensitivity <= 1).
func NewGovernor(targetRPM, sensitivity float64, valve *Valve) (*Governor, error) {
	if targetRPM <= 0 {
		return nil, fmt.Errorf("target RPM must be positive, got %.2f", targetRPM)
	}
	if sensitivity <= 0 || sensitivity > 1 {
		return nil, fmt.Errorf("sensitivity must be in (0, 1], got %.2f", sensitivity)
	}
	if valve == nil {
		return nil, fmt.Errorf("governor requires a non-nil valve")
	}
	return &Governor{
		targetRPM:   targetRPM,
		sensitivity: sensitivity,
		valve:       valve,
	}, nil
}

// Activate engages the governor control loop.
func (g *Governor) Activate() {
	g.active = true
}

// Deactivate disengages the governor, leaving the valve at its current position.
func (g *Governor) Deactivate() {
	g.active = false
}

// IsActive reports whether the governor is currently engaged.
func (g *Governor) IsActive() bool {
	return g.active
}

// Update feeds the current RPM into the governor and adjusts the
// control valve position accordingly. It returns the new valve position.
func (g *Governor) Update(currentRPM float64) float64 {
	g.currentRPM = currentRPM
	if !g.active {
		return g.valve.Position()
	}

	error := (g.targetRPM - currentRPM) / g.targetRPM
	adjustment := error * g.sensitivity
	newPosition := g.valve.Position() + adjustment

	if newPosition > 1.0 {
		newPosition = 1.0
	}
	if newPosition < 0.0 {
		newPosition = 0.0
	}

	_ = g.valve.SetPosition(newPosition)
	return newPosition
}

// TargetRPM returns the governor's target rotational speed.
func (g *Governor) TargetRPM() float64 {
	return g.targetRPM
}

// CurrentRPM returns the last observed RPM.
func (g *Governor) CurrentRPM() float64 {
	return g.currentRPM
}

// String returns a human-readable description of the governor state.
func (g *Governor) String() string {
	status := "inactive"
	if g.active {
		status = "active"
	}
	return fmt.Sprintf("Governor{target=%.1f rpm, current=%.1f rpm, valve=%.2f, status=%s}",
		g.targetRPM, g.currentRPM, g.valve.Position(), status)
}
