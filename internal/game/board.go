package game

// Position 立体坐标
type Position struct {
	Q int `json:"q"`
	R int `json:"r"`
	S int `json:"s"`
}

type BoardSize struct {
	QSize int `json:"qSize"`
	RSize int `json:"rSize"`
	SSize int `json:"sSize"`
}

type CellType int

const (
	CellTypeBlocked CellType = iota
	CellTypeNormal
	CellTypeControlPoint
	CellTypeCastle
)

// BoardCell 棋盘格子信息
type BoardCell struct {
	UnitID       *string  `json:"unitId,omitempty"`       // 可选字段使用指针
	ControlledBy *string  `json:"controlledBy,omitempty"` // playerId
	CellType     CellType `json:"cellType"`
}
