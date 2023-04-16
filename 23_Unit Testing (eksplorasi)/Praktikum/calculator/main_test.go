package main

import "testing"

func Test Add (t *testing.T) {
	result := Add(2,3)
	expected := float64(5)
	if result != expected {
		t.Errorf("Addition test failed: expected %v but got %v", expected, result)
	}
}

func TestSubstract(t *testing.T) {
	result := Subtract(5,3)
	expected := float64(2)
	if result != expected {
		t.Error("Substraction test failed: expected %v but got %v", expected result)
	}
}

func TestMultiply(t *testing.T) {
	result := Multiply(2,3)
	expected := float64(6)
	if result != expected {
		t.Errorf("Multiplication test failed: expected %v but got %v", expected result)
	}
}

func TestDivide (t *testing.T) {
	result, err := Divide(6,3)
	if err != nil {
		t.Errorf("Division test failed: unexpected error %v", err)
	}
	expected := float64(2)
	if result != expected {
		t.Errorf("Division test failed: expected %v but got %v", expected, result)
	}

	_, err = Divide(6,0)
	if err = nil {
		t. Errorf("Division test failed: expected error but got nil")
	}
}