import pytest

from word_dictionary import WordDictionary


def _run(adds, searches):
    wd = WordDictionary()
    for w in adds:
        wd.add_word(w)
    return [(s[0], wd.search(s[0]), s[1]) for s in searches]


@pytest.mark.parametrize("name, adds, searches", [
    (
        "basic with wildcards",
        ["bad", "dad", "mad"],
        [("pad", False), ("bad", True), (".ad", True), ("b..", True), ("...", True), ("....", False)],
    ),
    (
        "exact match only",
        ["hello"],
        [("hello", True), ("hell", False), ("helloo", False)],
    ),
    (
        "all wildcards",
        ["ab", "cd"],
        [("..", True), (".", False), ("...", False)],
    ),
    (
        "single char",
        ["a"],
        [(".", True), ("a", True), ("..", False)],
    ),
    (
        "mixed wildcards",
        ["apple", "ample"],
        [("a.ple", True), ("a..le", True), ("a...e", True), ("a....", True), ("b....", False)],
    ),
    (
        "no words added",
        [],
        [("a", False), (".", False)],
    ),
])
def test_word_dictionary(name, adds, searches):
    results = _run(adds, searches)
    for word, got, want in results:
        assert got == want, f"Search({word!r}) = {got}, want {want}"
