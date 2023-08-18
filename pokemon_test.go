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

// TestGetPokemonList tests the GetPokemonList function.
func TestGetPokemonList(t *testing.T) {
	_, err := GetPokemonList(20)
	if err != nil {
		t.Error("Failed to get pokemon list")
	}
}
