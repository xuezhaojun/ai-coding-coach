import sys
def pytest_runtest_setup(item):
    for mod_name in list(sys.modules.keys()):
        if mod_name in ('helpers', 'solution'):
            del sys.modules[mod_name]
