from time_based_kv import TimeMap


def test_basic_set_and_get():
    tm = TimeMap()
    tm.set("foo", "bar", 1)
    assert tm.get("foo", 1) == "bar"
    assert tm.get("foo", 3) == "bar"
    tm.set("foo", "bar2", 4)
    assert tm.get("foo", 4) == "bar2"
    assert tm.get("foo", 5) == "bar2"


def test_get_before_any_set():
    tm = TimeMap()
    assert tm.get("foo", 1) == ""


def test_get_with_timestamp_before_first_set():
    tm = TimeMap()
    tm.set("key", "val", 5)
    assert tm.get("key", 3) == ""


def test_multiple_keys():
    tm = TimeMap()
    tm.set("a", "v1", 1)
    tm.set("b", "v2", 1)
    assert tm.get("a", 1) == "v1"
    assert tm.get("b", 1) == "v2"


def test_get_exact_timestamp():
    tm = TimeMap()
    tm.set("k", "a", 1)
    tm.set("k", "b", 2)
    tm.set("k", "c", 3)
    assert tm.get("k", 1) == "a"
    assert tm.get("k", 2) == "b"
    assert tm.get("k", 3) == "c"
