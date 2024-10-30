import re

with open("input.txt") as file:
    r = r"Blueprint \d+: Each ore robot costs (\d+) ore\. Each clay robot costs (\d+) ore\. Each obsidian robot costs (\d+) ore and (\d+) clay\. Each geode robot costs (\d+) ore and (\d+) obsidian\."
    bps = [tuple(map(lambda x: int(x), re.search(r, x).groups()))
           for x in file.read().splitlines()]


def dfs(bp):
    global maxg
    co, cc, cobo, cobc, cgo, cgob = bp  # Costs of different robots
    # Stack of states ( we start with 24 minutes and 1 ore robot)
    s = [(24, 0, 0, 0, 0, 1, 0, 0, 0, -1)]
    visited = set()

    mo = max(co, cc, cobo, cgo)  # Maximum ore cost

    while len(s) != 0:
        state = s.pop()
        if state in visited:
            continue
        visited.add(state)
        t, o, c, ob, g, ro, rc, rob, rg, b = state

        # Finish building the robot we have started constructing
        match b:
            case 0:
                ro += 1
            case 1:
                rc += 1
            case 2:
                rob += 1
            case 3:
                rg += 1

        # TODO: terribly slow needs more pruning
        if t == 0 or ro > mo or rc > cobc or rob > cgob:
            maxg = max(maxg, g)
            continue
        # print(f"{state}")
        t -= 1

        # See if we can build
        if o >= co:
            s.append((t, o+ro-co, c+rc, ob+rob, g+rg, ro, rc, rob, rg, 0))
        if o >= cc:
            s.append((t, o+ro-cc, c+rc, ob+rob, g+rg, ro, rc, rob, rg, 1))
        if o >= cobo and c >= cobc:
            s.append((t, o+ro-cobo, c+rc-cobc, ob +
                      rob, g+rg, ro, rc, rob, rg, 2))
        if o >= cgo and ob >= cgob:
            s.append((t, o+ro-cgo, c+rc, ob +
                      rob-cgob, g+rg, ro, rc, rob, rg, 3))
        # We are just collecting resources this minute
        s.append((t, o+ro, c+rc, ob+rob, g+rg, ro, rc, rob, rg, -1))


res = 0
for i, bp in enumerate(bps):
    maxg = 0
    dfs(bp)
    res += (i+1)*maxg
    print(maxg)
print(res)
