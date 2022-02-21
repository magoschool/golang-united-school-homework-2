package solution

import (
	"math"
	"testing"
)

const errorMessage = "CalcSquare failed: '%v' expected but '%v' found"

func testCircle(t *testing.T) {
	testValue := CalcSquare(2.0, SidesCircle)
	expectedValue := math.Pi * 4

	if testValue != expectedValue {
		t.Errorf(errorMessage, expectedValue, testValue)
	}
}

func testTriangle(t *testing.T) {
	testValue := CalcSquare(2.0, SidesTriangle)
	expectedValue := math.Sqrt(3)

	if testValue != expectedValue {
		t.Errorf(errorMessage, expectedValue, testValue)
	}
}

func testSquare(t *testing.T) {
	testValue := CalcSquare(2.0, SidesSquare)
	expectedValue := 4.0

	if testValue != expectedValue {
		t.Errorf(errorMessage, expectedValue, testValue)
	}
}

func testOther(t *testing.T) {
	testValue := CalcSquare(2.0, 1)
	expectedValue := 0.0

	if testValue != expectedValue {
		t.Errorf(errorMessage, expectedValue, testValue)
	}
}

func TestCalcSquare(t *testing.T) {
	testCircle(t)
	testTriangle(t)
	testSquare(t)
	testOther(t)
}
