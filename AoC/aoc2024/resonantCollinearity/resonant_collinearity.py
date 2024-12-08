"""
Functional implementation of the Go solution
"""

import sys
from itertools import combinations, groupby
from functools import partial, reduce


with open("./input.txt", "r", encoding=sys.getdefaultencoding()) as f:
  rows = f.readlines()

m, n = len(rows), len(rows[0]) - 1


def validate_coord(x, y):
  return not (x < 0 or y < 0 or x >= m or y >= n)


def process_pair(xy0, xy1, lb, ub):
  res = []

  for p0, p1 in [[xy0, xy1], [xy1, xy0]]:
    d = [c1 - c0 for c0, c1 in zip(p0, p1)]
    k = lb
    while k != ub:
      aux = [c + k * cd for c, cd in zip(p1, d)]
      if validate_coord(aux[0], aux[1]):
        res.append(aux)
      else:
        break
      k += 1

  return res


def process_group(antennae, lb, ub):
  res = [
      tuple(e)
      for cxy0, cxy1 in combinations(antennae, 2)
      for e in process_pair(cxy0[1:], cxy1[1:], lb, ub)
  ]
  return set(res)


grouped_antennae = [
    list(x) for _, x in groupby(
        sorted(
            [
              (c, i, j)
                for i, j_row in enumerate(map(enumerate, rows))
                for j, c in j_row
                if c not in [".", "\n"]
            ],
            key=lambda x: x[0]
        ),
        lambda x: x[0]
    )
]

for processor in [
    partial(process_group, lb=1, ub=2),
    partial(process_group, lb=0, ub=-1)
]:
  mapped_antennae = map(processor, grouped_antennae)
  print(len(reduce(set.union, mapped_antennae)))
