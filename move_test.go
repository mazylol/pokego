package pokego

import "testing"

// TestGetMove tests the GetMove function.
func TestGetMove(t *testing.T) {
	_, err := GetMove("pound")
	if err != nil {
		t.Fatal("Failed to get move by name")
	}
}

// TestGetAilment tests the GetAilment function.
func TestGetAilment(t *testing.T) {
	_, err := GetAilment("paralysis")
	if err != nil {
		t.Fatal("Failed to get move ailment by name")
	}
}

// TestGetBattleStyle tests the GetBattleStyle function.
func TestGetBattleStyle(t *testing.T) {
	_, err := GetBattleStyle("attack")
	if err != nil {
		t.Fatal("Failed to get move battle style by name")
	}
}

// TestGetCategory tests the GetCategory function.
func TestGetCategory(t *testing.T) {
	_, err := GetCategory("damage")
	if err != nil {
		t.Fatal("Failed to get move category by name")
	}
}

// TestGetDamageClass tests the GetDamageClass function.
func TestGetDamageClass(t *testing.T) {
	_, err := GetDamageClass("physical")
	if err != nil {
		t.Fatal("Failed to get move damage class by name")
	}
}

// TestGetLearnMethod tests the GetLearnMethod function.
func TestGetLearnMethod(t *testing.T) {
	_, err := GetLearnMethod("level-up")
	if err != nil {
		t.Fatal("Failed to get move learn method by name")
	}
}

// TestGetTarget tests the GetTarget function.
func TestGetTarget(t *testing.T) {
	_, err := GetTarget("specific-move")
	if err != nil {
		t.Fatal("Failed to get move target by name")
	}
}
