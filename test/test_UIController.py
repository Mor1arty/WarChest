import unittest
from show import UIController, UnitCoin, GlobalCoin
from show import get_pixel_coordinate


class MyTestCase(unittest.TestCase):
    def test_UIController(self):
        ui = UIController()

        ui.set_coin(coin_name=UnitCoin.LIGHT_CAVALRY, pos=get_pixel_coordinate((2, 3)))
        ui.set_coin(coin_name=UnitCoin.LIGHT_CAVALRY, pos=get_pixel_coordinate((2, 3)))
        ui.set_coin(coin_name=GlobalCoin.PHOENIX_CONTROL_MARKER, pos=get_pixel_coordinate((6, 2)))

        ui.update()


if __name__ == '__main__':
    unittest.main()
