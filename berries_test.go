package pokego

import "testing"

// TestGetBerry tests the GetBerry function.
func TestGetBerry(t *testing.T) {
	_, err := GetBerry("cheri")
	if err != nil {
		t.Errorf("GetBerry(\"cheri\") failed with error: %v", err)
	}
}

// TestGetBerryFirmness tests the GetBerryFirmness function.
func TestGetBerryFirmness(t *testing.T) {
	_, err := GetBerryFirmness("very-soft")
	if err != nil {
		t.Errorf("GetBerryFirmness(\"very-soft\") failed with error: %v", err)
	}
}

// TestGetBerryFlavor tests the GetBerryFlavor function.
func TestGetBerryFlavor(t *testing.T) {
	_, err := GetBerryFlavor("spicy")
	if err != nil {
		t.Errorf("GetBerryFlavor(\"spicy\") failed with error: %v", err)
	}
}
