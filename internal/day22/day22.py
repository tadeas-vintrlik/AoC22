import re

with open("internal/day22/input.txt") as file:
    board, instructions = [x for x in file.read().split("\n\n")]

# Parse the board, adjust all lines to be the same length
board = [x for x in board.split("\n")]
xmax = max(list(map(lambda x: len(x), board))) + 2
board = [" " + x.ljust(xmax) + " " for x in board]
board = [" "*xmax] + board + [" "*xmax]
ymax = len(board)

# Parse the instruction, split, convert the ingetegers to ingeters
instructions = list(map(lambda x: int(x) if not x.isalpha()
                        else x, re.split(r"(\d+)", instructions)[1:-1]))

# Using complex number a coordinate, imag is y, real is x
pos = board[1].index(".") + 1j
# The direction we are facing is in fact a vector we add for each movement
dir = 1 + 0j


def help_print():
    for i, row in enumerate(board):
        for j, cell in enumerate(row):
            if i == pos.imag and j == pos.real:
                print("P", end="")
            else:
                print(cell, end="")
        print("")


for ins in instructions:
    match ins:
        case "R":
            dir *= 1j
        case "L":
            dir /= 1j
        case x:
            for _ in range(x):
                npos = pos + dir
                # Check for overflow
                while board[int(npos.imag)][int(npos.real)] == " ":
                    npos += dir
                    npos = npos.real % (xmax-1) + \
                        (npos.imag % (ymax-1)) * 1j
                if board[int(npos.imag)][int(npos.real)] == "#":
                    break  # We hit a wall the movement makes no sense anymore
                pos = npos

# Finished following instructions print the result now
dirscore = 0
match dir:
    case (1j):
        dirscore = 1
    case (-1+0j):
        dirscore = 2
    case (-1j):
        dirscore = 3
print(int(pos.imag * 1000 + pos.real * 4 + dirscore))
