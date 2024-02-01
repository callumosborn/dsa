package datastructures

import "testing"

func TestIsEmpty(t *testing.T) {
	var stack Stack

	if stack.isEmpty() == false {
		t.Error("1")
	}

	stack.push(1)

	if stack.isEmpty() == true {
		t.Error("2")
	}

	stack.pop()

	if stack.isEmpty() == false {
		t.Error("3")
	}
}

func TestPush(t *testing.T) {
	var stack Stack
	stack.push(1)

	if len(stack) != 1 {
		t.Errorf("1")
	}

	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(5)

	if len(stack) != 5 {
		t.Error("2")
	}
}

func TestPop(t *testing.T) {
	var stack Stack

	val := stack.pop()

	if val != nil {
		t.Error("1")
	}

	stack.push(1)
	stack.push(2)
	stack.push(3)

	val1 := stack.pop()

	val2, ok := val1.(int)

	if !ok {
		t.Error("1")
	}

	if val2 != 3 {
		t.Error("1")
	}

	stack.pop()
	stack.pop()

	val = stack.pop()

	if val != nil {
		t.Error("4")
	}
}
