package game_map

type Cell struct {
	Position Cube
	Type     CellType
	Occupant *ChessPiece
}

func NewCell(pos Cube, cellType CellType) *Cell {
	return &Cell{
		Position: pos,
		Type:     cellType,
	}
}
