#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
class PlayerProfile(object):
    def __init__(self):
        self.name = "Player1"

    def set_name(self, name):
        self.name = name


class Player(object):
    def __init__(self):
        self.profile = None
        self.unit_types = None
        self.hand = None
        self.supply = None
        self.discard = None
        self.graveyard = None
