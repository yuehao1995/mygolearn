package queue

type Queue []int

func (q *Queue) Push(v int) {
	//go语言指针对象所指向的对象本身可以改变
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
