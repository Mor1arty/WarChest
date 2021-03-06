#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description:
# ----------------------------------------------------------------------


class Coin(object):
    def __init__(self, unit_type):
        self.unit_type = unit_type  # 兵种


class Terrain (object):
    def __init__(self, terrain):
        self.terrain = terrain


class Unit(object):
    def __init__(self, coin):
        self.coin = coin
        self.hp = 1
        self.unit_type = coin.unit_type


class Grid(object):
    def __init__(self):
        self.terrain = None
        self.unit = None


class Area(object):
    def __init__(self):
        self.coins = []


class Board(object):
    def __init__(self):
        self.grids = None
        self.init_1v1()

    def init_1v1(self):
        cols = [4, 5, 6, 7, 6, 5, 4]
        self.grids = [[Grid() for i in range(cols[row])] for row in range(len(cols))]
