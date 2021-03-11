#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
from enum import Enum
import logging
from utils import singleton


class Event(object):
    """
    Deploy: target: pos
    Bolster: target: Unit

    Move: source: Unit, target: pos
    Attack: source: Unit, target: Unit
    Control: source: Unit

    Claim Initiative:
    Recruit: target: Coin
    Pass: source:
    """
    def __init__(self, source=None, target=None, coin=None, player=None):
        self.source = source
        self.target = target
        self.coin = coin
        self.player = player


class EventType(Enum):
    GAME_START = 1
    DRAW_AFTER = 2
    GAME_END = 5
    ACTION_AFTER = 12
    # Placement action
    DEPLOY_AFTER = 13
    BOLSTER_AFTER = 14
    # Face up to discard(Maneuvers)
    MOVE_AFTER = 15
    ATTACK_AFTER = 16
    CONTROL_AFTER = 17
    TACTIC_AFTER = 18
    # Face down to discard
    CLAIM_INITIATIVE_AFTER = 19
    RECRUIT_AFTER = 20
    PASS_AFTER = 21

    PROCLAIM_AFTER = 33


@singleton
class EventManager(object):
    def __init__(self):
        self.events = {}  # key: EventType. Value: subscribers (list)

    def register(self, event_type, subscriber):
        subscribers = self.events.get(event_type)
        if subscribers is None:
            self.events[event_type] = []
            subscribers = self.events[event_type]
        subscribers.append(subscriber)

    def unregister(self, event_type, subscriber):
        subscribers = self.events.get(event_type)
        if subscribers is None:
            logging.error("No Such Event!")
            return
        if subscriber not in subscribers:
            logging.error(f"{subscriber} is not in {event_type}")
            return
        subscribers.remove(subscriber)

    def notify(self, event_type, event=None):
        subscribers = self.events.get(event_type)
        if subscribers is None:
            logging.error("No Such Event!")
            return
        for subscriber in subscribers:
            subscriber(event)
