package main

// https://www.hackerrank.com/challenges/dijkstrashortreach/problem

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Edge struct {
	node int
	cost float64
}

type Node struct {
	cost float64
	num  int
	loc  int
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	bufs := nextBufs(s)
	tCount := toInt(&bufs[0])
	for i := 0; i < tCount; i++ {

		bufs = nextBufs(s)
		nodeCount := toInt(&bufs[0])
		edgeCount := toInt(&bufs[1])

		nodeList := make([]*Node, nodeCount)   // used for printing
		nodeHeap := make(NodeHeap, nodeCount)  // used for dijkstras
		nodeEdges := make([][]Edge, nodeCount) // used for stack storage of edges

		// intialize nodes
		var newNode *Node
		for j := 0; j < nodeCount; j++ {
			newNode = new(Node)
			newNode.cost = math.Inf(1)
			newNode.loc = int(j)
			newNode.num = j
			nodeList[j] = newNode
			nodeHeap[j] = newNode
		}

		// intialize edges
		var a, b int
		var cost float64

		for j := 0; j < edgeCount; j++ {
			bufs = nextBufs(s)
			a = toInt(&bufs[0]) - 1 //  0 indexed
			b = toInt(&bufs[1]) - 1 //  0 indexed
			cost = float64(toInt(&bufs[2]))
			nodeEdges[a] = append(nodeEdges[a], Edge{b, cost})
			nodeEdges[b] = append(nodeEdges[b], Edge{a, cost})
		}

		// intialize start node
		bufs = nextBufs(s)
		start := toInt(&bufs[0])
		start--
		nodeList[start].cost = 0

		// init heap
		heap.Init(&nodeHeap)

		// eval all
		var src, dst *Node
		var newCost float64

		for len(nodeHeap) > 0 {

			// eval connections
			src = heap.Pop(&nodeHeap).(*Node)
			for _, edge := range nodeEdges[src.num] {
				dst = nodeList[edge.node]
				newCost = src.cost + edge.cost
				if newCost < dst.cost {
					dst.cost = newCost
					heap.Fix(&nodeHeap, dst.loc)
				}
			}
		}

		// print nodes
		for i, slvd := range nodeList {
			if int(slvd.cost) < 0 {
				slvd.cost = -1
			}
			if i != start {
				fmt.Printf("%v ", slvd.cost)
			}
		}
		fmt.Printf("\n")
	}
}

func toInt(buf *[]byte) (n int) {
	for _, v := range *buf {
		n = n*10 + int(v-'0')
	}
	return
}

func nextBufs(s *bufio.Scanner) [][]byte {
	s.Scan()
	return bytes.Split(s.Bytes(), []byte{' '})
}

// Heap Implementation
type NodeHeap []*Node

func (h NodeHeap) Len() int {
	return len(h)
}
func (h NodeHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].loc = i
	h[j].loc = j
}
func (h *NodeHeap) Push(x interface{}) {
	node, _ := x.(*Node)
	*h = append(*h, node)
	node.loc = len(*h) - 1
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	x.loc = -1
	*h = old[0 : n-1]
	return x
}

func (h NodeHeap) printHeap() {
	fmt.Println("*#######* heap *#######*")
	var totalRows, totalItems, cellSize, totalWidth int // heap level
	var itemsInRow, offset, indent, rowBeg, rowEnd int  // row level

	totalRows = int(math.Log2(float64(len(h))) + 1)
	totalItems = int(math.Pow(2.0, float64(totalRows-1)))
	cellSize = 9
	totalWidth = totalItems * cellSize

	for r := 0; r < totalRows; r++ {
		itemsInRow = int(math.Pow(2.0, float64(r)))
		offset = totalWidth/itemsInRow - cellSize + 1
		indent = offset / 2
		for x := 0; x < indent; x++ {
			fmt.Printf(" ")
		}

		rowBeg = itemsInRow - 1   // 0 indexed
		rowEnd = itemsInRow*2 - 1 // 0 indexed
		if rowEnd > len(h) {
			rowEnd = len(h)
		}
		for i := rowBeg; i < rowEnd; i++ {
			cost := h[i].cost
			if h[i].cost == math.Inf(1) {
				cost = -1
			}
			fmt.Printf(" %2.v|%1.v|%-2.f ", h[i].num, h[i].loc, cost)
			for x := 0; x < offset; x++ {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
