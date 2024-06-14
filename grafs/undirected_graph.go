package main

import "fmt"

func main() {

	// Undirected Graph
	var graph = make(map[string]*Node)

	addEdge(graph, "A", "B")
	addEdge(graph, "A", "E")
	addEdge(graph, "A", "C")
	addEdge(graph, "A", "D")

	// addEdge(graph, "B", "A")
	addEdge(graph, "B", "C")

	// addEdge(graph, "C", "A")
	// addEdge(graph, "C", "B")

	// addEdge(graph, "D", "A")

	addEdge(graph, "C", "G")

	addEdge(graph, "G", "V")

	printGraph(graph)

	bfs(graph, "A")

}

type Node struct {
	id            string
	adjacentNodes []*Node
	marked        bool
}

func NewNode(id string) *Node {
	return &Node{id: id, marked: false}
}

func addNode(graph map[string]*Node, id string) *Node {
	if val, ok := graph[id]; ok {
		return val
	} else {
		node := NewNode(id)
		graph[id] = node
		return node
	}
}

func addEdge(graph map[string]*Node, id1 string, id2 string) {
	node1 := addNode(graph, id1)
	node2 := addNode(graph, id2)
	node1.adjacentNodes = append(node1.adjacentNodes, node2)
	node2.adjacentNodes = append(node2.adjacentNodes, node1)
}

func printGraph(graph map[string]*Node) {
	for _, node := range graph {
		printNode(node)
	}
}

func printNode(node *Node) {
	fmt.Print(node.id, " -> ")
	for _, adjacentNode := range node.adjacentNodes {
		fmt.Print(adjacentNode.id, " ")
	}
	fmt.Println()
}

// BFS - поиск в ширину
func bfs(graf map[string]*Node, startNodeId string) {
	clearMarked(graf)
	startNode := graf[startNodeId]

	startNode.marked = true

	queue := make([]*Node, 0)
	queue = append(queue, startNode)

	for len(queue) > 0 { // пока очередь не пустая
		node := queue[0] // берем первый элемент
		queue = queue[1:]
		fmt.Println(node.id)

		for _, adjacentNode := range node.adjacentNodes {
			if !adjacentNode.marked {
				adjacentNode.marked = true
				queue = append(queue, adjacentNode)
			}
		}
	}

}

func clearMarked(graf map[string]*Node) {
	for _, node := range graf {
		node.marked = false
	}
}
