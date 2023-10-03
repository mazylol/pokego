package pokego

import "testing"

func TestGetResourceList(t *testing.T) {
	_, err := GetResourceList("pokemon", 10)
	if err != nil {
		t.Fatal("Failed to get resource list")
	}
}
