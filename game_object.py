#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description:
# ----------------------------------------------------------------------
from enum import Enum


class TerrainType(Enum):
    """
    Image path of coins except unit coins.(I call them global coins)
    Default in folder imgs/.
    """
    PHOENIX_CONTROL_MARKER = "phoenix_control_marker.png"
    LION_CONTROL_MARKER = "lion_control_marker.png"
    PHOENIX_ROYAL_COIN = "phoenix_royal_coin.png"
    LION_ROYAL_COIN = "lion_royal_coin.png"


class UnitType(Enum):
    """
    Image path of unit coins.(Default in folder imgs/)
    """
    LIGHT_CAVALRY = "light_cavalry.png"
    ARCHER = 1
    BERSERKER = 2
    CROSSBOWMAN = 3
    ENSIGN = 4
    FOOTMAN = 5
    LANCER = 6
    MARSHALL = 7
    MERCENARY = 8
    PIKEMAN = 9
    WARRIOR＿PRIEST = 10


class Coin(object):
    def __init__(self, unit_type):
        self.unit_type = unit_type  # 兵种
        self.area = None
        self.deploy_limit = 1


class Terrain (object):
    def __init__(self, terrain_type):
        self.terrain_type = terrain_type


class Area(object):
    def __init__(self):
        self.coins = []


class Unit(Area):
    def __init__(self, coin: Coin):
        super().__init__()
        self.coins = [coin]
        coin.area = self
        self.hp = 1
        self.grid = None

    @property
    def unit_type(self):
        return self.coins[0].unit_type

    def add_coin(self, coin):
        self.coins.append(coin)
        self.hp += 1
        coin.area = self


class Grid(object):
    def __init__(self):
        self.terrain = None
        self.unit = None

    def add_coin(self, coin):
        if self.unit is None:
            self.unit = Unit(coin)
            self.unit.grid = self
        else:
            self.unit.add_coin(coin)


class Board(object):
    def __init__(self, mode="1v1"):
        self.grids = None
        if mode == "1v1":
            self.init_1v1()

    def init_1v1(self):
        cols = [4, 5, 6, 7, 6, 5, 4]
        self.grids = [[Grid() for i in range(cols[row])] for row in range(len(cols))]

        self.grids[1][0].terrain = Terrain(TerrainType.PHOENIX_CONTROL_MARKER)
        self.grids[4][0].terrain = Terrain(TerrainType.PHOENIX_CONTROL_MARKER)
        self.grids[2][5].terrain = Terrain(TerrainType.LION_CONTROL_MARKER)
        self.grids[5][4].terrain = Terrain(TerrainType.LION_CONTROL_MARKER)

    def add_coin(self, coin, pos):
        self.grids[pos[0]][pos[1]].add_coin(coin)

    def unit_num(self, unit_type):
        num = 0
        for grid in self:
            if grid.unit is not None:
                if grid.unit.unit_type == unit_type:
                    num += 1
        return num

    def __iter__(self):
        return iter(sum(self.grids, []))
