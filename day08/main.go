package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	root := makeTree(input)

	metaSum := 0
	root.Walk(func(n Node) {
		for _, meta := range n.metadata {
			metaSum += meta
		}
	})

	fmt.Println("A:", metaSum)
	fmt.Println("B:", root.Value())
}

func readInput() []int {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	result := []int{}
	for _, word := range strings.Split(strings.TrimSpace(string(input)), " ") {
		n, _ := strconv.Atoi(word)
		result = append(result, n)
	}

	return result
}

type Node struct {
	children []Node
	metadata []int
}

func (n *Node) Walk(f func(Node)) {
	f(*n)
	for _, child := range n.children {
		child.Walk(f)
	}
}

func (n *Node) Value() int {
	result := 0
	if len(n.children) == 0 {
		for _, meta := range n.metadata {
			result += meta
		}
	} else {
		for _, meta := range n.metadata {
			if meta <= len(n.children) {
				result += n.children[meta-1].Value()
			}
		}
	}
	return result
}

func makeTree(nums []int) Node {
	var nextNode func([]int, int) (Node, int)
	nextNode = func(nums []int, idx int) (Node, int) {
		numChildren, numMeta := nums[idx], nums[idx+1]
		children := []Node{}
		metadata := []int{}
		i := idx + 2

		for len(children) < numChildren {
			child, pos := nextNode(nums, i)
			children = append(children, child)
			i = pos
		}

		for len(metadata) < numMeta {
			metadata = append(metadata, nums[i])
			i++
		}

		result := Node{
			children: children,
			metadata: metadata,
		}

		return result, i
	}
	root, _ := nextNode(nums, 0)
	return root
}
