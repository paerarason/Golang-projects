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

}
