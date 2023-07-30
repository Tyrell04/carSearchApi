package helper

import (
	"testing"
)

func TestBinaryFindString_TargetExists(t *testing.T) {
	arr := []string{"ADLERWERKE", "ARDIE-WERK", "ROSSKNECHT", "WUERZ", "BAYERMOTWERKEBMW"}
	target := "BAYERMOTWERKEBMW"

	found := BinaryFindString(arr, target)

	if !found {
		t.Errorf("Expected target %s to be found, but it wasn't", target)
	}
}

func TestBinaryFindString_TargetDoesNotExist(t *testing.T) {
	arr := []string{"apple", "banana", "cherry", "grape", "orange", "pear"}
	target := "watermelon"

	found := BinaryFindString(arr, target)

	if found {
		t.Errorf("Expected target %s not to be found, but it was", target)
	}
}

func TestBinaryFindString_EmptyArray(t *testing.T) {
	arr := []string{}
	target := "banana"

	found := BinaryFindString(arr, target)

	if found {
		t.Errorf("Expected target %s not to be found in an empty array, but it was", target)
	}
}

func TestBinaryFindString_SingleElementArray_TargetExists(t *testing.T) {
	arr := []string{"apple"}
	target := "apple"

	found := BinaryFindString(arr, target)

	if !found {
		t.Errorf("Expected target %s to be found in a single-element array, but it wasn't", target)
	}
}

func TestBinaryFindString_SingleElementArray_TargetDoesNotExist(t *testing.T) {
	arr := []string{"apple"}
	target := "orange"

	found := BinaryFindString(arr, target)

	if found {
		t.Errorf("Expected target %s not to be found in a single-element array, but it was", target)
	}
}
