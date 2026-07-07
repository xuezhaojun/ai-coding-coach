import pytest

from car_fleet import car_fleet


@pytest.mark.parametrize(
    "target,position,speed,want",
    [
        (12, [10, 8, 0, 5, 3], [2, 4, 1, 1, 3], 3),
        (10, [3], [3], 1),
        (10, [0, 2, 4], [2, 2, 2], 3),
        (10, [0, 2], [2, 1], 1),
        (10, [], [], 0),
        (100, [0, 50], [1, 1], 2),
    ],
    ids=[
        "example_1",
        "single_car",
        "all_same_speed",
        "all_merge",
        "no_cars",
        "two_separate_fleets",
    ],
)
def test_car_fleet(
    target: int, position: list[int], speed: list[int], want: int
) -> None:
    assert car_fleet(target, position, speed) == want
