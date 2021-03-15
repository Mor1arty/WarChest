import unittest
from game_manager import GameManager
from game_object import Coin
from ui import UIController, UnitType, TerrainType


class TestGameManager(unittest.TestCase):

    def setUp(self):
        pass

    def test_game_start(self):
        gm = GameManager()
        gm.game_start()

        ui = UIController()
        ui.set_board(gm.board)
        ui.display()
        # self.assertEqual(True, False)

    def test_deploy(self):
        gm = GameManager()
        gm.game_start()
        coin = Coin(UnitType.LIGHT_CAVALRY)
        gm.deploy(coin, pos=(1, 0))
        ui = UIController()
        ui.set_board(gm.board)
        ui.display()




if __name__ == '__main__':
    unittest.main()
