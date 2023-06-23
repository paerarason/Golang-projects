package queue

func dequeue(queue *List) {
	remove_last(queue)
}

func enqueue(queue *List, num int) {
	add(queue, num)
}

func createQueue() *List {
	var list List
	list.head = &Node{}
	return &list
}
