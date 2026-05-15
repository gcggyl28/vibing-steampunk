package steampunk

import "fmt"

// FlywheelCoupling connects a Crankshaft to a Flywheel, transferring torque
// and synchronising rotational speed between the two components.
type FlywheelCoupling struct {
	crankshaft *Crankshaft
	flywheel   *Flywheel
	gearRatio  float64 // flywheel RPM = crankshaft RPM * gearRatio
}

// NewFlywheelCoupling creates a coupling between a crankshaft and flywheel.
// gearRatio must be positive; 1.0 means direct drive.
func NewFlywheelCoupling(cs *Crankshaft, fw *Flywheel, gearRatio float64) (*FlywheelCoupling, error) {
	if cs == nil {
		return nil, fmt.Errorf("crankshaft must not be nil")
	}
	if fw == nil {
		return nil, fmt.Errorf("flywheel must not be nil")
	}
	if gearRatio <= 0 {
		return nil, fmt.Errorf("gear ratio must be positive, got %.4f", gearRatio)
	}
	return &FlywheelCoupling{
		crankshaft: cs,
		flywheel:   fw,
		gearRatio:  gearRatio,
	}, nil
}

// Update transfers the crankshaft's current torque to the flywheel over
// the time step dt (seconds), then synchronises the flywheel's displayed RPM.
func (fc *FlywheelCoupling) Update(dt float64) {
	if dt <= 0 {
		return
	}
	torque := fc.crankshaft.Torque()
	fc.flywheel.ApplyTorque(torque, dt)
}

// FlywheelRPM returns the flywheel's RPM adjusted by the gear ratio.
func (fc *FlywheelCoupling) FlywheelRPM() float64 {
	return fc.crankshaft.RPM() * fc.gearRatio
}

// SyncRPM forces the flywheel's RPM to match the crankshaft via the gear ratio.
// Useful for initialisation or hard resets.
func (fc *FlywheelCoupling) SyncRPM() {
	fc.flywheel.SetRPM(fc.FlywheelRPM())
}

// GearRatio returns the configured gear ratio.
func (fc *FlywheelCoupling) GearRatio() float64 {
	return fc.gearRatio
}

// String returns a summary of the coupling state.
func (fc *FlywheelCoupling) String() string {
	return fmt.Sprintf("FlywheelCoupling(ratio=%.2f, crankRPM=%.1f, flywheelRPM=%.1f)",
		fc.gearRatio, fc.crankshaft.RPM(), fc.flywheel.RPM())
}
