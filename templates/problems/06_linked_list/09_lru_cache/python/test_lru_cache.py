from lru_cache import LRUCache


def _run(capacity, operations, keys, values, expected):
    cache = LRUCache(capacity)
    for i, op in enumerate(operations):
        if op == "put":
            cache.put(keys[i], values[i])
        elif op == "get":
            got = cache.get(keys[i])
            if expected[i] != -2:
                assert got == expected[i], (
                    f"operation {i}: Get({keys[i]}) = {got}, want {expected[i]}"
                )


def test_basic_get_and_put():
    _run(
        2,
        ["put", "put", "get", "put", "get", "put", "get", "get", "get"],
        [1, 2, 1, 3, 2, 4, 1, 3, 4],
        [1, 2, 0, 3, 0, 4, 0, 0, 0],
        [-2, -2, 1, -2, -1, -2, -1, 3, 4],
    )


def test_capacity_one():
    _run(
        1,
        ["put", "get", "put", "get", "get"],
        [1, 1, 2, 1, 2],
        [10, 0, 20, 0, 0],
        [-2, 10, -2, -1, 20],
    )


def test_update_existing_key():
    _run(
        2,
        ["put", "put", "get", "put", "get"],
        [1, 1, 1, 2, 1],
        [1, 2, 0, 3, 0],
        [-2, -2, 2, -2, 2],
    )


def test_get_miss():
    _run(
        2,
        ["get", "put", "get", "get"],
        [5, 5, 5, 10],
        [0, 50, 0, 0],
        [-1, -2, 50, -1],
    )


def test_eviction_order_with_get_refresh():
    _run(
        2,
        ["put", "put", "get", "put", "get"],
        [1, 2, 1, 3, 2],
        [1, 2, 0, 3, 0],
        [-2, -2, 1, -2, -1],
    )
