from reverse_bits import reverse_bits


def test_example_1():
    assert reverse_bits(0b00000010100101000001111010011100) == 964176192


def test_example_2():
    assert reverse_bits(0b11111111111111111111111111111101) == 3221225471


def test_zero():
    assert reverse_bits(0) == 0


def test_all_ones():
    assert reverse_bits(0xFFFFFFFF) == 0xFFFFFFFF


def test_one():
    assert reverse_bits(1) == 0x80000000


def test_power_of_two():
    assert reverse_bits(0x80000000) == 1
