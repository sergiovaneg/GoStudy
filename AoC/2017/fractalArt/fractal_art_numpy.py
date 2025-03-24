"""
Day 21, 2017, optimized using numpy vector functions
(Original in https://github.com/ednl/adventofcode/blob/main/2017/21.py)
"""

import sys

import numpy as np
from time import monotonic_ns

SEED = ".#./..#/###"

t0 = monotonic_ns()


def deserialize(serial: str) -> np.ndarray[bool]:
    n = serial.count("/") + 1
    return np.array(list(serial.replace("/", ""))).reshape((n, n)) == "#"


def get_equivalent_fractals(frac: np.ndarray[bool]) -> list[bytes]:
    perms = []
    for _ in range(4):
        perms.append(frac.tobytes())
        perms.append(np.fliplr(frac).tobytes())
        frac = np.rot90(frac)
    return perms


ruleset = {}
with open("./input.txt", "r", encoding=sys.getdefaultencoding()) as file:
    for line in file.readlines():
        line = line.strip()
        src, dst = [deserialize(serial) for serial in line.split(" => ")]
        rule_n = src.shape[0]
        for src_equiv in get_equivalent_fractals(src):
            ruleset[src_equiv] = dst


def evolve(seed: np.ndarray[bool]) -> np.ndarray[bool]:
    c_size = 2 if seed.shape[0] % 2 == 0 else 3
    sf_count = seed.shape[0] // c_size
    r_size = sf_count * (c_size + 1)

    return np.vectorize(
        lambda sf: ruleset[sf.tobytes()],
        signature="(i,j)->(k,l)"
    )(
        seed.reshape(
            (sf_count, c_size, sf_count, c_size)
        ).swapaxes(1, 2)
    ).swapaxes(1, 2).reshape((r_size, r_size))


fractal = deserialize(SEED)
print(0, np.sum(fractal, dtype=int))
for step in range(1, 19):
    fractal = evolve(fractal)
    print(step, np.sum(fractal, dtype=int))

t1 = monotonic_ns()
print(f"Time: {(t1 - t0) / 1E9: .3f} s")
