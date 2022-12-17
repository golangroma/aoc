package main

type Tile struct {
	Kind int
	X, Y int
	W    World
}

func (t *Tile) PathNeighbors() []*Tile {
	var neighbors []*Tile

	for _, offset := range [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	} {
		if n := t.W.Tile(t.X+offset[0], t.Y+offset[1]); n != nil {
			current := t.Kind
			next := n.Kind

			if current == 'S' {
				current = 'a'
			}

			if next == 'E' {
				next = 'z'
			}

			if next-current > 1 {
				continue
			}

			neighbors = append(neighbors, n)
		}
	}

	return neighbors
}

func (t *Tile) PathEstimatedCost(toT *Tile) float64 {
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
