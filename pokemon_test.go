package pokego

import (
	"testing"
)

// TestGetPokemon tests the GetPokemon function.
func TestGetPokemon(t *testing.T) {
	_, err := GetPokemon("charizard")
	if err != nil {
		t.Error("Failed to get pokemon by name")
	}

	_, err = GetPokemon(6)
	if err != nil {
		t.Error("Failed to get pokemon by id")
	}
}
