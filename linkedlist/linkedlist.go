package linkedlist

type Node struct {
	value interface{}
	next  *Node
}

// 由各个内存结构通过一个 Next 指针链接在一起组成，
// 每一个内存结构都存在后继内存结构【链尾除外】，内存结构由数据域和 Next 指针域组成
type LinkedList struct {
	head *Node
}

// 向单向链表末尾插入一个元素
func (l *LinkedList) InsertEnd(v interface{}) {
	node := &Node{value: v}
	p := l.head // 首先头部是为了定位
	if l.head == nil {
		l.head = node
	} else { // 这里以
		for p.next != nil {
			p = p.next
		}
		p.next = node
	}
}

func (l *LinkedList) Travese(callback func(node *Node)) {
	iterator := l.head
	for iterator.next != nil {
		callback(iterator)
		iterator = iterator.next
	}
}

func (l *LinkedList) Search(v interface{}) *Node {
	iterator := l.head
	for iterator.next != nil {
		if iterator.value == v {
			return iterator
		}
		iterator = iterator.next
	}
	return nil
}

func (l *LinkedList) Delete(value interface{}) {
	iterator := l.head
	for iterator.next != nil {
		if iterator.next.next != nil && iterator.next.value == value {
			iterator.next = iterator.next.next
		}
		iterator = iterator.next
	}
}
