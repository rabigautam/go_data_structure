package main

import "fmt"

type Queue []string

func (q *Queue) IsEmpty() bool {

	return len(*q) == 0
}

func (q *Queue) Enqueue(s string) {
	*q = append(*q, s)
}

func (q *Queue) Dequeue() (string, bool) {

	if q.IsEmpty() {
		return "", false
	}
	element := (*q)[0]
	*q = (*q)[1:]

	return element, true

}

func main() {
	var queue Queue

	queue.Enqueue("Data")
	queue.Enqueue("Structure")
	queue.Enqueue("And")
	queue.Enqueue("Algorithm")

	fmt.Println(queue)

	for !queue.IsEmpty() {
		element, has_data := queue.Dequeue()
		if has_data == true {
			fmt.Println(element)

		}

	}
}
