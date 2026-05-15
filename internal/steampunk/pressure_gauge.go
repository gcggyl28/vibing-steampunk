package steampunk

import "fmt"

// PressureGauge monitors pressure levels and provides readings
// with configurable warning and critical thresholds.
type PressureGauge struct {
	name        string
	minPressure float64
	maxPressure float64
	warnLevel   float64
	critLevel   float64
	current     float64
}

// NewPressureGauge creates a new PressureGauge with the given name and pressure range.
// warnRatio and critRatio are fractions of maxPressure (e.g., 0.8 and 0.95).
func NewPressureGauge(name string, minPressure, maxPressure, warnRatio, critRatio float64) (*PressureGauge, error) {
	if maxPressure <= minPressure {
		return nil, fmt.Errorf("maxPressure %.2f must be greater than minPressure %.2f", maxPressure, minPressure)
	}
	if warnRatio <= 0 || warnRatio >= 1 {
		return nil, fmt.Errorf("warnRatio %.2f must be between 0 and 1", warnRatio)
	}
	if critRatio <= warnRatio || critRatio >= 1 {
		return nil, fmt.Errorf("critRatio %.2f must be between warnRatio and 1", critRatio)
	}
	return &PressureGauge{
		name:        name,
		minPressure: minPressure,
		maxPressure: maxPressure,
		warnLevel:   maxPressure * warnRatio,
		critLevel:   maxPressure * critRatio,
		current:     minPressure,
	}, nil
}

// Update sets the current pressure reading.
func (g *PressureGauge) Update(pressure float64) {
	g.current = pressure
}

// Reading returns the current pressure value.
func (g *PressureGauge) Reading() float64 {
	return g.current
}

// IsWarning returns true when pressure exceeds the warning threshold.
func (g *PressureGauge) IsWarning() bool {
	return g.current >= g.warnLevel
}

// IsCritical returns true when pressure exceeds the critical threshold.
func (g *PressureGauge) IsCritical() bool {
	return g.current >= g.critLevel
}

// IsOverpressure returns true when pressure exceeds the maximum safe level.
func (g *PressureGauge) IsOverpressure() bool {
	return g.current > g.maxPressure
}

// Status returns a human-readable status string.
func (g *PressureGauge) Status() string {
	switch {
	case g.IsOverpressure():
		return "OVERPRESSURE"
	case g.IsCritical():
		return "CRITICAL"
	case g.IsWarning():
		return "WARNING"
	default:
		return "NORMAL"
	}
}

// String returns a formatted gauge reading.
func (g *PressureGauge) String() string {
	return fmt.Sprintf("PressureGauge(%s): %.2f PSI [%s]", g.name, g.current, g.Status())
}
