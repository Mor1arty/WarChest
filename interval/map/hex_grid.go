package game_map

type HexGrid struct {
	Q_Size int
	R_Size int
	S_Size int
}

func NewHexGrid(qSize, rSize, sSize int) *HexGrid {
	return &HexGrid{
		Q_Size: qSize,
		R_Size: rSize,
		S_Size: sSize,
	}
}

var Directions = []Cube{
	{Q: 1, R: -1, S: 0}, // 右上
	{Q: 0, R: -1, S: 1}, // 上
	{Q: -1, R: 0, S: 1}, // 左上
	{Q: -1, R: 1, S: 0}, // 左下
	{Q: 0, R: -1, S: 1}, // 下
	{Q: 1, R: 0, S: -1}, // 右下
}

func (g *HexGrid) IsInBounds(pos Cube) bool {
	return (pos.Q >= -g.Q_Size && pos.Q <= g.Q_Size) &&
		(pos.R >= -g.R_Size && pos.R <= g.R_Size) &&
		(pos.S >= -g.S_Size && pos.S <= g.S_Size)
}

func (g *HexGrid) GetNeighbor(pos Cube, direction Cube) Cube {
	return Cube{
		Q: pos.Q + direction.Q,
		R: pos.R + direction.R,
		S: pos.S + direction.S,
	}
}

func (g *HexGrid) GetNeighbors(pos Cube) []Cube {
	neighbors := []Cube{}
	for _, direction := range Directions {
		neighbor := g.GetNeighbor(pos, direction)
		if g.IsInBounds(neighbor) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}
