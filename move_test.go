package pokego

import "testing"

// TestGetMove tests the GetMove function.
func TestGetMove(t *testing.T) {
	_, err := GetMove("pound")
	if err != nil {
		t.Error("Failed to get move by name")
	}
}

// TestGetMoveList tests the GetMoveList function.
func TestGetMoveList(t *testing.T) {
	_, err := GetMoveList(20)
	if err != nil {
		t.Error("Failed to get move list")
	}
}
