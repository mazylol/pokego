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
