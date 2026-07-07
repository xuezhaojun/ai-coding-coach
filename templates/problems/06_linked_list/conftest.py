"""Pytest conftest for the linked_list category.

Each problem's ``python/`` directory contains same-named shared modules
(``helpers``, ``solution``). With pytest's default (prepend) import mode, the
first-imported instance stays cached in ``sys.modules`` and is reused by
subsequent problem directories, which breaks collection of later problems.

Before each test module is imported, drop any cached shared module so it is
re-imported from the current test file's directory (which pytest inserts at
``sys.path[0]``).
"""
import sys

import pytest

# Module names shared across every problem's python/ directory.
_SHARED_MODULES = ("helpers", "solution")


@pytest.hookimpl(hookwrapper=True)
def pytest_collectstart(collector):
    if collector.__class__.__name__ == "Module":
        for name in _SHARED_MODULES:
            sys.modules.pop(name, None)
    yield
