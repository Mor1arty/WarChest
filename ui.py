from PIL import Image
from enum import Enum
import game_object  #暂时用后台的数据结构。虽然前后端大概不会分离
import player # 同上
from game_object import TerrainType, UnitType  # Same as above
from unit import UnitType


class UIPiece(object):
    def __init__(self, name=None):
        self.name = name

    @property
    def img_path(self) -> str:
        return f"imgs/{self.name}.png"


class Grid(object):
    def __init__(self, x=0, y=0, piece=None):
        self.x = x
        self.y = y
        self.piece = piece


class UIBoard(object):
    def __init__(self):
        self.grids = []

    def __iter__(self):
        return iter(self.grids)

    def clear(self):
        self.grids = []


class UIController(object):
    def __init__(self):
        self.terrain_board = UIBoard()  # 棋盘，记录棋盘上所有的地形及其位置
        self.coin_board = UIBoard()  # 棋盘，记录棋盘上所有的棋子及其位置
        self.background = Image.open("imgs/background.png")
        self.player = None
        self.player_supply_slots = [(18, 728), (118, 728), (218, 728), (318, 728)]  # slots[i] is the slot list of ith player.
        self.player_hand_slots = []
        self.player_discard_slots = []
        self.enemy = None
        self.enemy_supply_slots = []
        self.enemy_hand_slots = []
        self.enemy_discard_slots = []

    def set_board(self, new_board: game_object.Board):
        self.coin_board.clear()
        self.terrain_board.clear()
        for row in range(len(new_board.grids)):
            for col in range(len(new_board.grids[row])):
                grid = new_board.grids[row][col]
                if grid.terrain is not None:
                    self.add_piece(grid.terrain.terrain_type.value, row, col, is_terrain=True)
                if grid.unit is not None:  # TODO if unit has multi hp
                    self.add_piece(grid.unit.unit_type.value, row, col, is_terrain=False)

    def set_piece(self, piece: UIPiece, pos, is_terrain=False):
        img = self.__read_coin(piece.img_path, is_terrain)
        pos = self.board_pixel_coordinate(pos)
        self.background.paste(img, pos, img)

    # def set_unit(self, unit: Piece, pos, hp):
    #     coin = Image.open(unit.img_path)
    #     coin = coin.resize((78, 78))

    def display_area(self, area_slots, area):
        for i in range(area.size):
            unit = area[i]
            img = Image.open(UIPiece(unit.unit_type.value).img_path)
            img = img.resize((78, 78))
            pos = area_slots[i]
            self.background.paste(img, pos, img)

    def display_player(self, friendly_player):
        self.display_area(self.player_supply_slots, friendly_player.supply)
        pass

    # def display_hand(self, player_idx, coins):
    #     pass
    #
    # def display_discard(self, player_idx, coins):
    #     pass
    #
    # def display_deck(self, player_idx, coins):
    #     pass

    def display_board(self):
        for grid in self.terrain_board:
            self.set_piece(grid.piece, (grid.x, grid.y), is_terrain=True)

        for grid in self.coin_board:
            self.set_piece(grid.piece, (grid.x, grid.y), is_terrain=False)

    def display(self) -> None:
        self.display_board()
        self.display_player(self.player)
        self.background.show()

    def add_piece(self, name: str, x: int, y: int, is_terrain=False):
        if is_terrain:
            self.terrain_board.grids.append(Grid(x, y, UIPiece(name)))
        else:
            self.coin_board.grids.append(Grid(x, y, UIPiece(name)))

    @staticmethod
    def __read_coin(img_path: str, is_terrain=False) -> Image.Image:
        coin = Image.open(img_path)
        if is_terrain:
            coin = coin.resize((85, 85))
        else:
            coin = coin.resize((78, 78))

        return coin

    @staticmethod
    def board_pixel_coordinate(logic_pos: tuple) -> tuple:
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
