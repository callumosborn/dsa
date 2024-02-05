package datastructures

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	var stack Stack

	if stack.isEmpty() == false {
		t.Error("Stack should be empty, but it is not.")
	}

	stack.push(1)

	if stack.isEmpty() == true {
		t.Error("Stack should not be empty, but is is.")
	}

	stack.pop()

	if stack.isEmpty() == false {
		t.Error("Stack should be empty, but it is not.")
	}
}

func TestPush(t *testing.T) {
	var stack Stack
	stack.push(1)

	expectedLength := 1

	if len(stack) != expectedLength {
		t.Errorf("Expected stack length %d, but got %d", expectedLength, len(stack))
	}

	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)

	expectedLength = 5

	if len(stack) != expectedLength {
		t.Errorf("Expected stack length %d, but got %d", expectedLength, len(stack))
	}
}

func TestPop(t *testing.T) {
	var stack Stack

	val := stack.pop()

	if val != nil {
		t.Errorf("Expected nil value, but got %v", val)
	}

	stack.push(1)
	stack.push(2)
	stack.push(3)

	val1 := stack.pop()

	val2, _ := val1.(int)

	expectedVal := 3

	if val2 != expectedVal {
		t.Errorf("Expected val to be %d, but got %d", expectedVal, val2)
	}

	stack.pop()
	stack.pop()

	val = stack.pop()

	if val != nil {
		t.Errorf("Expected val to be nil, but got %d", val)
	}
}

func TestIterator(t *testing.T) {
	var stack Stack

	iterator := stack.Iterator()

	if iterator.HasNext() {
		t.Errorf("Expected HasNext to return false, but got true")
	}

	if iterator.GetNext() != nil {
		t.Errorf("Expected GetNext to return nil, but got a value")
	}

	expectedVal := 5

	stack.push(expectedVal)

	iterator = stack.Iterator()

	if !iterator.HasNext() {
		t.Errorf("Expected HasNext to return true, but got false")
	}

	val := iterator.GetNext().(int)

	if val != expectedVal {
		t.Errorf("Expected GetNext to return %d, but got %d", expectedVal, val)
	}
}
