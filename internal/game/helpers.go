package game

import "github.com/Mor1arty/WarChest/pkg/utils"

func GenerateUnit(t UnitType, o string, s UnitStatus) *Unit {
	id := "unit_" + utils.GenerateUUID()
	return &Unit{
		ID:     id,
		Type:   t,
		Owner:  o,
		Status: s,
	}
}

func GetUnitDefinition(unitType UnitType) (UnitDefinition, bool) {
	def, exists := UnitDefinitions[unitType]
	return def, exists
}

func GetAllUnits() []UnitType {
	units := make([]UnitType, 0, int(UnitTypeMax))
	for i := UnitType(0); i < UnitTypeMax; i++ {
		units = append(units, i)
	}
	return units
}

func GenerateTeam(userID string, room *GameRoom) Team {
	team := TeamWhite
	for existUserID, existTeam := range room.Players {
		if existUserID != userID {
			if existTeam == TeamWhite {
				team = TeamBlack
				break
			}
		}
	}
	return team
}

func GetAllUnitDefinitions() map[UnitType]UnitDefinition {
	return UnitDefinitions
}
