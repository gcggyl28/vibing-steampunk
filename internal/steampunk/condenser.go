package steampunk

import "fmt"

// Condenser converts exhaust steam back into water for reuse in the boiler.
// It models a simple surface condenser with cooling efficiency.
type Condenser struct {
	coolingArea   float64 // square meters
	efficiency    float64 // 0.0 to 1.0
	inletPressure float64 // bar
	outletTemp    float64 // degrees Celsius
}

// NewCondenser creates a Condenser with the given cooling area (m²) and efficiency.
// Efficiency must be between 0 and 1; cooling area must be positive.
func NewCondenser(coolingArea, efficiency float64) (*Condenser, error) {
	if coolingArea <= 0 {
		return nil, fmt.Errorf("cooling area must be positive, got %.2f", coolingArea)
	}
	if efficiency < 0 || efficiency > 1 {
		return nil, fmt.Errorf("efficiency must be between 0 and 1, got %.2f", efficiency)
	}
	return &Condenser{
		coolingArea: coolingArea,
		efficiency:  efficiency,
		outletTemp:  25.0,
	}, nil
}

// ReceiveSteam accepts exhaust steam at the given pressure (bar).
func (c *Condenser) ReceiveSteam(pressure float64) {
	if pressure < 0 {
		pressure = 0
	}
	c.inletPressure = pressure
}

// OutletTemperature returns the condensed water temperature in Celsius.
// Higher cooling area and efficiency yield lower outlet temperatures.
func (c *Condenser) OutletTemperature() float64 {
	base := 100.0 // boiling point at 1 bar
	reduction := c.coolingArea * c.efficiency * 5.0
	temp := base - reduction
	if temp < c.outletTemp {
		return c.outletTemp
	}
	return temp
}

// CondensateFlow returns the estimated condensate flow rate in kg/s
// based on inlet pressure and condenser efficiency.
func (c *Condenser) CondensateFlow() float64 {
	if c.inletPressure <= 0 {
		return 0
	}
	return c.inletPressure * c.coolingArea * c.efficiency * 0.1
}

// IsEffective reports whether the condenser is operating above 50% efficiency
// and receiving steam.
func (c *Condenser) IsEffective() bool {
	return c.efficiency >= 0.5 && c.inletPressure > 0
}

// String returns a human-readable description of the condenser state.
func (c *Condenser) String() string {
	return fmt.Sprintf("Condenser(area=%.1fm² eff=%.0f%% inlet=%.2fbar outlet=%.1f°C)",
		c.coolingArea, c.efficiency*100, c.inletPressure, c.OutletTemperature())
}
