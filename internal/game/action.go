package game

import (
	"encoding/json"
	"fmt"
)

type ActionType int

const (
	ActionTypeMove ActionType = iota
	ActionTypeAttack
	ActionTypeDeploy
	ActionTypeBolster
	ActionTypeClaimInitiative
	ActionTypeRecruit
	ActionTypePass
	ActionTypeControl
	ActionTypeClear
)

type ServerActionType int

const (
	ServerActionTypeUpdateGameState ServerActionType = iota
	ServerActionTypeClearGame
)

// MovePayload 移动操作的数据结构
type MovePayload struct {
	UnitID string   `json:"unitId"`
	To     Position `json:"to"`
}

// AttackPayload 攻击操作的数据结构
type AttackPayload struct {
	AttackerID       string   `json:"attackerId"`
	AttackerPosition Position `json:"attackerPosition"`
	TargetID         string   `json:"targetId"`
	TargetPosition   Position `json:"targetPosition"`
}

// DeployPayload 部署操作的数据结构
type DeployPayload struct {
	UnitID   string   `json:"unitId"`
	Position Position `json:"position"`
}

// BolsterPayload 增强操作的数据结构
type BolsterPayload struct {
	UnitID string `json:"unitId"`
}

// ClaimInitiativePayload 获取先手权操作的数据结构
type ClaimInitiativePayload struct {
	UnitID string `json:"unitId"`
}

// RecruitPayload 招募操作的数据结构
type RecruitPayload struct {
	UnitID string `json:"unitId"`
}

// PassPayload 跳过操作的数据结构
type PassPayload struct {
	UnitID string `json:"unitId"`
}

// ControlPayload 控制点操作的数据结构
type ControlPayload struct {
	UnitID   string   `json:"unitId"`
	Position Position `json:"position"`
}

// GameAction 游戏动作结构体，用于包装所有可能的操作
type GameAction struct {
	Type    ActionType  `json:"type"`
	Payload interface{} `json:"payload"`
}

// UnmarshalGameAction 根据动作类型解析对应的 Payload
func UnmarshalGameAction(actionType ActionType, payloadData []byte) (interface{}, error) {
	var payload interface{}

	switch actionType {
	case ActionTypeMove:
		payload = &MovePayload{}
	case ActionTypeAttack:
		payload = &AttackPayload{}
	case ActionTypeDeploy:
		payload = &DeployPayload{}
	case ActionTypeBolster:
		payload = &BolsterPayload{}
	case ActionTypeClaimInitiative:
		payload = &ClaimInitiativePayload{}
	case ActionTypeRecruit:
		payload = &RecruitPayload{}
	case ActionTypePass:
		payload = &PassPayload{}
	case ActionTypeControl:
		payload = &ControlPayload{}
	default:
		return nil, fmt.Errorf("unknown action type: %v", actionType)
	}

	if err := json.Unmarshal(payloadData, payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %v", err)
	}

	return payload, nil
}

// ValidatePayload 验证 Payload 数据的有效性
func ValidatePayload(actionType ActionType, payload interface{}) error {
	switch actionType {
	case ActionTypeMove:
		return validateMovePayload(payload.(*MovePayload))
	case ActionTypeAttack:
		return validateAttackPayload(payload.(*AttackPayload))
		// ... 其他验证函数
	}
	return nil
}

// 验证函数示例
func validateMovePayload(payload *MovePayload) error {
	if payload.UnitID == "" {
		return fmt.Errorf("unitId cannot be empty")
	}
	// 验证位置的有效性
	if !isValidPosition(payload.To) {
		return fmt.Errorf("invalid position")
	}
	return nil
}

func validateAttackPayload(payload *AttackPayload) error {
	if payload.AttackerID == "" || payload.TargetID == "" {
		return fmt.Errorf("attacker and target IDs cannot be empty")
	}
	// 验证位置的有效性
	if !isValidPosition(payload.AttackerPosition) || !isValidPosition(payload.TargetPosition) {
		return fmt.Errorf("invalid position")
	}
	return nil
}

// 辅助函数
func isValidPosition(pos Position) bool {
	// 验证六边形坐标的有效性：q + r + s 应该等于 0
	return pos.Q+pos.R+pos.S == 0
}
