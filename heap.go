package main

func heapify(arr *[]int, n int, i int) {
	l := 2*i + 1
	r := 2*i + 2
	largest := i
	if l < n && (*arr)[largest] < (*arr)[l] {
		largest = l
	} else if r < n && (*arr)[largest] < (*arr)[r] {
		largest = r
	}
	if largest != i {
		(*arr)[i], (*arr)[largest] = (*arr)[largest], (*arr)[i]
		heapify(arr, n, largest)
	}

}

func heapsort(arr *[]int, n int) {

	for i := (n / 2) - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		(*arr)[0], (*arr)[i] = (*arr)[i], (*arr)[0]
		heapify(arr, i, 0)
	}
}
