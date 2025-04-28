"""
Script to solve part 2 of day 13, 2020.
"""

import sys
import cvxpy
import numpy as np

with open("./input.txt", "r", encoding=sys.getdefaultencoding()) as f:
    buses = [int(x) if x != "x" else None for x in f.readlines()[-1].split(",")]

mask = np.array([x is not None for x in buses], dtype=bool)
n_buses = np.count_nonzero(mask)


a_mat = np.diag([x for x in buses if x is not None])
b_vec = np.array([
    idx for idx, x in enumerate(buses)
    if x is not None
], dtype=float)


c_vec = np.eye(n_buses + 1, 1).flatten()
ts = cvxpy.Variable()
x = cvxpy.Variable(n_buses, integer=True)

prob = cvxpy.Problem(
    cvxpy.Minimize(ts),
    [
        a_mat @ x == ts + b_vec,
        x >= 0,
        ts >= 0
    ]
)

prob.solve(
    verbose=False,
    solver="GUROBI",
    reoptimize=True
)

print(ts.value)
