package datastructures

import (
	"testing"
)

func TestXxx(t *testing.T) {
	var queue Queue

	if queue.isEmpty() == false {
		t.Error("Stack should be empty, but it is not.")
	}

	queue.enqueue(1)

	if queue.isEmpty() == true {
		t.Error("Stack should not be empty, but is is.")
	}

	queue.dequeue()

	if queue.isEmpty() == false {
		t.Error("Stack should be empty, but it is not.")
	}
}

func TestEnqueue(t *testing.T) {
	var queue Queue
	queue.enqueue(1)

	expectedLength := 1

	if len(queue) != expectedLength {
		t.Errorf("Expected queue length %d, but got %d", expectedLength, len(queue))
	}
}

func TestDequeue(t *testing.T) {
	var queue Queue

	val1 := queue.dequeue()

	if val1 != nil {
		t.Errorf("Expected nil value, but got %v", val1)
	}

	expectedVal := 1

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)

	val2 := queue.dequeue()
	val3, _ := val2.(int)

	if val3 != expectedVal {
		t.Errorf("Expected val to be %d, but got %d", expectedVal, val3)
	}

	queue.dequeue()
	queue.dequeue()

	expectedLength := 0

	if len(queue) != expectedLength {
		t.Errorf("Expected stack length %d, but got %d", expectedLength, len(queue))
	}
}
