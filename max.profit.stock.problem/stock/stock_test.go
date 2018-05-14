package stock

import "testing"

func TestIncorrectInput(t *testing.T) {
	_, err := GetData("test")

	if err == nil {
		t.Logf("error: %s", err)
		t.Fail()
	}
}

func TestValidInput(t *testing.T) {
	prices, err := GetData("GOOGL")

	if err != nil {
		t.Fail()
	}

	if len(prices) < 2 {
		t.Fail()
	}
}
