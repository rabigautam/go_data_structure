package main

import "fmt"

const Alphabet_Size = 26

type TriNode struct {
	children   [Alphabet_Size]*TriNode
	endOfWords bool
}
//setup the node 
func getNode() *TriNode {
	node := new(TriNode)
	node.endOfWords = false
	for i := 0; i < Alphabet_Size; i++ {
		node.children[i] = nil
	}
	return node
}
func insertNode(root *TriNode, key string) {
	temp := root

	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] == nil {
			temp.children[index] = getNode()
		}
		temp = temp.children[index] 

	}  
	temp.endOfWords = true
}

func search(root *TriNode, key string) bool {
	temp := root

	for i := 0; i <len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			return false
		}

	}
	return (temp != nil && temp.endOfWords)
}
func main() {
	words := []string{"a", "and", "go", "golang", "mango", "mongo"}
	root := getNode()
	for i := 0; i < len(words); i++ {

		insertNode(root, words[i])

	}
	fmt.Println("contains the word and", search(root, "and"))
}
