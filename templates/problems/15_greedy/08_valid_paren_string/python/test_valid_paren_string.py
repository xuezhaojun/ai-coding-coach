from valid_paren_string import check_valid_string


def test_simple_valid():
    assert check_valid_string("()") == True


def test_star_as_empty():
    assert check_valid_string("(*)") == True


def test_star_as_paren():
    assert check_valid_string("(*))") == True


def test_empty_string():
    assert check_valid_string("") == True


def test_only_stars():
    assert check_valid_string("***") == True


def test_unmatched_open():
    assert check_valid_string("((") == False


def test_unmatched_close():
    assert check_valid_string("))") == False


def test_star_as_open():
    assert check_valid_string("*(") == False
