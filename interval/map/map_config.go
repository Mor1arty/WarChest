package game_map

type MapConfig struct {
	Grid          *HexGrid
	Layout        map[Cube]CellType
	ControlPoints []Cube
	Castles       []Cube
}

// 预定义地图配置
var (
	DefaultMapConfig = &MapConfig{
		Grid: NewHexGrid(3, 3, 3), // 双人模式
		// Grid: NewHexGrid(5, 3, 3), // 四人模式
	}
)
