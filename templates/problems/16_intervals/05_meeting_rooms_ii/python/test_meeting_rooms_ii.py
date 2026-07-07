from meeting_rooms_ii import min_meeting_rooms


def test_two_rooms():
    assert min_meeting_rooms([[0, 30], [5, 10], [15, 20]]) == 2


def test_one_room():
    assert min_meeting_rooms([[7, 10], [2, 4]]) == 1


def test_all_overlap():
    assert min_meeting_rooms([[1, 5], [2, 6], [3, 7]]) == 3


def test_single_meeting():
    assert min_meeting_rooms([[1, 10]]) == 1


def test_sequential():
    assert min_meeting_rooms([[0, 5], [5, 10], [10, 15]]) == 1


def test_nested():
    assert min_meeting_rooms([[1, 10], [2, 5], [6, 9]]) == 2


def test_empty():
    assert min_meeting_rooms([]) == 0
