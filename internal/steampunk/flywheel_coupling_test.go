package steampunk

import (
	"testing"
)

func makeTestCoupling(t *testing.T) (*FlywheelCoupling, *Crankshaft, *Flywheel) {
	t.Helper()
	cs, err := NewCrankshaft(0.1, 0.05)
	if err != nil {
		t.Fatalf("crankshaft: %v", err)
	}
	fw, err := NewFlywheel(0.4, 15.0)
	if err != nil {
		t.Fatalf("flywheel: %v", err)
	}
	fc, err := NewFlywheelCoupling(cs, fw, 1.0)
	if err != nil {
		t.Fatalf("coupling: %v", err)
	}
	return fc, cs, fw
}

func TestNewFlywheelCoupling(t *testing.T) {
	fc, _, _ := makeTestCoupling(t)
	if fc.GearRatio() != 1.0 {
		t.Errorf("expected gear ratio 1.0, got %f", fc.GearRatio())
	}
}

func TestNewFlywheelCouplingNilCrankshaft(t *testing.T) {
	fw, _ := NewFlywheel(0.4, 15)
	_, err := NewFlywheelCoupling(nil, fw, 1.0)
	if err == nil {
		t.Error("expected error for nil crankshaft")
	}
}

func TestNewFlywheelCouplingNilFlywheel(t *testing.T) {
	cs, _ := NewCrankshaft(0.1, 0.05)
	_, err := NewFlywheelCoupling(cs, nil, 1.0)
	if err == nil {
		t.Error("expected error for nil flywheel")
	}
}

func TestNewFlywheelCouplingInvalidRatio(t *testing.T) {
	cs, _ := NewCrankshaft(0.1, 0.05)
	fw, _ := NewFlywheel(0.4, 15)
	_, err := NewFlywheelCoupling(cs, fw, 0)
	if err == nil {
		t.Error("expected error for zero gear ratio")
	}
	_, err = NewFlywheelCoupling(cs, fw, -2)
	if err == nil {
		t.Error("expected error for negative gear ratio")
	}
}

func TestFlywheelCouplingUpdateZeroDt(t *testing.T) {
	fc, _, fw := makeTestCoupling(t)
	fw.SetRPM(200)
	fc.Update(0)
	if fw.RPM() != 200 {
		t.Errorf("expected flywheel RPM unchanged at 200, got %f", fw.RPM())
	}
}

func TestFlywheelCouplingSyncRPM(t *testing.T) {
	fc, cs, fw := makeTestCoupling(t)
	_ = cs
	fc.SyncRPM()
	if fw.RPM() != fc.FlywheelRPM() {
		t.Errorf("expected flywheel RPM=%f after sync, got %f", fc.FlywheelRPM(), fw.RPM())
	}
}

func TestFlywheelCouplingGearRatio(t *testing.T) {
	cs, _ := NewCrankshaft(0.1, 0.05)
	fw, _ := NewFlywheel(0.4, 15)
	fc, err := NewFlywheelCoupling(cs, fw, 2.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if fc.GearRatio() != 2.5 {
		t.Errorf("expected ratio 2.5, got %f", fc.GearRatio())
	}
}

func TestFlywheelCouplingString(t *testing.T) {
	fc, _, _ := makeTestCoupling(t)
	if len(fc.String()) == 0 {
		t.Error("expected non-empty string")
	}
}
