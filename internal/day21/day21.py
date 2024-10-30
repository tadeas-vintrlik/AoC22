import re

with open("input.txt") as file:
    data = {key: value for (key, value)
            in [line.split(": ") for line in file.read().splitlines()]}


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


print(getval("root"))
