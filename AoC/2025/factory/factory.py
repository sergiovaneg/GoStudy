"""
Day 10, 2025: Factory
"""

import sys
import re
import numpy as np
from scipy import optimize


def parse_machine(machine: str) -> tuple[np.ndarray, np.ndarray]:
    joltages = np.asarray(
        [
            float(num)
            for num in re.findall(r"\{.*\}", machine)[0][1:-1].split(",")
        ]
    )

    buttons = re.findall(r"\([^\)]*\)", machine)

    mask = np.zeros([len(joltages), len(buttons)])
    for j, button in enumerate(buttons):
        idxs = [int(x) for x in button[1:-1].split(",")]
        mask[idxs, j] = 1.

    return mask, joltages


def optimize_pushes(mask: np.ndarray, joltages: np.ndarray) -> int:
    x = optimize.linprog(
        c=np.ones(mask.shape[1]),
        A_eq=mask,
        b_eq=joltages,
        bounds=(0, np.sum(joltages)),
        integrality=1
    ).x

    return int(np.sum(x))


total = 0
with open("input.txt", "r", encoding=sys.getdefaultencoding()) as f:
    for line in f.readlines():
        total += optimize_pushes(*parse_machine(line))

print(total)
