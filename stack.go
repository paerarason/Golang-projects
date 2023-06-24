package main

func pop(queue *List) {
	remove_last(queue)
}

func push(queue *List, num int) {
	add(queue, num)
}

func createStack() List {
	var list List
	list.head = &Node{}
	return list
}

func printstack(l *List) {
	Printlink(l)
}
