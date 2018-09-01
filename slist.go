package gds

import (
	"errors"
)

// Slist is singly linked list
type Slist struct {
	Head   *Item
}

// Item is node of Slist.
type Item struct {
	Data interface{}
	Next *Item
}

func newItem(x interface{}) *Item {
	return &Item{
		Data: x,
	}
}

// ErrEmpty indicates that a Slist is empty.
var ErrEmpty = errors.New("empty Slist")

// NewSlist creates a new Slist instance.
func NewSlist() *Slist {
	return &Slist{}
}

// IsEmpty returns true if the Slist is empty, otherwise false.
func (s *Slist) IsEmpty() bool {
	return s.Head == nil
}

// PushHead pushes a new item at the head of the Slist.
func (s *Slist) PushHead(x interface{}) {
	i := newItem(x)
	if s.IsEmpty() {
		s.Head = i
	} else {
		i.Next = s.Head
		s.Head = i
	}
}

// func (s *Slist) PushTail(x interface{}) {

// }

// PopHead remove a item from the head of the Slist and returns its data.
func (s *Slist) PopHead() (interface{}, error) {
	if s.IsEmpty() {
		return 0, ErrEmpty
	}

	v := s.Head.Data
	s.Head = s.Head.Next
	return v, nil
}

// func (s *Slist) PopTail() (interface{}, error) {

// }

// GetHead gets the data of the item at the head of the Slist.
func (s *Slist) GetHead() (interface{}, error) {
	if s.IsEmpty() {
		return 0, ErrEmpty
	}

	return s.Head.Data, nil
}

// func (s *Slist) GetTail() (interface{}, error) {

// }

// FindItem searches an item with value x.
func (s *Slist) FindItem(x interface{}) *Item {
	for i := s.Head; i != nil; i = i.Next {
		if i.Data == x {
			return i
		}
	}
	return nil
}
