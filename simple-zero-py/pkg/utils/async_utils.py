# !/usr/bin/env python
# -*- coding: utf-8 -*-
import asyncio
import selectors


def async_run(feature):
    """
    python3.7-
    :param feature:
    :return:
    """
    # loop = asyncio._get_running_loop()
    # if loop is None:
    #     loop = asyncio.new_event_loop()
    #     asyncio.set_event_loop(loop)
    selector = selectors.SelectSelector()       # New line
    loop = asyncio.SelectorEventLoop(selector)
    return loop.run_until_complete(feature)


def get_loop():
    tmp_loop = asyncio._get_running_loop()
    if tmp_loop is None:
        tmp_loop = asyncio.new_event_loop()
        asyncio.set_event_loop(tmp_loop)
    return tmp_loop
