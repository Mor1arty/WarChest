#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
from enum import Enum


class CoinType(Enum):
    """
    Image path of unit coins.(Default in folder imgs/)
    """
    UNKNOWN = 0
    PHOENIX = "phoenix"
    LION = "lion"
    OTHERS = 1
    LIGHT_CAVALRY = "light_cavalry"
    ARCHER = "archer"
    BERSERKER = "berserker"
    CROSSBOWMAN = "crossbowman"
    ENSIGN = "ensign"
    FOOTMAN = "footman"
    LANCER = "lancer"
    MARSHALL = "marshall"
    MERCENARY = 8
    PIKEMAN = 9
    WARRIOR_PRIEST = 10


class UnitCard(object):
    def __init__(self, unit_type: CoinType = CoinType.UNKNOWN):
        self.unit_type: CoinType = unit_type  # 兵种
        self.deploy_limit = 1
        self.coin_limit = 5  # TODO read config

    def __str__(self):
        return f"UnitCard {self.unit_type}"


class Coin(UnitCard):
    def __init__(self, unit_type):
        super().__init__(unit_type)
        self.unit = None
        self.is_revealed = True

    @property
    def area(self):
        return self.unit.area

    def __str__(self):
        return f"Coin {self.unit_type}"


class Unit(UnitCard):
    def __init__(self, unit_type: CoinType):
        super().__init__(unit_type)
        self.coins: list = []
        self.hp = 0  # how many coin in this unit
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
        # coin.area = self

    def pop_coin(self, coin=None) -> Coin:
        self.hp -= 1
        if self.unit_type != CoinType.OTHERS:
            return self.coins.pop()
        else:
            self.coins.remove(coin)
            return coin

    def pop_coin_by_type(self, unit_type=None) -> Coin:
        self.hp -= 1
        if self.unit_type != CoinType.OTHERS:
            return self.coins.pop()
        else:
            for coin in self.coins:
                if coin.unit_type == unit_type:
                    self.coins.remove(coin)
                    return coin
        raise

    def __str__(self):
        return f"Unit type: {self.unit_type.name}\nCounts: {self.hp}\n"
