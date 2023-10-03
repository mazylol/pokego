package pokego

import "testing"

// TestGetMove tests the GetMove function.
func TestGetMove(t *testing.T) {
	_, err := GetMove("pound")
	if err != nil {
		t.Fatal("Failed to get move by name")
	}
}
