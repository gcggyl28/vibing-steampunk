package steampunk

import "testing"

func makeTestSafetyValve(t *testing.T) *SafetyValve {
	t.Helper()
	sv, err := NewSafetyValve("relief", 180, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return sv
}

func TestNewSafetyValve(t *testing.T) {
	sv := makeTestSafetyValve(t)
	if sv.IsOpen() {
		t.Error("new safety valve should be closed")
	}
}

func TestNewSafetyValveInvalidSetPoint(t *testing.T) {
	_, err := NewSafetyValve("bad", -10, 5)
	if err == nil {
		t.Error("expected error for non-positive setPoint")
	}
}

func TestNewSafetyValveInvalidBlowdown(t *testing.T) {
	_, err := NewSafetyValve("bad", 180, 200)
	if err == nil {
		t.Error("expected error for blowdown >= setPoint")
	}
}

func TestSafetyValveStaysClosed(t *testing.T) {
	sv := makeTestSafetyValve(t)
	vented := sv.Check(150)
	if vented != 0 {
		t.Errorf("expected 0 vented, got %.2f", vented)
	}
	if sv.IsOpen() {
		t.Error("valve should remain closed below setPoint")
	}
}

func TestSafetyValveOpensAtSetPoint(t *testing.T) {
	sv := makeTestSafetyValve(t)
	sv.Check(185)
	if !sv.IsOpen() {
		t.Error("valve should open at or above setPoint")
	}
}

func TestSafetyValveVentsExcess(t *testing.T) {
	sv := makeTestSafetyValve(t)
	vented := sv.Check(190)
	if vented != 10 { // 190 - 180 = 10
		t.Errorf("expected 10 vented, got %.2f", vented)
	}
}

func TestSafetyValveReseats(t *testing.T) {
	sv := makeTestSafetyValve(t)
	sv.Check(185)
	if !sv.IsOpen() {
		t.Fatal("valve should be open")
	}
	// pressure drops below setPoint - blowdown (180 - 10 = 170)
	sv.Check(165)
	if sv.IsOpen() {
		t.Error("valve should reseat below blowdown pressure")
	}
}

func TestSafetyValveTotalReleased(t *testing.T) {
	sv := makeTestSafetyValve(t)
	sv.Check(190) // +10
	sv.Check(195) // +15
	if sv.TotalReleased() != 25 {
		t.Errorf("expected 25 total released, got %.2f", sv.TotalReleased())
	}
}

func TestSafetyValveReset(t *testing.T) {
	sv := makeTestSafetyValve(t)
	sv.Check(190)
	sv.Reset()
	if sv.IsOpen() {
		t.Error("valve should be closed after reset")
	}
	if sv.TotalReleased() != 0 {
		t.Errorf("expected 0 released after reset, got %.2f", sv.TotalReleased())
	}
}

func TestSafetyValveString(t *testing.T) {
	sv := makeTestSafetyValve(t)
	if sv.String() == "" {
		t.Error("expected non-empty string")
	}
}
