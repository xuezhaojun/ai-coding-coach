def solve_eval_rpn(tokens: list[str]) -> int:
    """Evaluate a reverse polish notation expression using a stack.

    Time: O(n), Space: O(n)
    """
    stack: list[int] = []
    for token in tokens:
        if token in ("+", "-", "*", "/"):
            b = stack.pop()
            a = stack.pop()
            if token == "+":
                stack.append(a + b)
            elif token == "-":
                stack.append(a - b)
            elif token == "*":
                stack.append(a * b)
            elif token == "/":
                # Integer division truncating toward zero (matches Go behavior)
                stack.append(int(a / b))
        else:
            stack.append(int(token))
    return stack[0]
