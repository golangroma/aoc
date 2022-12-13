package main

type World map[int]map[int]*Tile

func (w World) Tile(x, y int) *Tile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

func (w World) SetTile(t *Tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*Tile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

func (w World) Find(kind int) *Tile {
	for _, row := range w {
		for _, t := range row {
			if t.Kind == kind {
				return t
			}
		}
	}
	return nil
}

func (w World) From() *Tile {
	return w.Find('S')
}

func (w World) To() *Tile {
	return w.Find('E')
}

func ParseWorld(input []string) World {
	w := World{}

	for y, row := range input {
		for x, raw := range row {
			w.SetTile(&Tile{
				Kind: int(raw),
			}, x, y)
		}
	}

	return w
}
