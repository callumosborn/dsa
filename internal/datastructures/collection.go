package datastructures

type Collection interface {
	Iterator() Interator
}

type Interator interface {
	HasNext() bool
	GetNext() any
}
