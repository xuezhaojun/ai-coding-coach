from encode_decode_strings import encode, decode


def test_basic_strings():
    strs = ["hello", "world"]
    assert decode(encode(strs)) == strs


def test_empty_list():
    strs: list[str] = []
    assert decode(encode(strs)) == strs


def test_single_empty_string():
    strs = [""]
    assert decode(encode(strs)) == strs


def test_multiple_empty_strings():
    strs = ["", "", ""]
    assert decode(encode(strs)) == strs


def test_strings_with_special_characters():
    strs = ["he:llo", "wor#ld", "foo;bar"]
    assert decode(encode(strs)) == strs


def test_strings_with_delimiters_and_colons():
    strs = ["4:abcd", "3:xyz"]
    assert decode(encode(strs)) == strs


def test_single_character_strings():
    strs = ["a", "b", "c"]
    assert decode(encode(strs)) == strs


def test_mixed_empty_and_non_empty():
    strs = ["", "a", "", "b", ""]
    assert decode(encode(strs)) == strs
