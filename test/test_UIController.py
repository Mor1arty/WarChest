import unittest
from show import UIController, UnitType, TerrainType


class TestUIController(unittest.TestCase):
    def test_UIController(self):
        ui = UIController()

        ui.add_piece(name=TerrainType.PHOENIX_CONTROL_MARKER.value, x=1, y=0, is_terrain=True)
        ui.add_piece(name=TerrainType.PHOENIX_CONTROL_MARKER.value, x=4, y=0, is_terrain=True)
        ui.add_piece(name=TerrainType.LION_CONTROL_MARKER.value, x=2, y=5, is_terrain=True)
        ui.add_piece(name=TerrainType.LION_CONTROL_MARKER.value, x=5, y=4, is_terrain=True)

        ui.add_piece(name=UnitType.LIGHT_CAVALRY.value, x=1, y=0)
        ui.add_piece(name=UnitType.LIGHT_CAVALRY.value, x=2, y=2)
        ui.add_piece(name=UnitType.LIGHT_CAVALRY.value, x=2, y=5)
        ui.add_piece(name=UnitType.LIGHT_CAVALRY.value, x=4, y=0)

        ui.display()


if __name__ == '__main__':
    unittest.main()
