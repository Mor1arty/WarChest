package game_map

type GameMap struct {
	Config *MapConfig
	Grid   *HexGrid
	Cells  map[Cube]*Cell
}
