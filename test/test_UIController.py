import unittest
from ui import UIController, TerrainType
from player import Player
from unit import UnitType


class TestUIController(unittest.TestCase):
    def test_board_display(self):
        ui = UIController()

        ui.add_piece(name=TerrainType.PHOENIX_CONTROL_MARKER.value, x=1, y=0, is_terrain=True)
        ui.add_piece(name=TerrainType.PHOENIX_CONTROL_MARKER.value, x=4, y=0, is_terrain=True)
        ui.add_piece(name=TerrainType.LION_CONTROL_MARKER.value, x=2, y=5, is_terrain=True)
        ui.add_piece(name=TerrainType.LION_CONTROL_MARKER.value, x=5, y=4, is_terrain=True)

        ui.add_piece(name=UnitType.LIGHT_CAVALRY.value, x=1, y=0)
        ui.add_piece(name=UnitType.LIGHT_CAVALRY.value, x=2, y=2)
        ui.add_piece(name=UnitType.LIGHT_CAVALRY.value, x=2, y=5)
        ui.add_piece(name=UnitType.LIGHT_CAVALRY.value, x=4, y=0)

        ui.display_board()
        ui.background.show()

    def test_player_supply_display(self):
        ui = UIController()
        player1 = Player()
        player1.unit_types = [UnitType.LIGHT_CAVALRY, UnitType.ARCHER, UnitType.BERSERKER, UnitType.CROSSBOWMAN]
        # ui.player = player1

        ui.display_player(player1)
        ui.background.show()


if __name__ == '__main__':
    unittest.main()
