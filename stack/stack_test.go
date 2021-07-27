package stack

import (
	"fmt"
	"testing"
)

func TestStack_Empty(t *testing.T) {
	data := []struct{
		s Stack
		result bool
	} {
		{Stack{},true},
		{Stack{1},false},
		{Stack{"ok"},false},
	}
	for _,s2 := range data {
		if s2.s.Empty() != s2.result {
			t.Errorf("Empty to TestStack_Empty,get %v,excepted %v \n",
				s2.s.Empty(),s2.result)
		}
	}
}

func TestStack_Pop(t *testing.T) {
	data := []struct {
		s Stack
		result interface{}
	} {
        {Stack{1,2},2},
		{Stack{"ok","nice"},"nice"},
		{Stack{"fail","suc"},"suc"},
	}

	for _,s := range data {
		v,err := s.s.Pop()
		if err != nil {
			t.Errorf("pop is empty")
			continue
		}
		if v != s.result {
			t.Errorf("Pop to TestStack_Pop,get %b,excepted %b \n",
				v,s.result)
		}
	}
}

func TestStack_Push(t *testing.T) {
	data := []struct {
		s Stack
		result interface{}
	} {
		{Stack{1,2},2},
		{Stack{"ok","nice"},"nice"},
		{Stack{"fail","suc"},"suc"},
	}
	for _,s := range data {
		v,err := s.s.Pop()
		if err != nil {
			t.Errorf("pop is empty")
			continue
		}
		if v != s.result {
			t.Errorf("Push to TestStack_Push,get %v,excepted %v \n",
				v,s.result)
		}
	}

}

func ExampleStack_Empty() {
	s := Stack{}
	s.Push(1)
	fmt.Println(s.Empty())
	// Output:
	// false
}

func ExampleStack_Pop() {
	s := Stack{}
	s.Push(1)
	fmt.Println(s.Pop())
	// Output:
	// 1 <nil>
}

func ExampleStack_Push() {
	s := Stack{}
	s.Push(1)
	fmt.Println(s.Pop())
	// Output:
	// 1 <nil>
}


func BenchmarkStack_Pop(b *testing.B) {
	stack := Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}


func BenchmarkStack_Push(b *testing.B) {
	stack := Stack{}
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}