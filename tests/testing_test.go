package main

import "testing"

func TestFunction(t *testing.T) {
	t.Log("Testing reverse string")
	s := "Bonjour"
	reversed := reverseString(s);

	if reversed != "ruojhnoB" {
		t.Errorf("Expected string to equal ruojnoB")
	}
}
