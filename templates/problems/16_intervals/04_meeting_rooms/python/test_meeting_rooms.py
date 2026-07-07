from meeting_rooms import can_attend_meetings


def test_overlapping():
    assert can_attend_meetings([[0, 30], [5, 10], [15, 20]]) == False


def test_no_overlap():
    assert can_attend_meetings([[7, 10], [2, 4]]) == True


def test_empty():
    assert can_attend_meetings([]) == True


def test_single_meeting():
    assert can_attend_meetings([[1, 5]]) == True


def test_touching_boundaries():
    assert can_attend_meetings([[1, 5], [5, 10]]) == True


def test_same_time():
    assert can_attend_meetings([[1, 5], [1, 5]]) == False


def test_sequential():
    assert can_attend_meetings([[0, 1], [1, 2], [2, 3]]) == True
