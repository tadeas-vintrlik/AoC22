with open("internal/day22/test.txt") as file:
    board, instructions = [x for x in file.read().split("\n\n")]

# Parse the board, adjust all lines to be the same length
board = [x for x in board.split("\n")]
xmax = max(list(map(lambda x: len(x), board))) + 2
board = [" " + x.ljust(xmax) + " " for x in board]
board = [" "*xmax] + board + [" "*xmax]
ymax = len(board)

size = 4

xcube = xmax//size
ycube = ymax//size
cube = [[None for _ in range(xcube)] for _ in range(ycube)]

num = 1
for j in range(ycube):
    for i in range(xcube):
        if board[1 + (j*size)][1 + (i * size)] != ' ':
            cube[j][i] = num
            num += 1

adjecent = {x: {"left": None, "right": None,
                "bottom": None, "top": None} for x in range(1, num)}

# Get neighbours from the cube matrix itself
for key in range(1, num):
    i, j = None, None
    for j, row in enumerate(cube):
        for i, cell in enumerate(row):
            if cell == key:
                x, y = i, j

    adjecent[key]["top"] = cube[(y-1) % ycube][x % xcube]
    adjecent[key]["bottom"] = cube[(y+1) % ycube][x % xcube]
    adjecent[key]["left"] = cube[y % ycube][(x-1) % xcube]
    adjecent[key]["right"] = cube[y % ycube][(x+1) % xcube]

# Need to add a None key to make the next part easier and not have to check for None
adjecent[None] = {"left": None, "right": None,
                  "bottom": None, "top": None}

# Find our neighbours we do not know from our neighbours we do know
# Needs to run twice to fill all the values
# TODO: This actually does not work entirely as our left's top might not be the same as ours
# The orientation is a bit more difficult
for _ in range(2):
    for key in range(1, num):
        adjecent[key]["top"] = adjecent[key]["top"] or adjecent[adjecent[key]["left"]
                                                                ]["top"] or adjecent[adjecent[key]["right"]]["top"]
        adjecent[key]["bottom"] = adjecent[key]["bottom"] or adjecent[adjecent[key]["left"]
                                                                      ]["bottom"] or adjecent[adjecent[key]["right"]]["bottom"]
        adjecent[key]["left"] = adjecent[key]["left"] or adjecent[adjecent[key]["top"]
                                                                  ]["left"] or adjecent[adjecent[key]["bottom"]]["left"] or adjecent[key]["left"]
        adjecent[key]["right"] = adjecent[key]["right"] or adjecent[adjecent[key]["top"]
                                                                    ]["right"] or adjecent[adjecent[key]["bottom"]]["right"]

del adjecent[None]

for row in cube:
    print(row)
print(adjecent)
