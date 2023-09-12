package mascot_test

import (
	"testing"

	"example.com/go-demo-1/mascot"
)

// runs a test to determine if the best mascot is correct or not
func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
