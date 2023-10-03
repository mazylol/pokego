package pokego

import (
	"testing"
)

// TestGetAbility tests the GetPokemonAbility function.
func TestGetAbility(t *testing.T) {
	_, err := GetPokemonAbility("stench")
	if err != nil {
		t.Fatal("Failed to get ability by name")
	}
}

// TestGetCharacteristic tests the GetPokemonCharacteristic function.
func TestGetCharacteristic(t *testing.T) {
	_, err := GetPokemonCharacteristic(1)
	if err != nil {
		t.Fatal("Failed to get characteristic by id")
	}
}

// TestGetEggGroup tests the GetPokemonEggGroup function.
func TestGetEggGroup(t *testing.T) {
	_, err := GetPokemonEggGroup("monster")
	if err != nil {
		t.Fatal("Failed to get egg group by name")
	}
}

// TestGetGender tests the GetPokemonGender function.
func TestGetGender(t *testing.T) {
	_, err := GetPokemonGender("female")
	if err != nil {
		t.Fatal("Failed to get gender by name")
	}
}

// TestGetGrowthRate tests the GetPokemonGrowthRate function.
func TestGetGrowthRate(t *testing.T) {
	_, err := GetPokemonGrowthRate("slow")
	if err != nil {
		t.Fatal("Failed to get growth rate by name")
	}
}

// TestGetNature tests the GetPokemonNature function.
func TestGetNature(t *testing.T) {
	_, err := GetPokemonNature("adamant")
	if err != nil {
		t.Fatal("Failed to get nature by name")
	}
}

// TestGetPokeathlonStat tests the GetPokemonPokeathlonStats function.
func TestGetPokeathlonStat(t *testing.T) {
	_, err := GetPokemonPokeathlonStats("speed")
	if err != nil {
		t.Fatal("Failed to get pokeathlon stat by name")
	}
}

// TestGetPokemon tests the GetPokemon function.
func TestGetPokemon(t *testing.T) {
	_, err := GetPokemon("slugma")
	if err != nil {
		t.Fatal("Failed to get pokemon by name")
	}
}

// TestGetLocationArea tests the GetPokemonLocationArea function.
func TestGetLocationArea(t *testing.T) {
	_, err := GetPokemonLocationArea("pikachu")
	if err != nil {
		t.Fatal("Failed to get pokemon location area by name")
	}
}

// TestGetColor tests the GetPokemonColor function.
func TestGetColor(t *testing.T) {
	_, err := GetPokemonColor("red")
	if err != nil {
		t.Fatal("Failed to get pokemon color by name")
	}
}

// TestGetForm tests the GetPokemonForm function.
func TestGetForm(t *testing.T) {
	_, err := GetPokemonForm("arceus-bug")
	if err != nil {
		t.Fatal("Failed to get pokemon form by name")
	}
}

// TestGetHabitat tests the GetPokemonHabitat function.
func TestGetHabitat(t *testing.T) {
	_, err := GetPokemonHabitat("cave")
	if err != nil {
		t.Fatal("Failed to get pokemon habitat by name")
	}
}

// TestGetShape tests the GetPokemonShape function.
func TestGetShape(t *testing.T) {
	_, err := GetPokemonShape("ball")
	if err != nil {
		t.Fatal("Failed to get pokemon shape by name")
	}
}

// TestGetPokemonSpecies tests the GetPokemonSpecies function.
func TestGetPokemonSpecies(t *testing.T) {
	_, err := GetPokemonSpecies("slugma")
	if err != nil {
		t.Fatal("Failed to get pokemon species by name")
	}
}

// TestGetStat tests the GetPokemonStat function.
func TestGetStat(t *testing.T) {
	_, err := GetPokemonStat("speed")
	if err != nil {
		t.Fatal("Failed to get stat by name")
	}
}

// TestGetType tests the GetPokemonType function.
func TestGetType(t *testing.T) {
	_, err := GetPokemonType("fire")
	if err != nil {
		t.Fatal("Failed to get type by name")
	}
}
