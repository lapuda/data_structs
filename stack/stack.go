package stack

// Stack 的数据结构 FILO 后进先出
type Stack []interface{}

// 压数据入栈
func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

// 出栈
func (s *Stack) Pop() (interface{}, error) {
	length := len(*s)
	if length == 0 {
		return nil, StackError{
			"This stack is empty!",
		}
	}
	lastElement := (*s)[length-1]
	*s = (*s)[:length-1]
	return lastElement, nil
}

// 判栈是否为空
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// 获取栈的大小
func (s *Stack) Size() int {
	return len(*s)
}
