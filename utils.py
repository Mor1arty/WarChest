#!/usr/bin/env python
# coding: utf-8
# ----------------------------------------------------------------------
# Author: arshart@forevernine.com
# Description: 
# ----------------------------------------------------------------------


def singleton(cls):
    """单例装饰器函数"""
    _instance = {}

    def inner(*args, **kw):
        if cls not in _instance:
            _instance[cls] = cls(*args, **kw)
        return _instance[cls]
    return inner
