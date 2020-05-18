package tree

import "fmt"

type node struct {
	val   int
	lnode *node
	rnode *node
}

type tree struct {
	head *node
}

func (t tree) printtreeinorder() {
	printinorder(t.head)
}

func printinorder(n *node) {
	if n.lnode != nil {
		printinorder(n.lnode)
	}
	fmt.Println(n.val)
	if n.rnode != nil {
		printinorder(n.rnode)
	}
}
