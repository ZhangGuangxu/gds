package gds

import (
	"testing"
	"fmt"
)

func TestSlist(t *testing.T) {
	{
		s := NewSlist()
		for i := 0; i < 100; i++ {
			s.PushHead(i)
		}
		if i := s.FindItem(51); i == nil {
			t.Error("s.FindItem(51) got nil")
		}
		if i := s.FindItem(511); i != nil {
			t.Error("s.FindItem(51) got value")
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

	s := NewSlist()
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

func BenchmarkSlistPushHead(b *testing.B) {
	s := NewSlist()

	for i := 0; i < b.N; i++ {
		s.PushHead(i)
	}
}

func BenchmarkSlistPushHeadAndPopHead(b *testing.B) {
	s := NewSlist()

	for i := 0; i < b.N; i++ {
		s.PushHead(i)
	}
	for i := 0; i < b.N/4; i++ {
		s.PopHead()
	}
	for i := 0; i < b.N/4; i++ {
		s.PopHead()
	}
}
