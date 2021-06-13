#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
from enum import Enum



class UnitType(Enum):
    """
    Image path of unit coins.(Default in folder imgs/)
    """
    UNKNOWN = 0
    OTHERS = 1
    LIGHT_CAVALRY = "light_cavalry"
    ARCHER = "light_cavalry"
    BERSERKER = "light_cavalry"
    CROSSBOWMAN = "light_cavalry"
    ENSIGN = 4
    FOOTMAN = 5
    LANCER = 6
    MARSHALL = 7
    MERCENARY = 8
    PIKEMAN = 9
    WARRIOR_PRIEST = 10


class UnitCard(object):
    def __init__(self, unit_type: UnitType=None):
        self.unit_type: UnitType = unit_type  # 兵种
        self.deploy_limit = 1
        self.coin_limit = 5  # TODO read config


class Coin(UnitCard):
    def __init__(self, unit_type):
        super().__init__(unit_type)
        self.unit = None
        self.is_revealed = True

    @property
    def area(self):
        return self.unit.area


class Unit(UnitCard):
    def __init__(self, unit_type: UnitType):
        super().__init__(unit_type)
        self.coins: list = []
        self.hp = 0
        self.area = None
        self.grid = None  # Area=Board only. position in board.

    @staticmethod
    def create_unit_by_coin(coin: Coin):
        unit = Unit(coin.unit_type)
        unit.coins = [coin]
        coin.unit = unit
        unit.hp = 1
        return unit

    def add_coin(self, coin):
        self.coins.append(coin)
        self.hp += 1
        coin.area = self

    def __str__(self):
        return f"Unit type: {self.unit_type.name}\nHealth: {self.hp}\n"
