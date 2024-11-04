package game

import "fmt"

// Position 立体坐标
type Position struct {
	Q int `json:"q"`
	R int `json:"r"`
	S int `json:"s"`
}

func (p Position) String() string {
	return fmt.Sprintf("%d,%d,%d", p.Q, p.R, p.S)
}

type BoardSize struct {
	QSize int `json:"qSize"`
	RSize int `json:"rSize"`
	SSize int `json:"sSize"`
}

type CellType string

const (
	CellTypeBlocked      CellType = "Blocked"
	CellTypeNormal       CellType = "Normal"
	CellTypeControlPoint CellType = "ControlPoint"
	CellTypeCastle       CellType = "Castle"
)

// BoardCell 棋盘格子信息
type BoardCell struct {
	UnitID       *string  `json:"unitId,omitempty"`       // 可选字段使用指针
	ControlledBy *string  `json:"controlledBy,omitempty"` // playerId
	CellType     CellType `json:"cellType"`
}

type Board struct {
	Size  BoardSize             `json:"size"`
	Cells map[string]*BoardCell `json:"cells"` // key: position2string
}

type HistoryRecord struct {
}

func createCell(cellType CellType) *BoardCell {
	return &BoardCell{
		CellType: cellType,
	}
}

func createBoardCells(size BoardSize, controlPoints []Position, blockedCells []Position) map[string]*BoardCell {
	boardCells := make(map[string]*BoardCell)
	for q := -size.QSize; q <= size.QSize; q++ {
		for r := -size.RSize; r <= size.RSize; r++ {
			s := -q - r
			if s >= -size.SSize && s <= size.SSize {
				boardCells[Position{Q: q, R: r, S: s}.String()] = createCell(CellTypeNormal)
			}
		}
	}
	for _, cp := range controlPoints {
		boardCells[cp.String()] = createCell(CellTypeControlPoint)
	}
	for _, bc := range blockedCells {
		boardCells[bc.String()] = createCell(CellTypeBlocked)
	}
	return boardCells
}

func CreateBoard(playerCount int) *Board {
	var size BoardSize
	var controlPoints []Position
	var blockedCells []Position

	controlPoints = []Position{
		{Q: 1, R: 2, S: -3},
		{Q: -1, R: -2, S: 3},
		{Q: -2, R: 3, S: -1},
		{Q: 2, R: -3, S: 1},
		{Q: 2, R: 0, S: -2},
		{Q: -2, R: 0, S: 2},
		{Q: -1, R: 1, S: 0},
		{Q: 1, R: -1, S: 0},
		{Q: -3, R: 2, S: 1},
		{Q: -3, R: -2, S: -1},
		{Q: -4, R: 1, S: 3},
		{Q: -4, R: -1, S: -3},
		{Q: -5, R: 3, S: 2},
		{Q: -5, R: -3, S: 2},
	}

	if playerCount == 2 {
		// 2人游戏棋盘大小
		size = BoardSize{
			QSize: 5,
			RSize: 3,
			SSize: 3,
		}
		blockedCells = []Position{
			{Q: 4, R: -3, S: -1},
			{Q: 4, R: -2, S: -2},
			{Q: 4, R: -1, S: -3},
			{Q: 5, R: -3, S: -2},
			{Q: 5, R: -2, S: -3},
			{Q: -4, R: 1, S: 3},
			{Q: -4, R: 2, S: 2},
			{Q: -4, R: 3, S: 1},
			{Q: -5, R: 3, S: 2},
			{Q: -5, R: 2, S: 3},
		}
	} else {
		size = BoardSize{
			QSize: 5,
			RSize: 3,
			SSize: 3,
		}
		controlPoints = []Position{
			{Q: 0, R: 0, S: 0},
			{Q: 4, R: 2, S: 0},
		}
	}

	return &Board{
		Size:  size,
		Cells: createBoardCells(size, controlPoints, blockedCells),
	}
}
