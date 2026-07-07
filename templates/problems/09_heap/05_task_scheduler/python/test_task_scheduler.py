from task_scheduler import least_interval


def test_example_1():
    assert least_interval(["A", "A", "A", "B", "B", "B"], 2) == 8


def test_no_cooldown():
    assert least_interval(["A", "A", "A", "B", "B", "B"], 0) == 6


def test_example_3():
    assert (
        least_interval(
            ["A", "A", "A", "A", "A", "A", "B", "C", "D", "E", "F", "G"], 2
        )
        == 16
    )


def test_single_task():
    assert least_interval(["A"], 5) == 1


def test_all_different_tasks():
    assert least_interval(["A", "B", "C", "D"], 3) == 4


def test_two_tasks_with_cooldown_1():
    assert least_interval(["A", "A", "B", "B"], 1) == 4
