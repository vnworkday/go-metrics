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

func TestSetEquals(t *testing.T) {
	tests := []struct {
		name string
		s1   Set[int]
		s2   Set[int]
		want bool
	}{
		{
			"EqualSets",
			NewSet[int](1, 2, 3),
			NewSet[int](1, 2, 3),
			true,
		},
		{
			"DifferentSets",
			NewSet[int](1, 2, 3),
			NewSet[int](4, 5, 6),
			false,
		},
		{
			"EmptySetAndNonEmptySet",
			NewSet[int](),
			NewSet[int](1, 2, 3),
			false,
		},
		{
			"TwoEmptySets",
			NewSet[int](),
			NewSet[int](),
			true,
		},
		{
			"SetWithExtraElements",
			NewSet[int](1, 2, 3),
			NewSet[int](1, 2, 3, 4),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s1.Equals(tt.s2); got != tt.want {
				t.Errorf("Set.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToSet(t *testing.T) {
	tests := []struct {
		name string
		str  string
		sep  string
		want Set[string]
	}{
		{
			"EmptyString",
			"",
			",",
			NewSet[string](),
		},
		{
			"SingleElement",
			"element",
			",",
			NewSet[string]("element"),
		},
		{
			"MultipleElements",
			"element1,element2,element3",
			",",
			NewSet[string]("element1", "element2", "element3"),
		},
		{
			"MultipleElementsWithSpaces",
			"element1, element2, element3",
			",",
			NewSet[string]("element1", " element2", " element3"),
		},
		{
			"DifferentSeparator",
			"element1;element2;element3",
			";",
			NewSet[string]("element1", "element2", "element3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToSet(tt.str, tt.sep); !got.Equals(tt.want) {
				t.Errorf("StringToSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
