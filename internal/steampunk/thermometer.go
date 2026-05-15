package steampunk

import "fmt"

// Thermometer measures temperature in a steampunk system.
// It supports configurable min/max ranges and threshold-based status reporting.
type Thermometer struct {
	name     string
	minTemp  float64
	maxTemp  float64
	warnTemp float64
	critTemp float64
	current  float64
}

// NewThermometer creates a Thermometer with the given name and temperature range.
// warnRatio and critRatio are fractions of maxTemp at which warnings/critical alerts trigger.
func NewThermometer(name string, minTemp, maxTemp, warnRatio, critRatio float64) (*Thermometer, error) {
	if minTemp >= maxTemp {
		return nil, fmt.Errorf("thermometer: minTemp %.2f must be less than maxTemp %.2f", minTemp, maxTemp)
	}
	if warnRatio <= 0 || warnRatio >= 1 {
		return nil, fmt.Errorf("thermometer: warnRatio %.2f must be between 0 and 1", warnRatio)
	}
	if critRatio <= warnRatio || critRatio >= 1 {
		return nil, fmt.Errorf("thermometer: critRatio %.2f must be greater than warnRatio and less than 1", critRatio)
	}
	return &Thermometer{
		name:     name,
		minTemp:  minTemp,
		maxTemp:  maxTemp,
		warnTemp: minTemp + (maxTemp-minTemp)*warnRatio,
		critTemp: minTemp + (maxTemp-minTemp)*critRatio,
		current:  minTemp,
	}, nil
}

// SetTemperature updates the current temperature reading, clamped to [minTemp, maxTemp].
func (t *Thermometer) SetTemperature(temp float64) {
	if temp < t.minTemp {
		temp = t.minTemp
	}
	if temp > t.maxTemp {
		temp = t.maxTemp
	}
	t.current = temp
}

// Temperature returns the current temperature reading.
func (t *Thermometer) Temperature() float64 {
	return t.current
}

// IsWarning returns true if the current temperature is at or above the warning threshold.
func (t *Thermometer) IsWarning() bool {
	return t.current >= t.warnTemp
}

// IsCritical returns true if the current temperature is at or above the critical threshold.
func (t *Thermometer) IsCritical() bool {
	return t.current >= t.critTemp
}

// Status returns a human-readable status string for the thermometer.
func (t *Thermometer) Status() string {
	switch {
	case t.IsCritical():
		return fmt.Sprintf("%s: CRITICAL %.1f°C (crit=%.1f)", t.name, t.current, t.critTemp)
	case t.IsWarning():
		return fmt.Sprintf("%s: WARNING %.1f°C (warn=%.1f)", t.name, t.current, t.warnTemp)
	default:
		return fmt.Sprintf("%s: OK %.1f°C", t.name, t.current)
	}
}

// String implements fmt.Stringer.
func (t *Thermometer) String() string {
	return t.Status()
}
