package gds

import (
	"testing"
	"fmt"
)

func TestList(t *testing.T) {
	{
		s := NewList()
		for i := 0; i < 100; i++ {
			s.PushHead(i)
		}
		if i := s.FindItem(51); i == nil {
			t.Error("s.FindItem(51) got nil")
		}
		if i := s.FindItem(511); i != nil {
			t.Error("s.FindItem(511) got value")
		}

		var a interface{} = 1
		var b interface{} = "abc"
		if a != b {
			fmt.Println("a != b")
		}

		var c interface{} = 1
		var d interface{} = 1
		if c == d {
			fmt.Println("c == d")
		}

		var e interface{} = 1
		var f interface{} = 2
		if e != f {
			fmt.Println("e != f")
		}
	}

	{
		s := NewList()

		if !s.IsEmpty() {
			t.Error("should be empty")
		}

		s.PushHead(2)
		s.PushHead(1)
		s.PushTail(3)
		s.PushTail(4)

		v, _ := s.GetHead()
		if v != 1 {
			t.Error("shold be 1")
		}

		v, _ = s.GetTail()
		if v != 4 {
			t.Error("should be 4")
		}

		s.PopTail()
		v, _ = s.PopTail()
		if v != 3 {
			t.Error("should be 3")
		}
		s.PopHead()
		v, _ = s.PopTail()
		if v != 2 {
			t.Error("should be 2")
		}

		if !s.IsEmpty() {
			t.Error("should be empty now")
		}
	}

	s := NewList()
	if !s.IsEmpty() {
		t.Error("new slist should be empty")
	}

	s.PushHead(1)
	v, err := s.GetHead()
	if err != nil {
		t.Error("should not got error")
	}
	v2, ok := v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 1 {
		t.Errorf("s.GetHead() got %d, want %d", v2, 1)
	}

	s.PushHead(2)
	v, err = s.GetHead()
	if err != nil {
		t.Error("should not got error")
	}
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 2 {
		t.Errorf("s.GetHead() got %d, want %d", v2, 2)
	}

	s.PushHead(3)
	v, err = s.GetHead()
	if err != nil {
		t.Error("should not got error")
	}
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 3 {
		t.Errorf("s.GetHead() got %d, want %d", v2, 3)
	}

	v, err = s.PopHead()
	if err != nil {
		t.Error("should not got error")
	}
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 3 {
		t.Errorf("s.PopHead() got %d, want %d", v2, 3)
	}

	v, err = s.PopHead()
	if err != nil {
		t.Error("should not got error")
	}
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 2 {
		t.Errorf("s.PopHead() got %d, want %d", v2, 2)
	}

	v, err = s.PopHead()
	if err != nil {
		t.Error("should not got error")
	}
	v2, ok = v.(int)
	if !ok {
		t.Error("wrong type")
	}
	if v2 != 1 {
		t.Errorf("s.PopHead() got %d, want %d", v2, 1)
	}

	if _, err := s.PopHead(); err == nil {
		t.Error("should return error")
	}
	if _, err := s.GetHead(); err == nil {
		t.Error("should return error")
	}
}

func BenchmarkListPushHead(b *testing.B) {
	s := NewList()

	for i := 0; i < b.N; i++ {
		s.PushHead(i)
	}
}

func BenchmarkListGetHead(b *testing.B) {
	s := NewList()

	s.PushHead(1)

	for i := 0; i < b.N; i++ {
		s.GetHead()
	}
}

func BenchmarkListPushHeadAndPopHead(b *testing.B) {
	s := NewList()

	for i := 0; i < b.N; i++ {
		s.PushHead(i)
	}
	for i := 0; i < b.N; i++ {
		s.PopHead()
	}
}

func BenchmarkListPushTail(b *testing.B) {
	s := NewList()

	for i := 0; i < b.N; i++ {
		s.PushTail(i)
	}
}

func BenchmarkListGetTail(b *testing.B) {
	s := NewList()

	s.PushTail(1)

	for i := 0; i < b.N; i++ {
		s.GetTail()
	}
}

func BenchmarkListPushTailAndPopTail(b *testing.B) {
	s := NewList()

	for i := 0; i < b.N; i++ {
		s.PushTail(i)
	}
	for i := 0; i < b.N; i++ {
		s.PopTail()
	}
}
