import pytest

from trie import Trie


def _run(inserts, searches, prefixes):
    trie = Trie()
    for w in inserts:
        trie.insert(w)
    search_results = [(s[0], trie.search(s[0]), s[1]) for s in searches]
    prefix_results = [(p[0], trie.starts_with(p[0]), p[1]) for p in prefixes]
    return search_results, prefix_results


@pytest.mark.parametrize("name, inserts, searches, prefixes", [
    (
        "basic operations",
        ["apple"],
        [("apple", True), ("app", False), ("apples", False)],
        [("app", True), ("apple", True), ("b", False)],
    ),
    (
        "multiple words",
        ["apple", "app", "banana"],
        [("apple", True), ("app", True), ("banana", True), ("ban", False)],
        [("app", True), ("ban", True), ("cat", False)],
    ),
    (
        "empty string",
        [""],
        [("", True), ("a", False)],
        [("", True)],
    ),
    (
        "single char words",
        ["a", "b", "c"],
        [("a", True), ("b", True), ("d", False)],
        [("a", True), ("d", False)],
    ),
    (
        "overlapping words",
        ["the", "then", "them", "there"],
        [("the", True), ("then", True), ("they", False), ("there", True)],
        [("the", True), ("th", True), ("thx", False)],
    ),
])
def test_trie(name, inserts, searches, prefixes):
    search_results, prefix_results = _run(inserts, searches, prefixes)
    for word, got, want in search_results:
        assert got == want, f"Search({word!r}) = {got}, want {want}"
    for prefix, got, want in prefix_results:
        assert got == want, f"StartsWith({prefix!r}) = {got}, want {want}"
