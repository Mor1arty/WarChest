package game_map

// CellType 棋盘格子类型
type CellType int

const (
	CellTypeBlocked CellType = iota
	CellTypeNormal
	CellTypeControlPoint
	CellTypeCastle
)

// Cube 立体坐标
type Cube struct {
	Q int
	R int
	S int
}
