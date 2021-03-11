#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
from utils import singleton
from game_object import Board, Terrain, TerrainType
from player import Player
from event import Event, EventType, EventManager


@singleton
class GameManager(object):
    def __init__(self):
        self.board = Board()
        pass

    def game_start(self):
        self.board.init_1v1()
        EventManager().notify(EventType.GAME_START)

    def try_deploy(self, player: Player, coin, pos) -> bool:
        # check validation
        if coin.unit_type not in player.unit_types:
            return False
        if coin.Board.unit_num(coin.unit_type) >= coin.deploy_limit:
            return False
        if coin not in player.hand:
            return False

        self.deploy(coin, pos)
        return True

    def deploy(self, coin, pos):
        self.board.add_coin(coin, pos)
        EventManager().notify(EventType.DEPLOY_AFTER, event=Event(coin=coin, target=pos))

    def attack(self, attacker, defender):
        pass
