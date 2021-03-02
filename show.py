from PIL import Image
from enum import Enum


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


class Piece(object):
    def __init__(self, name=None):
        self.name = name

    @property
    def img_path(self) -> str:
        return f"imgs/{self.name}"


class Grid(object):
    def __init__(self, x=0, y=0, piece=None):
        self.x = x
        self.y = y
        self.piece = piece


class Board(object):
    def __init__(self):
        self.grids = []

    def __iter__(self):
        return iter(self.grids)


class UIController(object):
    def __init__(self):
        self.terrain_board = Board()  # 棋盘，记录棋盘上所有的地形及其位置
        self.coin_board = Board()  # 棋盘，记录棋盘上所有的棋子及其位置
        self.background = Image.open("imgs/background.png")

    def set_piece(self, piece: Piece, pos, is_terrain=False):
        img = self.__read_coin(piece.img_path, is_terrain)
        pos = self.get_pixel_coordinate(pos)
        self.background.paste(img, pos, img)

    def display(self) -> None:
        for grid in self.terrain_board:
            self.set_piece(grid.piece, (grid.x, grid.y), is_terrain=True)

        for grid in self.coin_board:
            self.set_piece(grid.piece, (grid.x, grid.y), is_terrain=False)
        self.background.show()

    def add_piece(self, name: str, x: int, y: int, is_terrain=False):
        if is_terrain:
            self.terrain_board.grids.append(Grid(x, y, Piece(name)))
        else:
            self.coin_board.grids.append(Grid(x, y, Piece(name)))

    @staticmethod
    def __read_coin(img_path: str, is_terrain=False) -> Image.Image:
        coin = Image.open(img_path)
        if is_terrain:
            coin = coin.resize((85, 85))
        else:
            coin = coin.resize((78, 78))

        return coin

    @staticmethod
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


def main():
    pass


if __name__ == '__main__':
    main()
