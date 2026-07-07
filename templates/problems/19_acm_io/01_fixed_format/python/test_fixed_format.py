from fixed_format import solve


def test_basic_sum():
    assert solve("5\n1 2 3 4 5\n") == "15\n"


def test_negative_numbers():
    assert solve("3\n-1 0 1\n") == "0\n"


def test_single_element():
    assert solve("1\n42\n") == "42\n"


def test_large_values():
    assert solve("2\n1000000000 -1000000000\n") == "0\n"
