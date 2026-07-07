from graph_input import solve


def test_basic_directed_graph():
    assert solve("4 5\n1 2\n1 3\n2 3\n2 4\n3 4\n") == "2\n2\n1\n0\n"


def test_sparse_graph():
    assert solve("3 2\n1 2\n3 2\n") == "1\n0\n1\n"


def test_single_edge():
    assert solve("2 1\n1 2\n") == "1\n0\n"


def test_self_loop():
    assert solve("2 2\n1 1\n1 2\n") == "2\n0\n"
