// Package steampunk provides a simulation of steampunk mechanical components.
//
// # Components
//
// The package models the following interconnected parts of a steam-powered
// machine:
//
//   - Boiler    – heats water and maintains steam pressure.
//   - Valve     – controls the flow of steam between components.
//   - Pipeline  – connects valves and routes steam to consumers.
//   - Engine    – orchestrates boiler, pipeline and power output.
//   - Gear      – a toothed wheel that meshes with other gears.
//   - GearTrain – a chain of meshed gears that transmits rotation.
//   - Piston    – converts steam pressure into linear mechanical force.
//   - Crankshaft – converts piston strokes into rotational torque.
//
// # Typical Usage
//
//	b, _ := steampunk.NewBoiler("main", 100, 12)
//	v, _ := steampunk.NewValve("inlet", 1.0)
//	p, _ := steampunk.NewPipeline("steam-line", 10)
//	p.AddValve(v)
//
//	piston, _ := steampunk.NewPiston("P1", 10.0, 15.0)
//	crank, _  := steampunk.NewCrankshaft("CS1", 0.05)
//	crank.AttachPiston(piston)
//
//	// Heat boiler, open valve, apply pressure to piston, rotate crank.
//	b.Heat(50)
//	v.Open()
//	piston.ApplyPressure(p.EffectivePressure(b.Pressure))
//	crank.Rotate(0.1)
//	fmt.Println(crank)
package steampunk
