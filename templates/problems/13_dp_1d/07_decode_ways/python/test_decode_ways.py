from decode_ways import num_decodings


def test_example_12():
    assert num_decodings("12") == 2


def test_example_226():
    assert num_decodings("226") == 3


def test_leading_zero():
    assert num_decodings("06") == 0


def test_single_digit():
    assert num_decodings("1") == 1


def test_example_11106():
    assert num_decodings("11106") == 2


def test_all_ones():
    assert num_decodings("1111") == 5


def test_example_10():
    assert num_decodings("10") == 1


def test_example_27():
    assert num_decodings("27") == 1
