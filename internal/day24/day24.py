from collections import deque

with open("input.txt") as file:
    board = [line[1:-1] for line in file.read().splitlines()[1:-1]]
    xmax = len(board[0])
    ymax = len(board)

def dfs(board, start_step, start, end):
    visited = {(start, start_step)}
    q = deque()
    q.append((start, start_step))
    while len(q) != 0:
        v = q.popleft()
        pos, step = v
        xpos, ypos = pos
        for x, y in (pos, (xpos+1, ypos), (xpos-1, ypos), (xpos, ypos+1), (xpos, ypos-1)):
            if (x, y) == end:
                return step
            if y == start[1] and x == start[0]:
                # In case we cannot move out of start turn one
                visited.add(((x, y), step+1))
                q.append(((x, y), step+1))
            if x < 0 or x >= xmax or y < 0 or y >= ymax:
                continue
            if board[y][(x-step) %xmax] == ">":
                continue
            if board[y][(x+step) %xmax] == "<":
                continue
            if board[(y-step)%ymax][x] == "v":
                continue
            if board[(y+step)%ymax][x] == "^":
                continue
            if ((x, y), step + 1) in visited:
                continue
            visited.add(((x, y), step+1))
            q.append(((x, y), step+1))

start = (0, -1)
end = (xmax-1, ymax)
there = dfs(board, 0, start, end)
print(there)
back = dfs(board, there, end, start)
there_again = dfs(board, back, start, end)
print(there_again)
    