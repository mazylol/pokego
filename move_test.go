package pokego

import "testing"

// TestGetMove tests the GetMove function.
func TestGetMove(t *testing.T) {
	_, err := GetMove("pound")
	if err != nil {
		t.Errorf("GetMove(\"pound\") failed with error: %v", err)
	}
}

// TestGetAilment tests the GetAilment function.
func TestGetAilment(t *testing.T) {
	_, err := GetAilment("paralysis")
	if err != nil {
		t.Errorf("GetAilment(\"paralysis\") failed with error: %v", err)
	}
}

// TestGetBattleStyle tests the GetBattleStyle function.
func TestGetBattleStyle(t *testing.T) {
	_, err := GetBattleStyle("attack")
	if err != nil {
		t.Errorf("GetBattleStyle(\"attack\") failed with error: %v", err)
	}
}

// TestGetCategory tests the GetCategory function.
func TestGetCategory(t *testing.T) {
	_, err := GetCategory("damage")
	if err != nil {
		t.Errorf("GetCategory(\"damage\") failed with error: %v", err)
	}
}

// TestGetDamageClass tests the GetDamageClass function.
func TestGetDamageClass(t *testing.T) {
	_, err := GetDamageClass("physical")
	if err != nil {
		t.Errorf("GetDamageClass(\"physical\") failed with error: %v", err)
	}
}

// TestGetLearnMethod tests the GetLearnMethod function.
func TestGetLearnMethod(t *testing.T) {
	_, err := GetLearnMethod("level-up")
	if err != nil {
		t.Errorf("GetLearnMethod(\"level-up\") failed with error: %v", err)
	}
}

// TestGetTarget tests the GetTarget function.
func TestGetTarget(t *testing.T) {
	_, err := GetTarget("specific-move")
	if err != nil {
		t.Errorf("GetTarget(\"specific-move\") failed with error: %v", err)
	}
}
