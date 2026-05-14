// Package steampunk provides a simulation of a Victorian-era steam-powered
// mechanical system. It models the core components of a steampunk engine:
//
//   - Boiler: generates steam pressure by heating water
//   - Valve: controls steam flow with adjustable positions
//   - Pipeline: connects components and routes steam through valves
//   - Piston: converts steam pressure into linear mechanical force
//   - Crankshaft: converts piston force into rotational torque
//   - Gear / GearTrain: transmits and transforms rotational motion
//   - Governor: regulates engine speed via feedback-controlled valve
//   - Engine: orchestrates the full power cycle
//   - Condenser: recovers exhaust steam as condensate for boiler reuse
//   - SteamCircuit: models a closed-loop boiler–pipeline–condenser system
//
// Components are designed to be composed into larger assemblies. Each
// component exposes a String() method for human-readable status output,
// suitable for terminal dashboards or logging.
//
// Units used throughout:
//
//	Pressure : bar
//	Temperature : degrees Celsius
//	Force : Newtons
//	Torque : Newton-metres
//	Angular velocity : radians per second
//	Flow rate : kg/s
//	Area : square metres
//
Example usage:
//
//	b, _ := steampunk.NewBoiler(200.0, 15.0)
//	b.Heat(250.0)
//	p, _ := steampunk.NewPipeline("main", 15.0)
//	c, _ := steampunk.NewCondenser(10.0, 0.8)
//	sc, _ := steampunk.NewSteamCircuit(b, p, c)
//	_ = sc.Start()
//	sc.Tick()
package steampunk
