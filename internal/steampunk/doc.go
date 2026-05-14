// Package steampunk provides primitives for simulating a steampunk mechanical
// system, including gears, gear trains, boilers, and steam engines.
//
// # Gears
//
// A [Gear] is a toothed wheel with a given number of teeth and a module (pitch
// size). Two gears can mesh together when they share the same module. Use
// [NewGear] to create a gear and [Gear.MeshWith] to validate meshing.
//
// # Gear Trains
//
// A [GearTrain] is an ordered sequence of meshing gears. The overall gear
// ratio is the product of each successive pair's tooth ratio. Use
// [NewGearTrain] to create a train and [GearTrain.AddGear] to extend it.
//
// # Boilers
//
// A [Boiler] generates steam pressure from heat. Pressure drives torque
// output. Use [NewBoiler] to create a boiler, [Boiler.Heat] to raise
// temperature, and [Boiler.Torque] to read available torque.
//
// # Engines
//
// An [Engine] couples a [Boiler] to a [GearTrain] to deliver mechanical
// power at the output shaft. Use [NewEngine] to assemble an engine,
// [Engine.Start] to begin operation, and [Engine.OutputTorque] /
// [Engine.OutputRPM] to read shaft characteristics.
//
// Example:
//
//	boiler := steampunk.NewBoiler(150.0, 20.0)
//	boiler.Heat(140.0)
//
//	g1, _ := steampunk.NewGear(20, 1.0)
//	g2, _ := steampunk.NewGear(60, 1.0)
//	gt := steampunk.NewGearTrain()
//	gt.AddGear(g1)
//	gt.AddGear(g2)
//
//	engine := steampunk.NewEngine(boiler, gt)
//	if err := engine.Start(); err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(engine)
package steampunk
