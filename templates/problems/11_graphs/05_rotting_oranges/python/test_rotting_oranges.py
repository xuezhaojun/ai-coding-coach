import copy
import pytest

from rotting_oranges import oranges_rotting


@pytest.mark.parametrize("name, grid, want", [
    (
        "standard case",
        [
            [2, 1, 1],
            [1, 1, 0],
            [0, 1, 1],
        ],
        4,
    ),
    (
        "impossible to rot all",
        [
            [2, 1, 1],
            [0, 1, 1],
            [1, 0, 1],
        ],
        -1,
    ),
    ("no fresh oranges", [[0, 2]], 0),
    ("empty grid cells only", [[0]], 0),
    ("already all rotten", [[2, 2], [2, 2]], 0),
    ("single fresh unreachable", [[1]], -1),
    ("one step", [[2, 1]], 1),
    (
        "multiple sources infect same round",
        [
            [2, 2],
            [1, 1],
            [0, 0],
            [2, 0],
        ],
        1,
    ),
])
def test_oranges_rotting(name, grid, want):
    g = copy.deepcopy(grid)
    got = oranges_rotting(g)
    assert got == want
