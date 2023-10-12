package pokego

import "testing"

// TestGetContestType tests the GetContestType function.
func TestGetContestType(t *testing.T) {
	_, err := GetContestType("cool")
	if err != nil {
		t.Errorf("GetContestType(\"cool\") failed with error: %v", err)
	}
}

// TestGetContestEffect tests the GetContestEffect function.
func TestGetContestEffect(t *testing.T) {
	_, err := GetContestEffect(1)
	if err != nil {
		t.Errorf("GetContestEffect(1) failed with error: %v", err)
	}
}

// TestGetSuperContestEffect tests the GetSuperContestEffect function.
func TestGetSuperContestEffect(t *testing.T) {
	_, err := GetSuperContestEffect(1)
	if err != nil {
		t.Errorf("GetSuperContestEffect(1) failed with error: %v", err)
	}
}
