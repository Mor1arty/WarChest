import unittest
from game_manager import GameManager
from unit import Coin
from ui import UIController, CoinType, TerrainType


class TestGameManager(unittest.TestCase):

    def setUp(self):
        pass

    def test_game_start(self):
        gm = GameManager()
        gm.game_start()

        # ui = UIController()
        # ui.set_board(gm.board)
        # ui.display()
        # self.assertEqual(True, False)

    def test_deploy(self):
        gm = GameManager()
        gm.game_start()
        coin = Coin(CoinType.LIGHT_CAVALRY)
        gm.deploy(coin, pos=(1, 0))
        self.assertEqual(gm.board.grids[1][0].unit.unit_type, CoinType.LIGHT_CAVALRY)

    def test_init_deck(self):
        gm = GameManager()
        gm.game_start()
        p1 = gm.players[0]
        self.assertEqual(p1.deck.size, 9) # each player have 9 coins in deck initial.


    # def test_init_deck(self):
    #     gm = GameManager()
    #     gm.game_start()
    #     p1 = gm.players[0]
    #     print(p1.supply)
    #     pass


if __name__ == '__main__':
    unittest.main()
