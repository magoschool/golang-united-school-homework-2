package solution

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	testValue := CalcSquare(2.0, SidesCircle)
	expectedValue := math.Pi * 4

	if testValue != expectedValue {
		t.Errorf("CalcSquare failed: '%v' expected but '%v' found", expectedValue, testValue)
	}
}

func TestTriangle(t *testing.T) {
	testValue := CalcSquare(2.0, SidesTriangle)
	expectedValue := math.Sqrt(3)

	if testValue != expectedValue {
		t.Errorf("CalcSquare failed: '%v' expected but '%v' found", expectedValue, testValue)
	}
}

func TestSquare(t *testing.T) {
	testValue := CalcSquare(2.0, SidesSquare)
	expectedValue := 4.0

	if testValue != expectedValue {
		t.Errorf("CalcSquare failed: '%v' expected but '%v' found", expectedValue, testValue)
	}
}

func TestOther(t *testing.T) {
	testValue := CalcSquare(2.0, 1)
	expectedValue := 0.0

	if testValue != expectedValue {
		t.Errorf("CalcSquare failed: '%v' expected but '%v' found", expectedValue, testValue)
	}
}
