package gds

import (
	"errors"
)

// List is doubly linked list
type List struct {
	Head *Item
	Tail *Item
}

// Item is node of List.
type Item struct {
	Data interface{}
	Prev *Item
	Next *Item
}

func newItem(x interface{}) *Item {
	return &Item{
		Data: x,
	}
}

// ErrEmpty indicates that a List is empty.
var ErrEmpty = errors.New("empty List")

// NewList creates a new List instance.
func NewList() *List {
	return &List{}
}

// IsEmpty returns true if the List is empty, otherwise false.
func (s *List) IsEmpty() bool {
	return s.Head == nil
}

func (s *List) addFirstItem(i *Item) bool {
	if s.IsEmpty() {
		s.Head = i
		s.Tail = i
		return true
	}
	return false
}

// PushHead pushes a new item at the head of the List.
func (s *List) PushHead(x interface{}) {
	i := newItem(x)
	if !s.addFirstItem(i) {
		s.Head.Prev = i
		i.Next = s.Head
		s.Head = i
	}
}

// PushTail pushes a new item at the head of the List.
func (s *List) PushTail(x interface{}) {
	i := newItem(x)
	if !s.addFirstItem(i) {
		s.Tail.Next = i
		i.Prev = s.Tail
		s.Tail = i
	}
}

func (s *List) isLastOne() bool {
	if !s.IsEmpty() {
		return s.Head.Prev == nil && s.Head.Next == nil
	}
	return false
}

func (s *List) onEmpty() {
	s.Head = nil
	s.Tail = nil
}

// PopHead remove a item from the head of the List and returns its data.
func (s *List) PopHead() (interface{}, error) {
	if s.IsEmpty() {
		return 0, ErrEmpty
	}

	v := s.Head.Data
	if s.isLastOne() {
		s.onEmpty()
	} else {
		s.Head = s.Head.Next
		s.Head.Prev = nil
	}
	return v, nil
}

func (s *List) PopTail() (interface{}, error) {
	if s.IsEmpty() {
		return 0, ErrEmpty
	}

	v := s.Tail.Data
	if s.isLastOne() {
		s.onEmpty()
	} else {
		s.Tail = s.Tail.Prev
		s.Tail.Next = nil
	}
	return v, nil
}

// GetHead gets the data of the item at the head of the List.
func (s *List) GetHead() (interface{}, error) {
	if s.IsEmpty() {
		return 0, ErrEmpty
	}

	return s.Head.Data, nil
}

func (s *List) GetTail() (interface{}, error) {
	if s.IsEmpty() {
		return 0, ErrEmpty
	}

	return s.Tail.Data, nil
}

// FindItem searches an item with value x.
func (s *List) FindItem(x interface{}) *Item {
	for i := s.Head; i != nil; i = i.Next {
		if i.Data == x {
			return i
		}
	}
	return nil
}
