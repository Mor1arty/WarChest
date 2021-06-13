#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description:
# ----------------------------------------------------------------------
from enum import Enum
from unit import Unit, CoinType, UnitCard, Coin


class TerrainType(Enum):
    """
    Image path of coins except unit coins.(I call them global coins)
    Default in folder imgs/.
    """
    UNKNOWN = 0
    PHOENIX_CONTROL_MARKER = "phoenix_control_marker"
    LION_CONTROL_MARKER = "lion_control_marker"
    PHOENIX_ROYAL_COIN = "phoenix_royal_coin"
    LION_ROYAL_COIN = "lion_royal_coin"


class Terrain (object):
    def __init__(self, terrain_type):
        self.terrain_type = terrain_type

    def __str__(self):
        return f"{self.terrain_type.name}"


class Area(object):
    def __init__(self, unit_types: CoinType=None):
        self.units: list = []
        self.other_unit = Unit(CoinType.OTHERS)
        if unit_types is not None:
            self.init_unit_types(unit_types)

    def init_unit_types(self, unit_types):
        for unit_type in unit_types:
            self.units.append(Unit(unit_type))

    def add_coin(self, coin: Coin):
        for unit in self.units:
            if unit.unit_type == coin.unit_type:
                unit.add_coin(coin)
                return
        self.other_unit.add_coin(coin)

    def pop_coin(self, coin: Coin):
        for unit in self.units:
            if unit.unit_type == coin.unit_type:
                return unit.pop_coin()
        return self.other_unit.pop_coin(coin)

    def pop_coin_by_type(self, unit_type: CoinType):
        for unit in self.units:
            if unit.unit_type == unit_type:
                return unit.pop_coin()
        return self.other_unit.pop_coin(unit_type)

    def __getitem__(self, key):
        return self.units[key]

    @property
    def size(self):
        return len(self.coins)

    @property
    def coins(self):
        coins = []
        for unit in self.units:
            for coin in unit.coins:
                coins.append(coin)
        return coins + self.other_unit.coins

    def __str__(self):
        info = f"The area have {self.size} coins\n"
        for unit in self.units:
            info += str(unit) + '\n'
        info += str(','.join(str(coin) for coin in self.other_unit.coins))
        return info



class Grid(object):
    def __init__(self):
        self.terrain = None
        self.unit = None

    def add_coin(self, coin):
        """
        Add coin to a grid. If the grid have no unit, make a unit. Otherwise, stack (bolster).

        """
        if self.unit is None:
            self.unit = Unit.create_unit_by_coin(coin)
            self.unit.grid = self
        else:
            self.unit.add_coin(coin)

    def __str__(self):
        return f"Terrain: {self.terrain}\nUnit: {self.unit}"


class Board(object):
    def __init__(self, mode="1v1"):
        self.grids = None  # 2D list.
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
        """
        Add coin to specific location.

        """
        self.grids[pos[0]][pos[1]].add_coin(coin)

    def unit_num(self, unit_type):
        """
        Calculate number of unit of specific unit types.

        """
        num = 0
        for grid in self:
            if grid.unit is not None:
                if grid.unit.unit_type == unit_type:
                    num += 1
        return num

    def __iter__(self):
        return iter(sum(self.grids, []))
