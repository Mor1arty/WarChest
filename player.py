#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
from game_object import Area
from unit import UnitCard, Coin, CoinType


class PlayerProfile(object):
    def __init__(self):
        self.name = "Player1"
        self.faction = "Lion"

    def set_name(self, name):
        self.name = name


class Player(object):
    def __init__(self, fiction_type=CoinType.UNKNOWN):
        self.profile = None  # TODO , uid, player name, and so on
        self.fiction_type: CoinType = fiction_type
        self._unit_types: list = []
        self.hand: Area = Area()
        self.deck: Area = Area()
        self.supply: Area = Area()
        self.discard: Area = Area()
        self.graveyard: Area = Area()

    @property
    def unit_types(self):
        return self._unit_types

    @unit_types.setter
    def unit_types(self, new_unit_types):
        self._unit_types = new_unit_types
        self.init_supply()
        self.init_deck()
        self.hand.init_unit_types(new_unit_types)
        self.discard.init_unit_types(new_unit_types)
        self.graveyard.init_unit_types(new_unit_types)

    def init_supply(self):
        self.supply.init_unit_types(self.unit_types)
        for unit_type in self.unit_types:
            uc = UnitCard(unit_type)
            for _ in range(uc.coin_limit):
                self.supply.add_coin(Coin(unit_type))

    def init_deck(self):
        self.deck.init_unit_types(self.unit_types)
        # add 2 coins of each unit from supply to deck.
        for unit_type in self.unit_types:
            for _ in range(2):
                coin = self.supply.pop_coin_by_type(unit_type)
                self.deck.add_coin(coin)
        self.deck.add_coin(Coin(self.fiction_type))  # Add royal coin.