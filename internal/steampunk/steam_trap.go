// Package steampunk provides a simulation of steampunk mechanical components.
package steampunk

import (
	"fmt"
	"math"
)

// SteamTrap removes condensate and non-condensable gases from a steam system
// without allowing live steam to escape. It acts as a thermostatic device
// that opens when condensate is present and closes when steam approaches.
type SteamTrap struct {
	// condensateRate is the rate at which condensate is discharged (kg/s)
	condensateRate float64
	// steamLoss is the fraction of steam lost through the trap (0.0–1.0)
	steamLoss float64
	// openTemp is the temperature threshold below which the trap opens (°C)
	openTemp float64
	// currentTemp is the current temperature at the trap inlet (°C)
	currentTemp float64
	// isOpen indicates whether the trap is currently discharging
	isOpen bool
	// failed indicates a failed-open or failed-closed condition
	failed bool
	// failedOpen means the trap is stuck open (steam blowing through)
	failedOpen bool
}

// NewSteamTrap creates a new SteamTrap with the given parameters.
// openTemp is the temperature (°C) at or below which the trap opens to
// discharge condensate. condensateRate is the maximum discharge rate (kg/s).
// steamLoss is the fractional steam loss when the trap is open (0.0–0.05 typical).
func NewSteamTrap(openTemp, condensateRate, steamLoss float64) (*SteamTrap, error) {
	if openTemp <= 0 {
		return nil, fmt.Errorf("steam trap: openTemp must be positive, got %.2f", openTemp)
	}
	if condensateRate <= 0 {
		return nil, fmt.Errorf("steam trap: condensateRate must be positive, got %.2f", condensateRate)
	}
	if steamLoss < 0 || steamLoss > 1 {
		return nil, fmt.Errorf("steam trap: steamLoss must be in [0, 1], got %.2f", steamLoss)
	}
	return &SteamTrap{
		openTemp:       openTemp,
		condensateRate: condensateRate,
		steamLoss:      steamLoss,
		currentTemp:    openTemp, // start at threshold
		isOpen:         false,
	}, nil
}

// Update sets the current inlet temperature and updates the trap state.
// The trap opens when temperature is at or below openTemp (condensate present)
// and closes when temperature exceeds openTemp (steam present).
func (st *SteamTrap) Update(inletTemp float64) {
	st.currentTemp = inletTemp
	if st.failed {
		st.isOpen = st.failedOpen
		return
	}
	st.isOpen = inletTemp <= st.openTemp
}

// IsOpen returns true if the trap is currently discharging condensate.
func (st *SteamTrap) IsOpen() bool {
	return st.isOpen
}

// DischargeRate returns the current condensate discharge rate (kg/s).
// When the trap is open, the rate is proportional to how far below
// the open temperature the inlet is. When closed, discharge is zero.
func (st *SteamTrap) DischargeRate() float64 {
	if !st.isOpen {
		return 0
	}
	if st.currentTemp >= st.openTemp {
		return 0
	}
	// Scale discharge rate by temperature differential, clamped to [0, 1]
	delta := st.openTemp - st.currentTemp
	factor := math.Min(delta/st.openTemp, 1.0)
	return st.condensateRate * factor
}

// SteamLossRate returns the rate of live steam escaping through the trap (kg/s).
// A properly functioning trap has minimal steam loss; a failed-open trap
// loses steam at the full condensate rate scaled by the steamLoss fraction.
func (st *SteamTrap) SteamLossRate() float64 {
	if !st.isOpen {
		return 0
	}
	return st.condensateRate * st.steamLoss
}

// Fail simulates a trap failure. If failOpen is true, the trap is stuck open
// (blowing live steam); if false, the trap is stuck closed (waterlogged system).
func (st *SteamTrap) Fail(failOpen bool) {
	st.failed = true
	st.failedOpen = failOpen
	st.isOpen = failOpen
}

// Reset clears any failure condition and returns the trap to normal operation.
func (st *SteamTrap) Reset() {
	st.failed = false
	st.isOpen = st.currentTemp <= st.openTemp
}

// IsFailed returns true if the trap is in a failed state.
func (st *SteamTrap) IsFailed() bool {
	return st.failed
}

// String returns a human-readable description of the steam trap state.
func (st *SteamTrap) String() string {
	status := "closed"
	if st.isOpen {
		status = "open"
	}
	failStatus := ""
	if st.failed {
		if st.failedOpen {
			failStatus = " [FAILED OPEN]"
		} else {
			failStatus = " [FAILED CLOSED]"
		}
	}
	return fmt.Sprintf("SteamTrap{status: %s, inletTemp: %.1f°C, openTemp: %.1f°C, discharge: %.3f kg/s%s}",
		status, st.currentTemp, st.openTemp, st.DischargeRate(), failStatus)
}
