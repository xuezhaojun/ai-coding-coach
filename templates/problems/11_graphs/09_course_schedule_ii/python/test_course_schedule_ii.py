import pytest

from course_schedule_ii import find_order


@pytest.mark.parametrize("name, num_courses, prerequisites, want_len, has_order", [
    ("two courses with dependency", 2, [[1, 0]], 2, True),
    ("four courses", 4, [[1, 0], [2, 0], [3, 1], [3, 2]], 4, True),
    ("cycle returns empty", 2, [[1, 0], [0, 1]], 0, False),
    ("single course", 1, [], 1, True),
    ("no prerequisites", 3, [], 3, True),
    ("three node cycle", 3, [[0, 1], [1, 2], [2, 0]], 0, False),
])
def test_find_order(name, num_courses, prerequisites, want_len, has_order):
    got = find_order(num_courses, prerequisites)
    if not has_order:
        assert got == [], f"find_order() = {got}, want empty list for cycle"
        return
    assert len(got) == want_len, (
        f"find_order() returned {len(got)} courses, want {want_len}"
    )
    # verify ordering: for each prerequisite [a, b], b must appear before a
    pos = {c: i for i, c in enumerate(got)}
    for p in prerequisites:
        if pos[p[1]] > pos[p[0]]:
            pytest.fail(f"find_order(): prerequisite {p} violated in order {got}")
