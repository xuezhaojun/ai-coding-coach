from sentinel_terminate import solve


def test_basic_with_sentinel():
    assert solve("1 2\n3 4\n0 0\n") == "3\n7\n"


def test_immediate_sentinel():
    assert solve("0 0\n") == ""


def test_negative_before_sentinel():
    assert solve("-5 10\n0 0\n") == "5\n"


def test_zeros_not_at_sentinel_position():
    assert solve("0 5\n5 0\n0 0\n") == "5\n5\n"
