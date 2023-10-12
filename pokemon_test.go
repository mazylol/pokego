package pokego

import (
	"testing"
)

// TestGetAbility tests the GetPokemonAbility function.
func TestGetAbility(t *testing.T) {
	_, err := GetPokemonAbility("stench")
	if err != nil {
		t.Errorf("GetPokemonAbility(\"stench\") failed with error: %v", err)
	}
}

// TestGetCharacteristic tests the GetPokemonCharacteristic function.
func TestGetCharacteristic(t *testing.T) {
	_, err := GetPokemonCharacteristic(1)
	if err != nil {
		t.Errorf("GetPokemonCharacteristic(1) failed with error: %v", err)
	}
}

// TestGetEggGroup tests the GetPokemonEggGroup function.
func TestGetEggGroup(t *testing.T) {
	_, err := GetPokemonEggGroup("monster")
	if err != nil {
		t.Errorf("GetPokemonEggGroup(\"monster\") failed with error: %v", err)
	}
}

// TestGetGender tests the GetPokemonGender function.
func TestGetGender(t *testing.T) {
	_, err := GetPokemonGender("female")
	if err != nil {
		t.Errorf("GetPokemonGender(\"female\") failed with error: %v", err)
	}
}

// TestGetGrowthRate tests the GetPokemonGrowthRate function.
func TestGetGrowthRate(t *testing.T) {
	_, err := GetPokemonGrowthRate("slow")
	if err != nil {
		t.Errorf("GetPokemonGrowthRate(\"slow\") failed with error: %v", err)
	}
}

// TestGetNature tests the GetPokemonNature function.
func TestGetNature(t *testing.T) {
	_, err := GetPokemonNature("adamant")
	if err != nil {
		t.Errorf("GetPokemonNature(\"adamant\") failed with error: %v", err)
	}
}

// TestGetPokeathlonStat tests the GetPokemonPokeathlonStats function.
func TestGetPokeathlonStat(t *testing.T) {
	_, err := GetPokemonPokeathlonStats("speed")
	if err != nil {
		t.Errorf("GetPokemonPokeathlonStats(\"speed\") failed with error: %v", err)
	}
}

// TestGetPokemon tests the GetPokemon function.
func TestGetPokemon(t *testing.T) {
	_, err := GetPokemon("slugma")
	if err != nil {
		t.Errorf("GetPokemon(\"slugma\") failed with error: %v", err)
	}
}

// TestGetLocationArea tests the GetPokemonLocationArea function.
func TestGetLocationArea(t *testing.T) {
	_, err := GetPokemonLocationArea("pikachu")
	if err != nil {
		t.Errorf("GetPokemonLocationArea(\"pikachu\") failed with error: %v", err)
	}
}

// TestGetColor tests the GetPokemonColor function.
func TestGetColor(t *testing.T) {
	_, err := GetPokemonColor("red")
	if err != nil {
		t.Errorf("GetPokemonColor(\"red\") failed with error: %v", err)
	}
}

// TestGetForm tests the GetPokemonForm function.
func TestGetForm(t *testing.T) {
	_, err := GetPokemonForm("arceus-bug")
	if err != nil {
		t.Errorf("GetPokemonForm(\"arceus-bug\") failed with error: %v", err)
	}
}

// TestGetHabitat tests the GetPokemonHabitat function.
func TestGetHabitat(t *testing.T) {
	_, err := GetPokemonHabitat("cave")
	if err != nil {
		t.Errorf("GetPokemonHabitat(\"cave\") failed with error: %v", err)
	}
}

// TestGetShape tests the GetPokemonShape function.
func TestGetShape(t *testing.T) {
	_, err := GetPokemonShape("ball")
	if err != nil {
		t.Errorf("GetPokemonShape(\"ball\") failed with error: %v", err)
	}
}

// TestGetPokemonSpecies tests the GetPokemonSpecies function.
func TestGetPokemonSpecies(t *testing.T) {
	_, err := GetPokemonSpecies("slugma")
	if err != nil {
		t.Errorf("GetPokemonSpecies(\"slugma\") failed with error: %v", err)
	}
}

// TestGetStat tests the GetPokemonStat function.
func TestGetStat(t *testing.T) {
	_, err := GetPokemonStat("speed")
	if err != nil {
		t.Errorf("GetPokemonStat(\"speed\") failed with error: %v", err)
	}
}

// TestGetType tests the GetPokemonType function.
func TestGetType(t *testing.T) {
	_, err := GetPokemonType("fire")
	if err != nil {
		t.Errorf("GetPokemonType(\"fire\") failed with error: %v", err)
	}
}
