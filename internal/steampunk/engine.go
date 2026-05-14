package steampunk

import "fmt"

// Engine combines a Boiler with a GearTrain to form a complete steam engine.
type Engine struct {
	Boiler    *Boiler
	GearTrain *GearTrain
	Running   bool
}

// NewEngine creates a new Engine with the given boiler and gear train.
func NewEngine(boiler *Boiler, gearTrain *GearTrain) *Engine {
	return &Engine{
		Boiler:    boiler,
		GearTrain: gearTrain,
		Running:   false,
	}
}

// Start attempts to start the engine. Returns an error if the boiler
// is not operational.
func (e *Engine) Start() error {
	if !e.Boiler.IsOperational() {
		return fmt.Errorf("cannot start engine: boiler pressure %.1f PSI is insufficient",
			e.Boiler.CurrentPressure)
	}
	e.Running = true
	return nil
}

// Stop shuts down the engine and vents boiler pressure.
func (e *Engine) Stop() {
	e.Running = false
	e.Boiler.Vent(e.Boiler.CurrentPressure * 0.1)
}

// OutputTorque returns the torque delivered at the final gear of the train,
// scaled by the gear ratio. Returns 0 if the engine is not running.
func (e *Engine) OutputTorque() float64 {
	if !e.Running {
		return 0
	}
	ratio := e.GearTrain.Ratio()
	if ratio == 0 {
		return 0
	}
	return e.Boiler.Torque() / ratio
}

// OutputRPM returns the RPM at the output shaft given a drive RPM.
// Returns 0 if the engine is not running.
func (e *Engine) OutputRPM(driveRPM float64) float64 {
	if !e.Running {
		return 0
	}
	return driveRPM * e.GearTrain.Ratio()
}

// String returns a human-readable description of the engine state.
func (e *Engine) String() string {
	status := "stopped"
	if e.Running {
		status = "running"
	}
	return fmt.Sprintf("Engine{status=%s, outputTorque=%.2f N·m, %s}",
		status, e.OutputTorque(), e.Boiler)
}
