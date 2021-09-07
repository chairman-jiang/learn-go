package queue

// interface{} 任意类型
type Queue []interface{}

func (q *Queue) Push(value interface{})  {
	// push 只存储int 类型, 如果外面传入string类型,编译不会保存, 运行时会报错
	*q = append(*q, value.(int))
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

