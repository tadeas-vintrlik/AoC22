import math

with open("input.txt") as file:
    # Collect the elves a complex number of the coordinates
    elves = [x + y*1j for (y, v) in [(y, v) for (y, v) in enumerate(
        file.read().splitlines())] for (x, c) in enumerate(v) if c == "#"]

# North, south, west and east directions
directions = [-1j, 1j, -1, 1]


def create_board(elves):
    ymin = int(min(list(map(lambda x: x.imag, elves))))
    ymax = int(max(list(map(lambda x: x.imag, elves))) + 1)
    xmin = int(min(list(map(lambda x: x.real, elves))))
    xmax = int(max(list(map(lambda x: x.real, elves))) + 1)
    return [["#" if x + y*1j in [e for e in elves] else "." for x in range(xmin, xmax)] for y in range(ymin, ymax)]


def draw_board(elves):
    for row in create_board(elves):
        for cell in row:
            print(cell, end="")
        print()


def distance(x, y):
    return math.sqrt((x.real - y.real)**2 + (y.imag - x.imag)**2)


def get_neighbours(elf, elves):
    return [neighbour for neighbour in list(filter(lambda x: distance(x, elf) < 2, elves)) if neighbour != elf]


def round(number):
    global elves
    moved = False
    # Part one of round
    proposed = {}
    for i, elf in enumerate(elves):
        n = get_neighbours(elf, elves)
        if len(n) == 0:
            continue  # He has no neighbours therefore no reason to move
        # Consider movement in direction
        for di in range(number, number+4):
            dir = directions[di % len(directions)]
            if dir.real == 0:
                # if movement in the imaginary axis check for colissions there
                cols = [neib for neib in n if (
                    neib-dir).imag == elf.imag]
            else:
                # If movement in real axis check for colissions there instead
                cols = [neib for neib in n if (
                    neib-dir).real == elf.real]
            # If there were no colissions put it as proposed movement unless contested
            new = elf + dir
            if len(cols) != 0:
                continue
            if new in proposed.values():
                del proposed[list(proposed.keys())[
                    list(proposed.values()).index(new)]]
                break
            else:
                proposed[elf] = new
                break
        # elves[i] = (elf[0], (elf[1] + 1))

    #  Part two of round
    for i, elf in enumerate(elves):
        if elf in proposed:
            elves[i] = proposed[elf]
            moved = True
    return moved


moved = True
i = 0
while moved:
    moved = round(i)
    i += 1
    if i == 10:
        # Part 1
        board = create_board(elves)
        print("".join([cell for row in board for cell in row]).count("."))

# Part 2
print(i)
