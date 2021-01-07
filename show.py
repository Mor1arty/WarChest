from PIL import Image
from enum import Enum


def get_pixel_coordinate(logic_pos: tuple) -> tuple:
    """
    Return pixel coordinate according to coin's logical position.Let the most left down position be (0, 0).X and y
    increase along up and right direction respectively.
    :param logic_pos: coin's logic position.
    :return: coin's pixel position according to its logic position.
    """
    init_pos = (231, 491)  # 最左下角的点的坐标
    x = init_pos[0] + logic_pos[0] * 78
    if logic_pos[0] < 4:
        y = init_pos[1] + logic_pos[0] * 44 - logic_pos[1] * 88
    else:
        y = init_pos[1] + (6 - logic_pos[0]) * 44 - logic_pos[1] * 88
    px_pos = (x, y)
    return px_pos


class GlobalCoin(Enum):
    """
    Image path of coins except unit coins.(I call them global coins)
    Default in folder imgs/.
    """
    PHOENIX_CONTROL_MARKER = "phoenix_control_marker.png"
    LION_CONTROL_MARKER = "lion_control_marker.png"
    PHOENIX_ROYAL_COIN = "phoenix_royal_coin.png"
    LION_ROYAL_COIN = "lion_royal_coin.png"


class UnitCoin(Enum):
    """
    Image path of unit coins.(Default in folder imgs/)
    """
    LIGHT_CAVALRY = "light_cavalry.png"


class UIController(object):
    def __init__(self):
        self.background = Image.open("imgs/background.png")
        self.set_coin(coin_name=GlobalCoin.PHOENIX_CONTROL_MARKER, pos=get_pixel_coordinate((1, 0)))
        self.set_coin(coin_name=GlobalCoin.PHOENIX_CONTROL_MARKER, pos=get_pixel_coordinate((4, 0)))
        self.set_coin(coin_name=GlobalCoin.LION_CONTROL_MARKER, pos=get_pixel_coordinate((2, 5)))
        self.set_coin(coin_name=GlobalCoin.LION_CONTROL_MARKER, pos=get_pixel_coordinate((5, 4)))

    def set_coin(self, coin_name: Enum, pos: tuple):
        img_path = f"imgs/{coin_name.value}"
        coin = UIController.__read_coin(img_path=img_path)
        self.background.paste(coin, pos, coin)

    def display(self):
        self.background.show()

    @staticmethod
    def __read_coin(img_path: str) -> Image.Image:
        coin = Image.open(img_path)
        coin = coin.resize((78, 78))
        return coin


def main():
    ui = UIController()

    ui.set_coin(coin_name=UnitCoin.LIGHT_CAVALRY, pos=get_pixel_coordinate((2, 3)))
    ui.set_coin(coin_name=UnitCoin.LIGHT_CAVALRY, pos=get_pixel_coordinate((2, 3)))
    ui.set_coin(coin_name=GlobalCoin.PHOENIX_CONTROL_MARKER, pos=get_pixel_coordinate((6, 2)))

    ui.display()


if __name__ == '__main__':
    main()
