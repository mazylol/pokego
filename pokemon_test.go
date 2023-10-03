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

// TestGetPokemonList tests the GetPokemonList function.
func TestGetPokemonList(t *testing.T) {
	_, err := GetPokemonList(20)
	if err != nil {
		t.Fatal("Failed to get pokemon list")
	}
}

// TestGetAbility tests the GetAbility function.
func TestGetAbility(t *testing.T) {
	_, err := GetAbility("stench")
	if err != nil {
		t.Fatal("Failed to get ability by name")
	}
}

// TestGetAbilityList tests the GetAbilityList function.
func TestGetAbilityList(t *testing.T) {
	_, err := GetAbilityList(20)
	if err != nil {
		t.Fatal("Failed to get ability list")
	}
}
