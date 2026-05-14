package steampunk

import "fmt"

// GearTrain represents a series of meshed gears transmitting motion.
type GearTrain struct {
	gears []*Gear
}

// NewGearTrain creates a new empty GearTrain.
func NewGearTrain() *GearTrain {
	return &GearTrain{}
}

// AddGear appends a gear to the train and propagates RPM from the previous gear.
func (gt *GearTrain) AddGear(g *Gear) {
	if len(gt.gears) > 0 {
		prev := gt.gears[len(gt.gears)-1]
		g.RPM = prev.MeshWith(g)
	}
	gt.gears = append(gt.gears, g)
}

// OutputRPM returns the RPM of the last gear in the train.
func (gt *GearTrain) OutputRPM() (float64, error) {
	if len(gt.gears) == 0 {
		return 0, fmt.Errorf("gear train is empty")
	}
	return gt.gears[len(gt.gears)-1].RPM, nil
}

// GearRatio returns the overall ratio from first to last gear.
func (gt *GearTrain) GearRatio() (float64, error) {
	if len(gt.gears) < 2 {
		return 0, fmt.Errorf("need at least two gears to compute ratio")
	}
	input := gt.gears[0].RPM
	if input == 0 {
		return 0, fmt.Errorf("input RPM is zero")
	}
	output, _ := gt.OutputRPM()
	return output / input, nil
}

// Gears returns a copy of the gear slice for inspection.
func (gt *GearTrain) Gears() []*Gear {
	copy := make([]*Gear, len(gt.gears))
	for i, g := range gt.gears {
		copy[i] = g
	}
	return copy
}
