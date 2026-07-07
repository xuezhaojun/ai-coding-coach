from letter_combinations import letter_combinations


def test_example_23():
    got = letter_combinations("23")
    expected = ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
    got.sort()
    expected.sort()
    assert got == expected


def test_empty_string():
    got = letter_combinations("")
    expected: list[str] = []
    got.sort()
    expected.sort()
    assert got == expected


def test_single_digit_2():
    got = letter_combinations("2")
    expected = ["a", "b", "c"]
    got.sort()
    expected.sort()
    assert got == expected


def test_digit_7_with_four_letters():
    got = letter_combinations("7")
    expected = ["p", "q", "r", "s"]
    got.sort()
    expected.sort()
    assert got == expected


def test_digit_9_with_four_letters():
    got = letter_combinations("9")
    expected = ["w", "x", "y", "z"]
    got.sort()
    expected.sort()
    assert got == expected


def test_three_digits():
    got = letter_combinations("234")
    expected = [
        "adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
        "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
        "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
    ]
    got.sort()
    expected.sort()
    assert got == expected
