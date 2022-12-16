import re
import functools


def parse(filename):
    with open(filename) as file:
        regex = r"Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)"
        graph = {key: value for (key, value) in [(lambda x: (x[0], {"flow": int(x[1]), "edges": x[2].split(", ")}))
                                                 (re.search(regex, line).groups()) for line in file.readlines()]}
    return graph


def FloydWarshallUnweighted(graph):
    keys = [k for k in graph.keys()]
    dist = {}
    for i in keys:
        for j in keys:
            dist[i, i] = 0
            dist[i, j] = float("inf")
            dist[j, i] = float("inf")
            if j in graph[i]["edges"]:
                dist[i, j] = 1
                dist[j, i] = 1

    for i in keys:
        for j in keys:
            for k in keys:
                dist[j, k] = min(dist[j, k], dist[j, i] + dist[i, k])

    return dist


@functools.cache
def part1(current, minutes, opened):
    """
Uses the global variable graph and dist to save stack memory in recursion.
    """
    if minutes <= 0:
        return 0
    best = 0
    pressure = 0

    if current != "AA":
        opened = set(opened)
        opened.add(current)
        minutes -= 1
        pressure = minutes * graph[current]["flow"]

    for k in [k for k in graph.keys() if k not in opened]:
        distance = dist[current, k]
        best = max(
            best,
            pressure + part1(k, minutes - distance, frozenset(opened))
        )

    return max(best, pressure)


@functools.cache
def part2(current, minutes, opened):
    """
Uses the global variable graph and dist to save stack memory in recursion.
    """
    if minutes <= 0:
        return part1("AA", 26, opened)
    best = 0
    pressure = 0

    if current != "AA":
        opened = set(opened)
        opened.add(current)
        minutes -= 1
        pressure = minutes * graph[current]["flow"]

    for k in [k for k in graph.keys() if k not in opened]:
        distance = dist[current, k]
        best = max(
            best,
            pressure + part2(k, minutes - distance, frozenset(opened))
        )

    return best


graph = parse("input.txt")
# Compute the distances between each pair of vertices
dist = FloydWarshallUnweighted(graph)
positiveKeys = [k for k in graph.keys() if graph[k]["flow"] > 0 or k == "AA"]
graph = {key: value for (key, value) in [(lambda k: (k, {"flow": graph[k]["flow"], "edges": [
    edge for edge in graph[k]["edges"] if edge in positiveKeys]}))(k) for k in positiveKeys]}
print(part1("AA", 30, frozenset(["AA"])))
print(part2("AA", 26, frozenset(["AA"])))
