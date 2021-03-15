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
import logging


@singleton
class GameManager(object):
    def __init__(self):
        self.board = Board()
        self.coins_draw_per_turn_num = 3
        self.players = []
        self.initiative_player = None
        self.is_initiative_change_this_turn = False
        pass

    def game_start(self):
        self.board.init_1v1()
        EventManager().notify(EventType.GAME_START)

    def players_generator(self):
        if self.initiative_player not in self.players:
            logging.error(f"{self.initiative_player} is not in {self.players}!")
        index = 0
        while self.players[index] != self.initiative_player:
            index += 1
        yield self.players[index]
        index += 1
        if index == len(self.players):
            index = 0
        while self.players[index] != self.initiative_player:
            yield self.players[index]

    def round_start(self):
        for player in self.players:
            self.draw_coins(player, min(self.coins_draw_per_turn_num, player.deck.size))
        for player in self.players_generator():
            self.take_turn(player)
        pass

    def draw_coins(self, players, num=1):
        pass

    def take_turn(self, player):
        pass

    def try_deploy(self, player: Player, coin, pos) -> bool:
        # check validation
        if coin.unit_type not in player.unit_types:
            return False
        if coin.UIBoard.unit_num(coin.unit_type) >= coin.deploy_limit:
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
