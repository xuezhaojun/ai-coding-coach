from multiply_strings import multiply


def test_example_1():
    assert multiply("2", "3") == "6"


def test_example_2():
    assert multiply("123", "456") == "56088"


def test_multiply_by_zero():
    assert multiply("0", "52") == "0"


def test_both_zeros():
    assert multiply("0", "0") == "0"


def test_single_digits():
    assert multiply("9", "9") == "81"


def test_large_numbers():
    assert multiply("999", "999") == "998001"


def test_one_and_number():
    assert multiply("1", "12345") == "12345"


def test_overflow_case():
    assert multiply("498828660196", "840477629533") == "419254329864656431168468"
