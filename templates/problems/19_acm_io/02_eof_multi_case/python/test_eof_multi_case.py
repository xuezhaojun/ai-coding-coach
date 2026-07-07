from eof_multi_case import solve


def test_basic_multi_case():
    assert solve("1 2\n3 4\n5 6\n") == "3\n7\n11\n"


def test_negative_numbers():
    assert solve("-1 2\n0 0\n") == "1\n0\n"


def test_single_case():
    assert solve("100 200\n") == "300\n"


def test_large_values():
    assert solve("1000000000 1000000000\n-1000000000 1\n") == "2000000000\n-999999999\n"
