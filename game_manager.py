#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
from utils import singleton
from game_object import Board

@singleton
class GameManager(object):
    def __init__(self):
        self.board = Board()
        pass


    def game_start(self):
        Board.