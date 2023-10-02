package pokego

import (
	"testing"
)

// TestGetPokemon tests the GetPokemon function.
func TestGetPokemon(t *testing.T) {
	_, err := GetPokemon("slugma")
	if err != nil {
		t.Fatal("Failed to get pokemon by name")
	}
}

// TestGetAbility tests the GetAbility function.
func TestGetAbility(t *testing.T) {
	_, err := GetAbility("stench")
	if err != nil {
		t.Fatal("Failed to get ability by name")
	}
}
