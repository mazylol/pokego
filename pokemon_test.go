package pokego

import (
	"testing"
)

// TestGetPokemon tests the GetPokemon function.
func TestGetPokemon(t *testing.T) {
	_, err := GetPokemon("slugma")
	if err != nil {
		t.Error("Failed to get pokemon by name")
	}
}
