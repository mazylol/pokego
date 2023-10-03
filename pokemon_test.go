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

// TestGetGrowthRate tests the GetGrowthRate function.
func TestGetGrowthRate(t *testing.T) {
	_, err := GetGrowthRate("slow")
	if err != nil {
		t.Fatal("Failed to get growth rate by name")
	}
}

// TestGetNature tests the GetNature function.
func TestGetNature(t *testing.T) {
	_, err := GetNature("adamant")
	if err != nil {
		t.Fatal("Failed to get nature by name")
	}
}

// TestGetPokeathlonStat tests the GetPokeathlonStat function.
func TestGetPokeathlonStat(t *testing.T) {
	_, err := GetPokeathlonStat("speed")
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

// TestGetLocationArea tests the GetLocationArea function.
func TestGetLocationArea(t *testing.T) {
	_, err := GetLocationArea("pikachu")
	if err != nil {
		t.Fatal("Failed to get pokemon location area by name")
	}
}

// TestGetColor tests the GetColor function.
func TestGetColor(t *testing.T) {
	_, err := GetColor("red")
	if err != nil {
		t.Fatal("Failed to get pokemon color by name")
	}
}

// TestGetForm tests the GetForm function.
func TestGetForm(t *testing.T) {
	_, err := GetForm("arceus-bug")
	if err != nil {
		t.Fatal("Failed to get pokemon form by name")
	}
}

// TestGetHabitat tests the GetHabitat function.
func TestGetHabitat(t *testing.T) {
	_, err := GetHabitat("cave")
	if err != nil {
		t.Fatal("Failed to get pokemon habitat by name")
	}
}

// TestGetShape tests the GetShape function.
func TestGetShape(t *testing.T) {
	_, err := GetShape("ball")
	if err != nil {
		t.Fatal("Failed to get pokemon shape by name")
	}
}

// TestGetPokemonSpecies tests the GetPokemonSpecies function.
func TestGetPokemonSpecies(t *testing.T) {
	_, err := GetSpecies("slugma")
	if err != nil {
		t.Fatal("Failed to get pokemon species by name")
	}
}

// TestGetStat tests the GetStat function.
func TestGetStat(t *testing.T) {
	_, err := GetStat("speed")
	if err != nil {
		t.Fatal("Failed to get stat by name")
	}
}

// TestGetType tests the GetType function.
func TestGetType(t *testing.T) {
	_, err := GetType("fire")
	if err != nil {
		t.Fatal("Failed to get type by name")
	}
}
