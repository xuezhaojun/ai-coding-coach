from string_lines import solve


def test_basic_reversal():
    assert solve("hello world\nI love coding\n") == "world hello\ncoding love I\n"


def test_four_words():
    assert solve("a b c d\n") == "d c b a\n"


def test_single_word():
    assert solve("hello\n") == "hello\n"


def test_multiple_spaces():
    assert solve("Go   is   great\n") == "great is Go\n"
