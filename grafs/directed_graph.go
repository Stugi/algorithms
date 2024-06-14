package main

import (
	"fmt"
	"sort"
)

func main() {

	// Directed Graph
	var graph = make(map[string]*Node)

	addEdge(graph, "A", "B")
	addEdge(graph, "B", "A")

	addEdge(graph, "A", "C")
	addEdge(graph, "C", "A")

	addEdge(graph, "B", "C")
	addEdge(graph, "C", "B")

	addEdge(graph, "A", "D")
	addEdge(graph, "D", "A")

	addEdge(graph, "A", "E")
	addEdge(graph, "E", "A")

	addEdge(graph, "E", "F")
	addEdge(graph, "F", "E")
	addEdge(graph, "E", "G")
	addEdge(graph, "G", "E")
	addEdge(graph, "E", "I")
	addEdge(graph, "I", "E")
	addEdge(graph, "E", "H")
	addEdge(graph, "H", "E")

	addEdge(graph, "F", "G")
	addEdge(graph, "G", "F")

	addEdge(graph, "G", "I")
	addEdge(graph, "I", "G")

	addEdge(graph, "H", "I")
	addEdge(graph, "I", "H")

	addEdge(graph, "J", "I")
	addEdge(graph, "I", "J")

	printGraph(graph)

	minDistance, err := findMinDistance(graph, "C", "J")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(minDistance)
	}

	path, err := findMinPath(graph, "C", "J")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(path)
	}
}

type Node struct {
	id            string
	adjacentNodes []*Node
	marked        bool
	distance      int
}

func NewNode(id string) *Node {
	return &Node{id: id, marked: false, distance: -1}
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

func findMinDistance(graf map[string]*Node, startNodeId string, endNodeId string) (int, error) {
	if startNodeId == endNodeId {
		return 0, nil
	}
	if _, ok := graf[startNodeId]; !ok {
		return 0, fmt.Errorf("startNodeId: %s not found", startNodeId)
	}
	if _, ok := graf[endNodeId]; !ok {
		return 0, fmt.Errorf("endNodeId: %s not found", endNodeId)
	}

	clear(graf)
	bfs(graf, startNodeId)

	distance := graf[endNodeId].distance

	clear(graf)

	return distance, nil
}

func findMinPath(graf map[string]*Node, startNodeId string, endNodeId string) ([]string, error) {
	if startNodeId == endNodeId {
		return []string{startNodeId}, nil
	}
	if _, ok := graf[startNodeId]; !ok {
		return nil, fmt.Errorf("startNodeId: %s not found", startNodeId)
	}
	if _, ok := graf[endNodeId]; !ok {
		return nil, fmt.Errorf("endNodeId: %s not found", endNodeId)
	}

	clear(graf)

	path := make([]string, 0)
	bfs(graf, startNodeId)

	path = append(path, graf[endNodeId].id)
	node := graf[endNodeId]

	for node.id != startNodeId {
		for _, adjacentNode := range node.adjacentNodes {
			if adjacentNode.distance == node.distance-1 {
				path = append(path, adjacentNode.id)
				node = adjacentNode
				break
			}
		}
	}

	clear(graf)
	return path, nil
}

func bfs(graf map[string]*Node, startNodeId string) {
	startNode := graf[startNodeId]
	startNode.marked = true
	startNode.distance = 0

	queue := make([]*Node, 0)
	queue = append(queue, startNode)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, adjacentNode := range node.adjacentNodes {
			if adjacentNode.distance == -1 {
				adjacentNode.distance = node.distance + 1
			} else {
				if adjacentNode.distance > node.distance+1 {
					adjacentNode.distance = node.distance + 1
				}
				// adjacentNode.distance =math.Min(adjacentNode.distance, node.distance+1)
			}
			if !adjacentNode.marked {
				queue = append(queue, adjacentNode)
			}
			adjacentNode.marked = true
			// sort queue by distance
			sort.Slice(queue, func(i, j int) (less bool) {
				return queue[i].distance < queue[j].distance
			})

		}
	}
}

func clear(graf map[string]*Node) {
	for _, node := range graf {
		node.marked = false
		node.distance = -1
	}
}
