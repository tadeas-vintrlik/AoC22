import re

with open("input.txt") as file:
    r = r"Blueprint \d+: Each ore robot costs (\d+) ore\. Each clay robot costs (\d+) ore\. Each obsidian robot costs (\d+) ore and (\d+) clay\. Each geode robot costs (\d+) ore and (\d+) obsidian\."
    bps = [tuple(map(lambda x: int(x), re.search(r, x).groups()))
           for x in file.read().splitlines()]


def dfs(bp):
    global maxg
    co, cc, cobo, cobc, cgo, cgob = bp  # Costs of different robots
    # Stack of states ( we start with 24 minutes and 1 ore robot)
    s = [(24, 0, 0, 0, 0, 1, 0, 0, 0)]
    visited = set()

    mo = max(co, cc, cobo, cgo)  # Maximum ore cost

    while len(s) != 0:
        state = s.pop()
        if state in visited:
            continue
        visited.add(state)
        t, o, c, ob, g, ro, rc, rob, rg = state

        # TODO: terribly slow needs more pruning
        if (
                # Out of time
                t == 0 or
                # Built more robots than neccessary
                ro > mo or rc > cobc or rob > cgob or
                # Stockpiled more resources than neccessary
                (o > mo and c > cobc and ob > cgob)
        ):
            maxg = max(maxg, g)
            continue
        # print(f"{state}")
        t -= 1

        # See if we can build
        if o >= co:
            s.append((t, o+ro-co, c+rc, ob+rob, g+rg, ro+1, rc, rob, rg))
        if o >= cc:
            s.append((t, o+ro-cc, c+rc, ob+rob, g+rg, ro, rc+1, rob, rg))
        if o >= cobo and c >= cobc:
            s.append((t, o+ro-cobo, c+rc-cobc, ob +
                      rob, g+rg, ro, rc, rob+1, rg))
        if o >= cgo and ob >= cgob:
            s.append((t, o+ro-cgo, c+rc, ob +
                      rob-cgob, g+rg, ro, rc, rob, rg+1))
        # We are just collecting resources this minute
        s.append((t, o+ro, c+rc, ob+rob, g+rg, ro, rc, rob, rg))


res = 0
for i, bp in enumerate(bps):
    maxg = 0
    dfs(bp)
    res += (i+1)*maxg
    print(maxg)
print(res)
