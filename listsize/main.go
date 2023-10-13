package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	count := 0
	node := l.Head
	for node != nil {
		count++
		node = node.Next
	}
	return count
}
