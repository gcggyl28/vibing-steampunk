package steampunk

import (
	"fmt"
	"math"
)

// Boiler represents a steam boiler that provides pressure to drive gear trains.
type Boiler struct {
	MaxPressure     float64 // in PSI
	CurrentPressure float64 // in PSI
	Temperature     float64 // in Celsius
	Volume          float64 // in liters
	Efficiency      float64 // 0.0 to 1.0
}

// NewBoiler creates a new Boiler with the given max pressure and volume.
func NewBoiler(maxPressure, volume float64) *Boiler {
	if maxPressure <= 0 {
		maxPressure = 100.0
	}
	if volume <= 0 {
		volume = 10.0
	}
	return &Boiler{
		MaxPressure:     maxPressure,
		CurrentPressure: 0,
		Temperature:     20.0,
		Volume:          volume,
		Efficiency:      0.85,
	}
}

// Heat increases the boiler temperature and builds pressure.
func (b *Boiler) Heat(deltaTemp float64) {
	b.Temperature += deltaTemp
	// Simplified steam pressure model based on temperature
	if b.Temperature > 100 {
		b.CurrentPressure = b.Efficiency * (b.Temperature - 100) * 0.5
		if b.CurrentPressure > b.MaxPressure {
			b.CurrentPressure = b.MaxPressure
		}
	}
}

// Torque calculates the output torque available from the boiler pressure.
func (b *Boiler) Torque() float64 {
	// Torque proportional to pressure and volume
	return b.CurrentPressure * math.Sqrt(b.Volume) * b.Efficiency
}

// IsOperational returns true if the boiler has enough pressure to operate.
func (b *Boiler) IsOperational() bool {
	return b.CurrentPressure >= 10.0
}

// Vent releases pressure by the given amount.
func (b *Boiler) Vent(amount float64) {
	b.CurrentPressure -= amount
	if b.CurrentPressure < 0 {
		b.CurrentPressure = 0
	}
}

// String returns a human-readable description of the boiler state.
func (b *Boiler) String() string {
	return fmt.Sprintf("Boiler{pressure=%.1f/%.1f PSI, temp=%.1f°C, torque=%.2f N·m}",
		b.CurrentPressure, b.MaxPressure, b.Temperature, b.Torque())
}
