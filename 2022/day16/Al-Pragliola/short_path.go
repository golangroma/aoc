package main

import "container/heap"

type Tile struct {
	Name string
}

func (t *Tile) PathNeighbors(w map[string][]*Tile) []*Tile {
	return w[t.Name]
}

func (t *Tile) PathEstimatedCost(to *Tile) float64 {
	return 1.0
}

type node struct {
	tile   *Tile
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

type nodeMap map[string]*node

func (nm nodeMap) get(p *Tile) *node {
	n, ok := nm[p.Name]
	if !ok {
		n = &node{
			tile: p,
		}
		nm[p.Name] = n
	}
	return n
}

func Path(from, to *Tile, world map[string][]*Tile) (path []*Tile, distance float64, found bool) {
	nm := nodeMap{}
	nq := &priorityQueue{}
	heap.Init(nq)
	fromNode := nm.get(from)
	fromNode.open = true
	heap.Push(nq, fromNode)
	for {
		if nq.Len() == 0 {
			return
		}
		current := heap.Pop(nq).(*node)
		current.open = false
		current.closed = true

		if current == nm.get(to) {
			var p []*Tile
			curr := current
			for curr != nil {
				p = append(p, curr.tile)
				curr = curr.parent
			}
			return p, current.cost, true
		}

		for _, neighbor := range current.tile.PathNeighbors(world) {
			cost := current.cost + 1.0
			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.PathEstimatedCost(to)
				neighborNode.parent = current
				heap.Push(nq, neighborNode)
			}
		}
	}
}
