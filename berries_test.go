package pokego

import "testing"

// TestGetBerry tests the GetBerry function.
func TestGetBerry(t *testing.T) {
	_, err := GetBerry("cheri")
	if err != nil {
		t.Fatal("Failed to get berry by name")
	}
}

// TestGetBerryFirmness tests the GetBerryFirmness function.
func TestGetBerryFirmness(t *testing.T) {
	_, err := GetBerryFirmness("very-soft")
	if err != nil {
		t.Fatal("Failed to get berry firmness by name")
	}
}

// TestGetBerryFlavor tests the GetBerryFlavor function.
func TestGetBerryFlavor(t *testing.T) {
	_, err := GetBerryFlavor("spicy")
	if err != nil {
		t.Fatal("Failed to get berry flavor by name")
	}
}