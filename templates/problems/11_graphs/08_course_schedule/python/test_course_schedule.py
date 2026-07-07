import pytest

from course_schedule import can_finish


@pytest.mark.parametrize("name, num_courses, prerequisites, want", [
    ("simple chain", 2, [[1, 0]], True),
    ("cycle detected", 2, [[1, 0], [0, 1]], False),
    ("no prerequisites", 3, [], True),
    ("diamond dependency", 4, [[1, 0], [2, 0], [3, 1], [3, 2]], True),
    ("three node cycle", 3, [[0, 1], [1, 2], [2, 0]], False),
    ("single course", 1, [], True),
    ("disconnected courses", 5, [[1, 0], [3, 2]], True),
    ("node with two prerequisites and cycle", 3, [[1, 0], [1, 2], [0, 1]], False),
])
def test_can_finish(name, num_courses, prerequisites, want):
    got = can_finish(num_courses, prerequisites)
    assert got == want
