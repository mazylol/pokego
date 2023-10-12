package pokego

import "testing"

// TestGetEncounterMethod tests the GetEncounterMethod function.
func TestGetEncounterMethod(t *testing.T) {
	_, err := GetEncounterMethod("walk")
	if err != nil {
		t.Errorf("GetEncounterMethod(\"walk\") failed with error: %v", err)
	}
}

// TestGetEncounterCondition tests the GetEncounterCondition function.
func TestGetEncounterCondition(t *testing.T) {
	_, err := GetEncounterCondition("swarm")
	if err != nil {
		t.Errorf("GetEncounterCondition(\"swarm\") failed with error: %v", err)
	}
}

// TestGetEncounterConditionValue tests the GetEncounterConditionValue function.
func TestGetEncounterConditionValue(t *testing.T) {
	_, err := GetEncounterConditionValue("swarm-yes")
	if err != nil {
		t.Errorf("GetEncounterConditionValue(\"swarm-yes\") failed with error: %v", err)
	}
}
