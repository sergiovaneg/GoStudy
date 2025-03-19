"""
Standard Library version of Bernal's solution
"""

import sys
from itertools import groupby


def get_sorted_values(raw_data: list[str]) -> tuple[list[int], list[int]]:
    parsed_data = [
        [int(num) for num in line.split("   ", 2)]
        for line in raw_data
    ]

    x, y = zip(*parsed_data)
    return sorted(x), sorted(y)


def i_hate_santa(x: list[int], y: list[int]) -> int:
    return sum([abs(a - b) for a, b in zip(x, y)])


def agh(x: list[int], y: list[int]):
    # Only works if already sorted
    x_grouped = {k: len(list(v)) for k, v in groupby(x)}
    y_grouped = {k: len(list(v)) for k, v in groupby(y)}

    return sum(
        [
            k * x_grouped[k] * y_grouped.get(k, 0) for k in x_grouped
        ]
    )


if __name__ == "__main__":
    with open("./input.txt", "r", encoding=sys.getdefaultencoding()) as f:
        data = f.readlines()

    left, right = get_sorted_values(data)

    print(i_hate_santa(left, right))
    print(agh(left, right))
