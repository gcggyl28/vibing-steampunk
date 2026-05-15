package steampunk

import "fmt"

// SafetyValve is an automatic relief valve that opens when pressure
// exceeds a set point, protecting the system from overpressure.
type SafetyValve struct {
	name      string
	setPoint  float64
	blowdown  float64 // pressure drop before valve reseats
	isOpen    bool
	released  float64 // cumulative pressure released
}

// NewSafetyValve creates a SafetyValve that opens at setPoint PSI.
// blowdown is the PSI drop required before the valve reseats (closes).
func NewSafetyValve(name string, setPoint, blowdown float64) (*SafetyValve, error) {
	if setPoint <= 0 {
		return nil, fmt.Errorf("setPoint %.2f must be positive", setPoint)
	}
	if blowdown <= 0 || blowdown >= setPoint {
		return nil, fmt.Errorf("blowdown %.2f must be positive and less than setPoint", blowdown)
	}
	return &SafetyValve{
		name:     name,
		setPoint: setPoint,
		blowdown: blowdown,
	}, nil
}

// Check evaluates the given pressure and opens or closes the valve accordingly.
// Returns the excess pressure vented (0 if valve is closed).
func (sv *SafetyValve) Check(pressure float64) float64 {
	if !sv.isOpen && pressure >= sv.setPoint {
		sv.isOpen = true
	}
	if sv.isOpen && pressure < (sv.setPoint-sv.blowdown) {
		sv.isOpen = false
	}
	if sv.isOpen && pressure > sv.setPoint {
		excess := pressure - sv.setPoint
		sv.released += excess
		return excess
	}
	return 0
}

// IsOpen returns whether the safety valve is currently open.
func (sv *SafetyValve) IsOpen() bool {
	return sv.isOpen
}

// TotalReleased returns the cumulative pressure released through the valve.
func (sv *SafetyValve) TotalReleased() float64 {
	return sv.released
}

// Reset clears the released counter and forces the valve closed.
func (sv *SafetyValve) Reset() {
	sv.isOpen = false
	sv.released = 0
}

// String returns a human-readable description of the safety valve state.
func (sv *SafetyValve) String() string {
	state := "CLOSED"
	if sv.isOpen {
		state = "OPEN"
	}
	return fmt.Sprintf("SafetyValve(%s): setPoint=%.2f PSI, state=%s, released=%.2f",
		sv.name, sv.setPoint, state, sv.released)
}
