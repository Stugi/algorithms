package main

import "fmt"

func main() {

	graph := &Graph{}
	graph.AddLink(1, 2)
	graph.AddLink(1, 3)
	graph.AddLink(1, 4)
	graph.AddLink(3, 5)

	graph.BFS()
}

type Graph struct {
	//array of adjacent vertices
	elems map[Elem][]Elem
}

type Elem struct {
	value  int
	marked bool
}

func (g *Graph) AddElem(value int) {
	if g.elems == nil {
		g.elems = make(map[Elem][]Elem)
	} else if _, ok := g.elems[Elem{value, false}]; ok {
		return
	}
	g.elems[Elem{value, false}] = []Elem{}

}

func (g *Graph) AddLink(from int, to int) {
	if g.elems == nil {
		g.elems = make(map[Elem][]Elem)
	}
	if _, ok := g.elems[Elem{from, false}]; !ok {
		g.AddElem(from)
	}
	if _, ok := g.elems[Elem{to, false}]; !ok {
		g.AddElem(to)
	}

	g.elems[Elem{from, false}] = append(g.elems[Elem{from, false}], Elem{to, false})
	g.elems[Elem{to, false}] = append(g.elems[Elem{to, false}], Elem{from, false})
}

// Обход в ширину
func (g *Graph) BFS() {

	if g.elems == nil {
		return
	}

	queue := make([]Elem, 0)
	queue = append(queue, Elem{1, false})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Println(current.value)
		for _, elem := range g.elems[current] {
			if !elem.marked {
				queue = append(queue, elem)
				elem.marked = true
			}
		}
	}
}
