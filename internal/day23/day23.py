with open("input.txt") as file:
    # Collect the elves a complex number of the coordinates
    elves = {x + y*1j for (y, v) in [(y, v) for (y, v) in enumerate(
        file.read().splitlines())] for (x, c) in enumerate(v) if c == "#"}

# North, south, west and east directions
directions = [-1j, 1j, -1, 1]


def create_board(elves):
    ymin = int(min(list(map(lambda x: x.imag, elves))))
    ymax = int(max(list(map(lambda x: x.imag, elves))) + 1)
    xmin = int(min(list(map(lambda x: x.real, elves))))
    xmax = int(max(list(map(lambda x: x.real, elves))) + 1)
    return [["#" if x + y*1j in elves else "." for x in range(xmin, xmax)] for y in range(ymin, ymax)]


def draw_board(elves):
    for row in create_board(elves):
        for cell in row:
            print(cell, end="")
        print()


def get_neighbours(elf, elves):
    return [neighbour for neighbour in [elf.real+x + (elf.imag+y)*1j for x in (0, -1, 1) for y in (0, -1, 1) if x != 0 or y != 0] if neighbour in elves]


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

    #  Part two of round
    elves = {elf if elf not in proposed else proposed[elf] for elf in elves}
    return len(proposed) != 0


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
