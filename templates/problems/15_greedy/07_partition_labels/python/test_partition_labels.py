from partition_labels import partition_labels


def test_example_1():
    assert partition_labels("ababcbacadefegdehijhklij") == [9, 7, 8]


def test_single_char():
    assert partition_labels("a") == [1]


def test_all_same():
    assert partition_labels("aaaa") == [4]


def test_all_unique():
    assert partition_labels("abcdef") == [1, 1, 1, 1, 1, 1]


def test_two_partitions():
    assert partition_labels("aabbb") == [2, 3]


def test_example_2():
    assert partition_labels("eccbbbbdec") == [10]


def test_interleaved():
    assert partition_labels("abac") == [3, 1]
