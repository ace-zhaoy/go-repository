package repository

import (
	"reflect"
	"testing"
)

func TestDict_KeyExists(t *testing.T) {
	dict := NewDict(map[string]int{
		"one": 1,
		"two": 2,
	})

	exists := dict.KeyExists("one")
	if !exists {
		t.Errorf("Expected key 'one' to exist")
	}

	exists = dict.KeyExists("three")
	if exists {
		t.Errorf("Expected key 'three' not to exist")
	}
}

func TestDict_Values(t *testing.T) {
	dict := NewDict(map[string]int{
		"one": 1,
		"two": 2,
	})

	values := dict.Values()
	expectedValues := []int{1, 2}

	if !reflect.DeepEqual(values, expectedValues) && !reflect.DeepEqual(values, []int{2, 1}) {
		t.Errorf("Expected values to be %v or %v, got %v", expectedValues, []int{2, 1}, values)
	}
}

func TestDict_Set(t *testing.T) {
	dict := NewDict(map[string]int{})

	dict.Set("three", 3)

	if val, ok := dict.m["three"]; !ok || val != 3 {
		t.Errorf("Expected 'three' to be set to 3, got %d", val)
	}
}
