package game

type UnitStatus int

const (
	UnitStatusInSupply UnitStatus = iota
	UnitStatusInHand
	UnitStatusInBag
	UnitStatusInDiscarded
	UnitStatusOnBoard
	UnitStatusDead
)

type UnitType int

const (
	UnitTypeLightCavalry UnitType = iota
	UnitTypeArcher
	UnitTypeSwordsman
	UnitTypePickman
	UnitTypeCrossbowman
	UnitTypeMarshal
)

type MovementDirectionType int

const (
	MovementDirectionTypeStraight MovementDirectionType = iota
	MovementDirectionTypeFlexible
)

type AttackDirectionType int

const (
	AttackDirectionTypeStraight AttackDirectionType = iota
	AttackDirectionTypeFlexible
)

type Unit struct {
	ID       string     `json:"id"`
	Type     UnitType   `json:"type"`
	Owner    string     `json:"owner"`
	Position Position   `json:"position,omitempty"`
	Status   UnitStatus `json:"status"`
}

type UnitDefinition struct {
	Type         UnitType              `json:"type"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	MoveRange    int                   `json:"moveRange"`
	MovementType MovementDirectionType `json:"movementType"`
	AttackRange  int                   `json:"attackRange"`
	AttackType   AttackDirectionType   `json:"attackType"`
	Abilities    []string              `json:"abilities"`
}

var UnitDefinitions = map[UnitType]UnitDefinition{
	UnitTypeLightCavalry: {
		Type:         UnitTypeLightCavalry,
		Name:         "轻骑兵",
		Description:  "可以移动两格",
		MoveRange:    2,
		MovementType: MovementDirectionTypeFlexible,
		AttackRange:  1,
		AttackType:   AttackDirectionTypeStraight,
		Abilities:    []string{},
	},
	UnitTypeArcher: {
		Type:         UnitTypeArcher,
		Name:         "弓箭手",
		Description:  "可以远程攻击",
		MoveRange:    1,
		MovementType: MovementDirectionTypeStraight,
		AttackRange:  2,
		AttackType:   AttackDirectionTypeFlexible,
		Abilities:    []string{},
	},
	UnitTypeSwordsman: {
		Type:         UnitTypeSwordsman,
		Name:         "剑士",
		Description:  "近战单位",
		MoveRange:    1,
		MovementType: MovementDirectionTypeStraight,
		AttackRange:  1,
		AttackType:   AttackDirectionTypeStraight,
		Abilities:    []string{},
	},
	UnitTypePickman: {
		Type:         UnitTypePickman,
		Name:         "枪兵",
		Description:  "长矛兵种",
		MoveRange:    1,
		MovementType: MovementDirectionTypeStraight,
		AttackRange:  1,
		AttackType:   AttackDirectionTypeStraight,
		Abilities:    []string{},
	},
	UnitTypeCrossbowman: {
		Type:         UnitTypeCrossbowman,
		Name:         "弩手",
		Description:  "强力远程单位",
		MoveRange:    1,
		MovementType: MovementDirectionTypeStraight,
		AttackRange:  2,
		AttackType:   AttackDirectionTypeStraight,
		Abilities:    []string{},
	},
	UnitTypeMarshal: {
		Type:         UnitTypeMarshal,
		Name:         "元帅",
		Description:  "指挥官单位",
		MoveRange:    1,
		MovementType: MovementDirectionTypeStraight,
		AttackRange:  1,
		AttackType:   AttackDirectionTypeStraight,
		Abilities:    []string{},
	},
}

// 获取单位定义的辅助函数
func GetUnitDefinition(unitType UnitType) (UnitDefinition, bool) {
	def, exists := UnitDefinitions[unitType]
	return def, exists
}

// 获取所有单位定义
func GetAllUnitDefinitions() map[UnitType]UnitDefinition {
	return UnitDefinitions
}
