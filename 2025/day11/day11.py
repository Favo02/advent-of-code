# https://adventofcode.com/2025/day/11
# https://github.com/Favo02/advent-of-code

import sys
from functools import cache
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

@cache
def dfs(cur, fft, dac):
    res = 0
    for adj in graph.get(cur, []):
        if adj == "out" and fft and dac:
            res += 1
        else:
            res += dfs(adj, fft or adj == "fft", dac or adj == "dac")
    return res

graph = {}

for f, t in map(lambda l: l.strip().split(": "), fin.readlines()):
    graph[f] = t.split()

part1 = dfs("you", True, True)
part2 = dfs("svr", False, False)

print("Part 1:", part1)
print("Part 2:", part2)
