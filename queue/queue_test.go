package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Empty(t *testing.T) {
	data := []struct{
		q Queue
		result bool
	} {
		{Queue{},true},
		{Queue{1},false},
		{Queue{"ok"},false},
	}
	for _,q := range data {
		if q.q.Empty() != q.result {
			t.Errorf("Empty to TestQueue_Empty,get %v,excepted %v \n",
				q.q.Empty(),q.result)
		}
	}
}

func TestQueue_Pop(t *testing.T) {
	data := []struct {
		q Queue
		result interface{}
	} {
		{Queue{1,2},2},
		{Queue{"ok","nice"},"ok"},
		{Queue{"fail","suc"},"fail"},
	}

	for _,q := range data {
		v,err := q.q.Pop()
		if err != nil {
			t.Errorf("pop is empty")
			continue
		}
		if v != q.result {
			t.Errorf("Pop to TestQueue_Pop,get %b,excepted %b \n",
				v,q.result)
		}
	}
}

func TestQueue_Push(t *testing.T) {
	data := []struct {
		q Queue
		result interface{}
	} {
		{Queue{1,2},2},
		{Queue{"ok","nice"},"ok"},
		{Queue{"fail","suc"},"fail"},
	}
	for _,q := range data {
		v,err := q.q.Pop()
		if err != nil {
			t.Errorf("pop is empty")
			continue
		}
		if v != q.result {
			t.Errorf("Push to TestQueue_Push,get %v,excepted %v \n",
				v,q.result)
		}
	}

}

func ExampleQueue_Empty() {
	q := Queue{}
	q.Push(1)
	fmt.Println(q.Empty())
	// Output:
	// false
}

func ExampleQueue_Pop() {
	q := Queue{}
	q.Push(1)
	fmt.Println(q.Pop())
	// Output:
	// 1 <nil>
}

func ExampleQueue_Push() {
	q := Queue{}
	q.Push(1)
	fmt.Println(q.Pop())
	// Output:
	// 1 <nil>
}


func BenchmarkQueue_Pop(b *testing.B) {
	queue := Queue{}
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}

	for i := 0; i < b.N; i++ {
		queue.Pop()
	}
}


func BenchmarkQueue_Push(b *testing.B) {
	queue := Queue{}
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
}