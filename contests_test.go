package pokego

import "testing"

// TestGetContestType tests the GetContestType function.
func TestGetContestType(t *testing.T) {
	_, err := GetContestType("cool")
	if err != nil {
		t.Fatal("Failed to get contest type")
	}
}

// TestGetContestEffect tests the GetContestEffect function.
func TestGetContestEffect(t *testing.T) {
	_, err := GetContestEffect(1)
	if err != nil {
		t.Fatal("Failed to get contest effect")
	}
}

// TestGetSuperContestEffect tests the GetSuperContestEffect function.
func TestGetSuperContestEffect(t *testing.T) {
	_, err := GetSuperContestEffect(1)
	if err != nil {
		t.Fatal("Failed to get super contest effect")
	}
}
