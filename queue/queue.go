package queue

// Stack 的数据结构 FIFO 后进先出
type Queue []interface{}

// 压数据入队列
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

// 出队
func (q *Queue) Pop() (interface{}, error) {
	length := len(*q)
	if length == 0 {
		return nil, QueueError{
			"This stack is empty!",
		}
	}
	first := (*q)[0]
	*q = (*q)[1:]
	return first, nil
}

// 判队列是否为空
func (q *Queue) Empty() bool {
	return len(*q) == 0
}

// 获取队列的大小
func (q *Queue) Size() int {
	return len(*q)
}
