package steampunk

import "fmt"

// PressureMonitor combines a PressureGauge and SafetyValve to provide
// integrated pressure monitoring and automatic overpressure protection.
type PressureMonitor struct {
	gauge       *PressureGauge
	safetyValve *SafetyValve
}

// NewPressureMonitor creates a PressureMonitor using the provided gauge and safety valve.
func NewPressureMonitor(gauge *PressureGauge, sv *SafetyValve) (*PressureMonitor, error) {
	if gauge == nil {
		return nil, fmt.Errorf("gauge must not be nil")
	}
	if sv == nil {
		return nil, fmt.Errorf("safety valve must not be nil")
	}
	return &PressureMonitor{
		gauge:       gauge,
		safetyValve: sv,
	}, nil
}

// Update feeds a new pressure reading into the gauge and safety valve.
// Returns the amount of pressure vented by the safety valve (0 if not triggered).
func (pm *PressureMonitor) Update(pressure float64) float64 {
	pm.gauge.Update(pressure)
	return pm.safetyValve.Check(pressure)
}

// GaugeReading returns the current pressure from the gauge.
func (pm *PressureMonitor) GaugeReading() float64 {
	return pm.gauge.Reading()
}

// GaugeStatus returns the status string from the gauge.
func (pm *PressureMonitor) GaugeStatus() string {
	return pm.gauge.Status()
}

// SafetyValveOpen returns whether the safety valve is currently open.
func (pm *PressureMonitor) SafetyValveOpen() bool {
	return pm.safetyValve.IsOpen()
}

// IsOperational returns true when pressure is within safe operating bounds.
func (pm *PressureMonitor) IsOperational() bool {
	return !pm.gauge.IsOverpressure() && !pm.safetyValve.IsOpen()
}

// String returns a combined status of the monitor.
func (pm *PressureMonitor) String() string {
	return fmt.Sprintf("PressureMonitor: %s | %s", pm.gauge.String(), pm.safetyValve.String())
}
