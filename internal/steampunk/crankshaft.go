package steampunk

import (
	"fmt"
	"math"
)

// Crankshaft converts the linear force of one or more pistons into rotational
// motion, expressed as torque and angular velocity.
type Crankshaft struct {
	ID          string
	CrankRadius float64   // metres
	Pistons     []*Piston
	Angle       float64   // current angle in radians
	RPM         float64   // current rotational speed
}

// NewCrankshaft creates a Crankshaft with the given crank radius (in metres).
func NewCrankshaft(id string, crankRadius float64) (*Crankshaft, error) {
	if crankRadius <= 0 {
		return nil, fmt.Errorf("crank radius must be positive, got %.4f", crankRadius)
	}
	return &Crankshaft{
		ID:          id,
		CrankRadius: crankRadius,
	}, nil
}

// AttachPiston adds a piston to the crankshaft.
func (c *Crankshaft) AttachPiston(p *Piston) {
	c.Pistons = append(c.Pistons, p)
}

// Torque returns the instantaneous torque (N·m) contributed by all attached
// pistons at the current crank angle.
func (c *Crankshaft) Torque() float64 {
	var total float64
	offset := 2 * math.Pi / math.Max(1, float64(len(c.Pistons)))
	for i, p := range c.Pistons {
		angle := c.Angle + float64(i)*offset
		// Effective tangential force = F * sin(angle)
		total += p.Force * math.Abs(math.Sin(angle)) * c.CrankRadius
	}
	return total
}

// Rotate advances the crankshaft by deltaAngle radians and propagates the
// position update to all attached pistons.
func (c *Crankshaft) Rotate(deltaAngle float64) {
	c.Angle = math.Mod(c.Angle+deltaAngle, 2*math.Pi)
	// Map crank angle to piston position (0→1 via sine)
	delta := (math.Sin(deltaAngle) + 1) / 2 * 0.1
	for _, p := range c.Pistons {
		p.Advance(delta)
	}
	// Approximate RPM from deltaAngle (assuming 60 fps tick)
	c.RPM = (deltaAngle / (2 * math.Pi)) * 60 * 60
}

// String returns a human-readable summary of the crankshaft.
func (c *Crankshaft) String() string {
	return fmt.Sprintf("Crankshaft(%s radius=%.3fm pistons=%d angle=%.2frad torque=%.2fNm rpm=%.1f)",
		c.ID, c.CrankRadius, len(c.Pistons), c.Angle, c.Torque(), c.RPM)
}
