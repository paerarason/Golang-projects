package main

import (
	"fmt"
)

func main() {
	fmt.Print("linked list code\n")

	list := createStack()
	push(&list, 45)
	pop(&list)
	push(&list, 45)
	push(&list, 45)
	push(&list, 5)
	push(&list, 35)
	push(&list, 45)
	printstack(&list)
	arr := []int{78, 9, 45, 56, 14}
	heapsort(&arr, 5)
	fmt.Println(arr)
	var p *int
	var num int 
	num=9
	p=&num
	fmt.Println(num)
	fmt.Println(p)
	fmt.Println(*p)

}
