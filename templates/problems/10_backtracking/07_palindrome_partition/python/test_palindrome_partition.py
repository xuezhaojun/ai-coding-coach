from palindrome_partition import partition


def sort_string_slices(result: list[list[str]]) -> list[list[str]]:
    for s in result:
        s.sort()
    result.sort(key=lambda g: (len(g), g))
    return result


def test_example_aab():
    got = partition("aab")
    expected = [["a", "a", "b"], ["aa", "b"]]
    assert sort_string_slices(got) == sort_string_slices(expected)


def test_single_char():
    got = partition("a")
    expected = [["a"]]
    assert sort_string_slices(got) == sort_string_slices(expected)


def test_all_same_chars():
    got = partition("aaa")
    expected = [["a", "a", "a"], ["a", "aa"], ["aa", "a"], ["aaa"]]
    assert sort_string_slices(got) == sort_string_slices(expected)


def test_no_palindrome_partitions_beyond_singles():
    got = partition("abc")
    expected = [["a", "b", "c"]]
    assert sort_string_slices(got) == sort_string_slices(expected)


def test_full_palindrome():
    got = partition("aba")
    expected = [["a", "b", "a"], ["aba"]]
    assert sort_string_slices(got) == sort_string_slices(expected)


def test_two_chars_same():
    got = partition("bb")
    expected = [["b", "b"], ["bb"]]
    assert sort_string_slices(got) == sort_string_slices(expected)
