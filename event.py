#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------
class Event(object):
    def __init__(self):
        self.subscribers = []

    def notify(self, event):
        for subscriber in self.subscribers:
            subscriber(event)