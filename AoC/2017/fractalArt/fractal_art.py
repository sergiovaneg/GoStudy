"""
Day 21, 2017
(Original in https://github.com/ednl/adventofcode/blob/main/2017/21.py)
"""

import sys

import numpy as np
from math import sqrt
from time import monotonic_ns

t0 = monotonic_ns()


def str2bit(s: str) -> str:
    return s.strip().replace("/", "").replace(".", "0").replace("#", "1")


def str2val(s: str) -> int:
    return int(str2bit(s), 2)


def str2arr(s: str) -> np.ndarray[np.uint8]:
    return (np.array([list(row)
            for row in s.strip().split("/")]) == "#").astype(np.uint8)


def arr2val(a: np.ndarray[np.uint8]) -> int:
    return int("".join(map(str, a.flatten())), 2)


def permvals(a: np.ndarray[np.uint8]) -> list[str]:
    p = set()
    p.add(arr2val(a))
    a = a.T
    p.add(arr2val(a))
    a = np.flipud(a)
    p.add(arr2val(a))
    a = a.T
    p.add(arr2val(a))
    a = np.flipud(a)
    p.add(arr2val(a))
    a = a.T
    p.add(arr2val(a))
    a = np.flipud(a)
    p.add(arr2val(a))
    a = a.T
    p.add(arr2val(a))
    return list(p)


rule = {2: {}, 3: {}}
with open("./input.txt", "r", encoding=sys.getdefaultencoding()) as f:
    for line in f:
        src, dst = line.split(" => ")
        f_size = 2 if len(src) == 5 else 3
        dst_val = str2val(dst)            # replacement
        for fractal in permvals(str2arr(src)):  # search patterns
            rule[f_size][fractal] = dst_val


def partition(bitstr: str, area: int, size: int, chunk: int) -> list[str]:
    step = size * chunk  # index step of 2 or 3 rows
    p = []
    for i in range(0, area, step):   # row index per 2 or 3 rows
        for j in range(0, size, chunk):  # col index per 2 or 3 cols
            s = ""
            for k in range(
                    0, step, size):  # extra row index for 2 or 3 conseq. rows
                n = i + j + k
                s += bitstr[n:n + chunk]
            p.append(int(s, 2))
    return p


def val2bit(val: int, partarea: int) -> str:
    s = bin(val)[2:]
    return "0" * (partarea - len(s)) + s


def evolve(bitstr: str) -> str:
    area = len(bitstr)
    size = int(sqrt(area))
    chunk = 2 if size % 2 == 0 else 3
    parts = partition(bitstr, area, size, chunk)
    transform = list(map(lambda x: rule[chunk][x], parts))
    partperrow = size // chunk
    parts = partperrow * partperrow
    chunk += 1
    size = partperrow * chunk
    area = size * size
    partarea = chunk * chunk
    a = list(map(lambda x: val2bit(x, partarea), transform))
    s = ""
    for i in range(0, parts, partperrow):
        for k in range(0, partarea, chunk):
            for j in range(partperrow):
                s += a[i + j][k:k + chunk]
    return s


image = ".#./..#/###"
art = str2bit(image)
print(0, sum(map(int, list(art))))
for f_size in range(1, 19):
    art = evolve(art)
    print(f_size, sum(map(int, list(art))))
# part 1: 179
# part 2: 2766750

t1 = monotonic_ns()
print(f"Time: {(t1 - t0) / 1E9: .3f} s")
