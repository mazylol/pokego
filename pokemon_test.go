package pokego

import (
	"testing"
)

// TestGetAbility tests the GetAbility function.
func TestGetAbility(t *testing.T) {
	_, err := GetAbility("stench")
	if err != nil {
		t.Fatal("Failed to get ability by name")
	}
}

// TestGetCharacteristic tests the GetCharacteristic function.
func TestGetCharacteristic(t *testing.T) {
	_, err := GetCharacteristic(1)
	if err != nil {
		t.Fatal("Failed to get characteristic by id")
	}
}

// TestGetEggGroup tests the GetEggGroup function.
func TestGetEggGroup(t *testing.T) {
	_, err := GetEggGroup("monster")
	if err != nil {
		t.Fatal("Failed to get egg group by name")
	}
}

// TestGetGender tests the GetGender function.
func TestGetGender(t *testing.T) {
	_, err := GetGender("female")
	if err != nil {
		t.Fatal("Failed to get gender by name")
	}
}

// TestGetPokemon tests the GetPokemon function.
func TestGetPokemon(t *testing.T) {
	_, err := GetPokemon("slugma")
	if err != nil {
		t.Fatal("Failed to get pokemon by name")
	}
}
