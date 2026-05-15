package steampunk

import (
	"fmt"
	"math"
)

// Flywheel stores rotational kinetic energy to smooth out power delivery
// from a steam engine's piston strokes.
type Flywheel struct {
	radius float64  // metres
	mass   float64  // kilograms
	rpm    float64  // current rotational speed
}

// NewFlywheel creates a Flywheel with the given radius (m) and mass (kg).
// Returns an error if either value is non-positive.
func NewFlywheel(radius, mass float64) (*Flywheel, error) {
	if radius <= 0 {
		return nil, fmt.Errorf("flywheel radius must be positive, got %.4f", radius)
	}
	if mass <= 0 {
		return nil, fmt.Errorf("flywheel mass must be positive, got %.4f", mass)
	}
	return &Flywheel{radius: radius, mass: mass}, nil
}

// MomentOfInertia returns I = ½mr² (kg·m²) for a solid disc.
func (f *Flywheel) MomentOfInertia() float64 {
	return 0.5 * f.mass * f.radius * f.radius
}

// KineticEnergy returns the stored rotational energy (J) at the current RPM.
func (f *Flywheel) KineticEnergy() float64 {
	omega := f.AngularVelocity()
	return 0.5 * f.MomentOfInertia() * omega * omega
}

// AngularVelocity returns ω in rad/s for the current RPM.
func (f *Flywheel) AngularVelocity() float64 {
	return f.rpm * 2 * math.Pi / 60
}

// SetRPM updates the rotational speed of the flywheel.
// Negative values are clamped to zero.
func (f *Flywheel) SetRPM(rpm float64) {
	if rpm < 0 {
		rpm = 0
	}
	f.rpm = rpm
}

// RPM returns the current rotational speed in revolutions per minute.
func (f *Flywheel) RPM() float64 {
	return f.rpm
}

// ApplyTorque accelerates or decelerates the flywheel by the given torque (N·m)
// over a time interval dt (s). Returns the new RPM.
func (f *Flywheel) ApplyTorque(torque, dt float64) float64 {
	if dt <= 0 {
		return f.rpm
	}
	// α = τ / I
	alpha := torque / f.MomentOfInertia()
	omega := f.AngularVelocity() + alpha*dt
	if omega < 0 {
		omega = 0
	}
	f.rpm = omega * 60 / (2 * math.Pi)
	return f.rpm
}

// String returns a human-readable description of the flywheel state.
func (f *Flywheel) String() string {
	return fmt.Sprintf("Flywheel(r=%.2fm, m=%.1fkg, I=%.4fkg·m², rpm=%.1f, KE=%.2fJ)",
		f.radius, f.mass, f.MomentOfInertia(), f.rpm, f.KineticEnergy())
}
