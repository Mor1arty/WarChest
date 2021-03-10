#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
from enum import Enum
import logging
from utils import singleton


class EventType(Enum):
    GAME_START = 1
    DRAW_AFTER = 2
    GAME_END = 5

    ACTION_AFTER = 12
    DEPLOY_AFTER = 13
    BOLSTER_AFTER = 14
    MOVE_AFTER = 15
    ATTACK_AFTER = 16
    CONTROL_AFTER = 17
    TACTIC_AFTER = 18
    CLAIM_INITIATIVE_AFTER = 19
    RECRUIT_AFTER = 20
    PASS_AFTER = 21

    PROCLAIM_AFTER = 33


@singleton
class EventManager(object):
    def __init__(self):
        self.events = {}

    def register(self, event_type, subscriber):
        event = self.events.get(event_type)
        if event is None:
            self.events[event_type] = Event(event_type)
            event = self.events[event_type]
        event.register(subscriber)

    def unregister(self, event_type, subscriber):
        event = self.events.get(event_type)
        if event is None:
            logging.error("No Such Event!")
        event.unregister(subscriber)


class Event(object):
    def __init__(self, event_type=None):
        self.event_type = event_type
        self.subscribers = []

    def register(self, subscriber):
        self.subscribers.append(subscriber)

    def unregister(self, subscriber):
        if subscriber not in self.subscribers:
            logging.error(f"{subscriber} is not in {self.event_type}")
        self.subscribers.remove(subscriber)

    def notify(self, event):
        for subscriber in self.subscribers:
            subscriber(event)
