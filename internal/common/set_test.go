package common

import (
	"testing"
)

func TestNewSetWithValues(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	if s.IsEmpty() || s.Len() != 3 {
		t.Errorf("New set should contain the initial values")
	}
}

func TestAddMultipleValuesToSet(t *testing.T) {
	s := NewSet[int]()
	s = s.Add(1, 2, 3)
	if s.IsEmpty() || s.Len() != 3 {
		t.Errorf("Set should contain all the added elements")
	}
}

func TestAddDuplicateValuesToSet(t *testing.T) {
	s := NewSet[int]()
	s = s.Add(1, 1, 1)
	if s.Len() != 1 {
		t.Errorf("Set should not contain duplicate elements")
	}
}

func TestRemoveNonExistentValueFromSet(t *testing.T) {
	s := NewSet[int]()
	s = s.Remove(1)
	if !s.IsEmpty() {
		t.Errorf("Set should remain empty when removing a non-existent element")
	}
}

func TestRemoveMultipleValuesFromSet(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	s = s.Remove(1, 2)
	if s.Len() != 1 || s.Contains(1) || s.Contains(2) {
		t.Errorf("Set should not contain the removed elements")
	}
}

func TestSetDoesNotContainRemovedElement(t *testing.T) {
	s := NewSet[int](1)
	s = s.Remove(1)
	if s.Contains(1) {
		t.Errorf("Set should not contain a removed element")
	}
}

func TestSetContainsMultipleElements(t *testing.T) {
	s := NewSet[int](1, 2, 3)
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("Set should contain all the added elements")
	}
}

func TestSetLengthWithDuplicateValues(t *testing.T) {
	s := NewSet[int](1, 1, 1)
	if s.Len() != 1 {
		t.Errorf("Set length should not count duplicate values")
	}
}

func TestSetIsEmptyAfterRemovingAllElements(t *testing.T) {
	s := NewSet[int](1)
	s = s.Remove(1)
	if !s.IsEmpty() {
		t.Errorf("Set should be empty after all elements are removed")
	}
}

func TestSetIsNotEmptyWithValues(t *testing.T) {
	s := NewSet[int](1)
	if s.IsEmpty() {
		t.Errorf("Set should not be empty when it contains elements")
	}
}
