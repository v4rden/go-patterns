package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter := GetInstance()

	if counter == nil {
		t.Error("Expect instance, got nil")
	}

	counter2 := GetInstance()

	if counter != counter2 {
		t.Error("Expect same instance, got different")
	}
}
