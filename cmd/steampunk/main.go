// Command steampunk is the main entry point for the vibing-steampunk simulation.
// It demonstrates a complete steampunk steam engine system with all components
// working together: boiler, pistons, crankshaft, governor, condenser, and gears.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func main() {
	logger := log.New(os.Stdout, "[steampunk] ", log.LstdFlags)

	logger.Println("Initialising steampunk engine system...")

	// Build the boiler
	boiler, err := steampunk.NewBoiler(10.0, 500.0, 15.0)
	if err != nil {
		logger.Fatalf("failed to create boiler: %v", err)
	}

	// Build the condenser
	condenser, err := steampunk.NewCondenser(2.0, 0.85)
	if err != nil {
		logger.Fatalf("failed to create condenser: %v", err)
	}

	// Build the governor control valve
	governorValve, err := steampunk.NewValve("governor-valve", 5.0)
	if err != nil {
		logger.Fatalf("failed to create governor valve: %v", err)
	}
	if err := governorValve.Open(); err != nil {
		logger.Fatalf("failed to open governor valve: %v", err)
	}

	// Build the governor
	governor, err := steampunk.NewGovernor(1500.0, 0.1, governorValve)
	if err != nil {
		logger.Fatalf("failed to create governor: %v", err)
	}

	// Build the steam circuit
	circuit, err := steampunk.NewSteamCircuit(boiler, condenser, governorValve)
	if err != nil {
		logger.Fatalf("failed to create steam circuit: %v", err)
	}

	// Build the engine
	engine, err := steampunk.NewEngine(boiler, governorValve)
	if err != nil {
		logger.Fatalf("failed to create engine: %v", err)
	}

	// Build the gear train
	gearTrain := steampunk.NewGearTrain()
	drive, err := steampunk.NewGear("drive", steampunk.Large)
	if err != nil {
		logger.Fatalf("failed to create drive gear: %v", err)
	}
	driven, err := steampunk.NewGear("driven", steampunk.Small)
	if err != nil {
		logger.Fatalf("failed to create driven gear: %v", err)
	}
	gearTrain.Add(drive)
	gearTrain.Add(driven)

	// Heat the boiler to operating temperature
	logger.Println("Heating boiler...")
	for boiler.Temperature() < 150.0 {
		if err := boiler.Heat(10.0); err != nil {
			logger.Fatalf("boiler heating error: %v", err)
		}
	}
	logger.Printf("Boiler ready: %.1f°C, %.2f bar", boiler.Temperature(), boiler.Pressure())

	// Start the circuit and engine
	if err := circuit.Start(); err != nil {
		logger.Fatalf("failed to start steam circuit: %v", err)
	}
	logger.Println("Steam circuit running.")

	if err := engine.Start(); err != nil {
		logger.Fatalf("failed to start engine: %v", err)
	}
	logger.Println("Engine started.")

	// Simulate a few regulation cycles
	for i := range 5 {
		simRPM := 1200.0 + float64(i)*100.0
		governor.Regulate(simRPM)
		logger.Printf(
			"cycle %d — RPM: %.0f, valve: %.0f%%, pressure: %.2f bar",
			i+1, simRPM,
			governorValve.Position()*100,
			boiler.Pressure(),
		)
		time.Sleep(50 * time.Millisecond)
	}

	// Print gear train summary
	fmt.Println()
	fmt.Println(gearTrain.String())

	// Shutdown
	engine.Stop()
	logger.Println("Engine stopped. Goodbye.")
}
