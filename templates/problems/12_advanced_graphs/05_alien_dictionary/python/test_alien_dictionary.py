import pytest

from alien_dictionary import alien_order


def _verify(words: list[str], got: str, want: str) -> None:
    if want == "":
        assert got == "", f"alien_order({words}) = {got!r}, want empty string"
        return
    assert len(got) == len(want), (
        f"alien_order({words}) = {got!r} (len {len(got)}), want len {len(want)}"
    )
    char_pos = {ch: i for i, ch in enumerate(got)}
    for i in range(len(words) - 1):
        w1, w2 = words[i], words[i + 1]
        for j in range(min(len(w1), len(w2))):
            if w1[j] != w2[j]:
                assert char_pos[w1[j]] <= char_pos[w2[j]], (
                    f"alien_order({words}) = {got!r} violates ordering of "
                    f"{words[i]!r} before {words[i + 1]!r}"
                )
                break


@pytest.mark.parametrize("name, words, want", [
    ("standard order", ["wrt", "wrf", "er", "ett", "rftt"], "wertf"),
    ("simple two words", ["z", "x"], "zx"),
    ("invalid order", ["z", "x", "z"], ""),
    ("single word", ["abc"], "abc"),
    ("prefix violation", ["abc", "ab"], ""),
    ("single characters", ["z", "z"], "z"),
])
def test_alien_order(name, words, want):
    got = alien_order(words)
    _verify(words, got, want)
