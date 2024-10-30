import re
import sympy

with open("input.txt") as file:
    data = {key: value for (key, value)
            in [line.split(": ") for line in file.read().splitlines()]}


def contains(root, name):
    if root == name:
        return True
    numre = r"(\d+)"
    val = data[root]
    if re.match(numre, val):
        return False
    opre = r"(\w+) [+\-*\/] (\w+)"
    groups = re.search(opre, data[root]).groups()
    return contains(groups[0], name) or contains(groups[1], name)


def create_expr(name):
    numre = r"(\d+)"  # \w for humn also
    val = data[name]

    # If it was a number yelling monkey we know the value
    if re.match(numre, val):
        return val
    if val == "humn":
        return val

    # Otherwise it was an operation monkey and we need to recurse further
    opre = r"(\w+) ([+\-*\/]) (\w+)"
    groups = re.search(opre, val).groups()
    expr1 = create_expr(groups[0])
    expr2 = create_expr(groups[2])
    op = groups[1]
    return f"({expr1} {op} {expr2})"


def getval(name):
    numre = r"(\d+)"
    val = data[name]

    # If it was a number yelling monkey we know the value
    if re.match(numre, val):
        return int(val)

    # Otherwise it was an operation monkey and we need to recurse further
    opre = r"(\w+) ([+\-*\/]) (\w+)"
    groups = re.search(opre, val).groups()
    val1 = getval(groups[0])
    val2 = getval(groups[2])
    op = groups[1]
    return int(eval(f"{val1} {op} {val2}"))


oppositeOp = {
    "+": "-",
    "-": "+",
    "*": "/",
    "/": "*",
}

val = data["root"]
opre = r"(\w+) [+\-*\/] (\w+)"
groups = re.search(opre, val).groups()
res = {}

humnleft = contains(groups[0], "humn")

known = getval(groups[1]) if humnleft else getval(groups[0])
data["humn"] = "humn"
todo = groups[0] if humnleft else groups[1]

# Use symbolic expression solver
# TODO: Solve this by hand one day
humn = sympy.Symbol("humn")
print(sympy.solve(create_expr(todo) + "-" + str(known), humn)[0])
