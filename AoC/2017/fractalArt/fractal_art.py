"""
Script to generate the figures related to Fractal Art (AoC 2017, D21).
"""

import os

import numpy as np

from matplotlib.axis import Axis
import matplotlib.pyplot as plt
import seaborn as sns

sns.set_theme(
    context="paper",
    style="dark",
    palette="Dark2",
    font="roboto",
    font_scale=2
)

os.makedirs("./figures/", exist_ok=True)


def serialize(f: list[list[str]]) -> str:
  return "/".join([x for x in ["".join(e) for e in f]])


def deserialize(s: str) -> list[list[str]]:
  return [list(x) for x in s.split("/")]


def mirror(f: list[list[str]]) -> list[list[str]]:
  return [list(reversed(x)) for x in f]


def rotate(f: list[list[str]]) -> list[list[str]]:
  return [[f[j][i] for j in reversed(range(3))] for i in range(3)]


def deserialize_boolean(s: str) -> np.ndarray:
  return [[True if e == "#" else False for e in x] for x in s.split("/")]


def generate_partitioned(n: int, n_sub: int) -> np.ndarray:
  ratio = n // n_sub
  return np.arange(
      ratio**2
  ).reshape(
      [ratio, ratio]
  ).repeat(n_sub, 0).repeat(n_sub, 1)


def plot_fractal(data: np.ndarray, ax: Axis):
  sns.heatmap(
      data,
      ax=ax,
      cbar=False,
      xticklabels=False,
      yticklabels=False,
      annot=True
  )


# Seed
fig, axs = plt.subplots(1, 1, figsize=(6, 6))
plot_fractal(deserialize_boolean(".#./..#/###"), axs)
fig.tight_layout()
fig.savefig("./figures/seed.png")

# Rules
for idx, rule in enumerate([
    "../.# => ##./#../...",
    ".#./..#/### => #..#/..../..../#..#"
]):
  fig, axs = plt.subplots(1, 2, figsize=(12, 6))
  src, dst = rule.split(" => ")

  plot_fractal(deserialize_boolean(src), axs[0])
  plot_fractal(deserialize_boolean(dst), axs[1])
  axs[0].set_title("Input")
  axs[1].set_title("Output")

  fig.tight_layout()

  fig.savefig(f"./figures/rule_{idx}.png")

# Invariants
f_ex = deserialize(".#./..#/###")
f_ex_m = mirror(f_ex)

fig, axs = plt.subplots(2, 4, figsize=(12, 6))
for idx in range(4):
  plot_fractal(
      deserialize_boolean(serialize(f_ex)),
      axs[0][idx]
  )
  plot_fractal(
      deserialize_boolean(serialize(f_ex_m)),
      axs[1][idx]
  )
  f_ex = rotate(f_ex)
  f_ex_m = rotate(f_ex_m)

fig.suptitle("Equivalent Fractals")
fig.tight_layout()
fig.savefig("./figures/equivalent.png")

# Loop
fig, axs = plt.subplots(2, 4, figsize=(24, 12))
for gen, i_shape, i_subshape, o_shape, o_subshape in [
    [0, 3, 3, 4, 4],
    [1, 4, 2, 6, 3],
    [2, 6, 2, 9, 3],
    [3, 9, 3, 12, 4]
]:
  plot_fractal(
      generate_partitioned(i_shape, i_subshape),
      axs[0][gen]
  )
  plot_fractal(
      generate_partitioned(o_shape, o_subshape),
      axs[1][gen]
  )
  axs[1][gen].set_xlabel(f"Gen. {gen}")

fig.suptitle("Iteration Behaviour")
axs[0][0].set_ylabel("Input")
axs[1][0].set_ylabel("Output")
fig.tight_layout()
fig.savefig("./figures/evo.png")
