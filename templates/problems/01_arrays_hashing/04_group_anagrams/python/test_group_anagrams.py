from group_anagrams import group_anagrams


def sort_groups(groups: list[list[str]]) -> list[list[str]]:
    for group in groups:
        group.sort()
    groups.sort(key=lambda g: (len(g), g))
    return groups


def test_mixed_anagram_groups():
    result = group_anagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
    assert sort_groups(result) == sort_groups([["bat"], ["nat", "tan"], ["ate", "eat", "tea"]])


def test_single_empty_string():
    result = group_anagrams([""])
    assert sort_groups(result) == sort_groups([[""]])


def test_single_non_empty_string():
    result = group_anagrams(["a"])
    assert sort_groups(result) == sort_groups([["a"]])


def test_no_anagrams():
    result = group_anagrams(["abc", "def", "ghi"])
    assert sort_groups(result) == sort_groups([["abc"], ["def"], ["ghi"]])


def test_all_anagrams():
    result = group_anagrams(["abc", "bca", "cab"])
    assert sort_groups(result) == sort_groups([["abc", "bca", "cab"]])


def test_empty_input():
    result = group_anagrams([])
    assert result == []


def test_multiple_empty_strings():
    result = group_anagrams(["", ""])
    assert sort_groups(result) == sort_groups([["", ""]])
