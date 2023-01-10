"""
This file is used only for assisting in the testing of the go package
"""

import itertools


def take(n, iterable):
    "Return first n items of the iterable as a list"
    return list(itertools.islice(iterable, n))


def n_count(start: int, step: int, n: int):
    for c in take(n, itertools.count(start, step)):
        print(c)


def accumulate(data: list[int]):
    for d in itertools.accumulate(data):
        print(d)


def pairwise(data: list[int]):
    for pair in itertools.pairwise(data):
        print(f"{pair[0]},{pair[1]}")
