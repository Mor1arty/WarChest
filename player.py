#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
from game_object import Area


class PlayerProfile(object):
    def __init__(self):
        self.name = "Player1"
        self.faction = "Lion"

    def set_name(self, name):
        self.name = name


class Player(object):
    def __init__(self):
        self.profile = None
        self._unit_types = None
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
        self.hand.init_unit_types(new_unit_types)
        self.deck.init_unit_types(new_unit_types)
        self.supply.init_unit_types(new_unit_types)
        self.discard.init_unit_types(new_unit_types)
        self.graveyard.init_unit_types(new_unit_types)
