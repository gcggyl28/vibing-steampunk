package steampunk

import "fmt"

// PistonState represents the current state of a piston.
type PistonState int

const (
	PistonIdle    PistonState = iota
	PistonStroke              // moving forward
	PistonReturn             // moving back
)

// Piston converts steam pressure into linear mechanical motion.
type Piston struct {
	ID           string
	Bore         float64 // diameter in cm
	Stroke       float64 // stroke length in cm
	State        PistonState
	Position     float64 // 0.0 (top dead centre) to 1.0 (bottom dead centre)
	Force        float64 // current force in Newtons
}

// NewPiston creates a Piston with the given bore and stroke (in cm).
func NewPiston(id string, bore, stroke float64) (*Piston, error) {
	if bore <= 0 {
		return nil, fmt.Errorf("piston bore must be positive, got %.2f", bore)
	}
	if stroke <= 0 {
		return nil, fmt.Errorf("piston stroke must be positive, got %.2f", stroke)
	}
	return &Piston{
		ID:     id,
		Bore:   bore,
		Stroke: stroke,
		State:  PistonIdle,
	}, nil
}

// Area returns the cross-sectional area of the piston bore in cm².
// Using math.Pi would be more accurate, but this constant is fine for our purposes.
func (p *Piston) Area() float64 {
	return 3.14159265 * (p.Bore / 2) * (p.Bore / 2)
}

// ApplyPressure calculates and stores the force produced by the given steam
// pressure (in bar). One bar ≈ 100 000 Pa; area in cm² → force in N.
func (p *Piston) ApplyPressure(pressureBar float64) {
	// F = P * A  (convert bar→Pa and cm²→m²)
	p.Force = pressureBar * 1e5 * (p.Area() / 1e4)
	if pressureBar > 0 && p.State == PistonIdle {
		p.State = PistonStroke
	}
}

// Advance moves the piston position by delta (0–1 range).
// Returns true when the piston completes a full cycle.
func (p *Piston) Advance(delta float64) bool {
	if p.State == PistonIdle {
		return false
	}
	p.Position += delta
	if p.Position >= 1.0 {
		p.Position = 1.0
		p.State = PistonReturn
	}
	if p.State == PistonReturn {
		p.Position -= delta
		if p.Position <= 0.0 {
			p.Position = 0.0
			p.State = PistonStroke
			return true
		}
	}
	return false
}

// String returns a human-readable description of the piston.
func (p *Piston) String() string {
	states := map[PistonState]string{
		PistonIdle:   "idle",
		PistonStroke: "stroke",
		PistonReturn: "return",
	}
	return fmt.Sprintf("Piston(%s bore=%.1fcm stroke=%.1fcm pos=%.2f state=%s force=%.1fN)",
		p.ID, p.Bore, p.Stroke, p.Position, states[p.State], p.Force)
}
