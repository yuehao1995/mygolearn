package queue

//interface可以表示任意类型
type Queue []interface{}

/*
//interface{}表示任意类型
func (q *Queue) Push(v interface{}){
	//go语言指针对象所指向的对象本身可以改变
	*q =append(*q,v)
}

func (q *Queue) Pop() interface{}{
	head :=(*q)[0]
	*q =(*q)[1:]
	return head
}
*/

//限定只能操作int类型
/*
在参数中限定类型
func (q *Queue) Push(v int){
	//go语言指针对象所指向的对象本身可以改变
	*q =append(*q,v)
}

func (q *Queue) Pop() int{
	head :=(*q)[0]
	*q =(*q)[1:]
	//对head类型进行了强制抓换,转换为int
	return head.(int)
}*/

//强制取出指定类型，从而进行类型限定
func (q *Queue) Push(v interface{}) {
	//go语言指针对象所指向的对象本身可以改变
	*q = append(*q, v.(int))
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	//对head类型进行了强制抓换,转换为int
	return head.(int)
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
