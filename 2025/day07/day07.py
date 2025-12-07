# https://adventofcode.com/2025/day/7
# https://github.com/Favo02/advent-of-code

import sys
from functools import cache
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

grid = list(map(str.strip, fin.read().strip().splitlines()))

splits = set()

@cache
def ways(x, y):
    if not (0 <= x < COLS): return 0
    if not (0 <= y+1 < ROWS): return 0

    while y+1 < ROWS:
        if grid[y+1][x] == "^":
            splits.add((x, y+1))
            return ways(x-1, y+1) + ways(x+1, y+1)
        y += 1
    return 1

ROWS, COLS = len(grid), len(grid[0])

assert "S" in grid[0], "Missing S (assumed on first line)"
part2 = ways(grid[0].index("S"), 0)
part1 = len(splits)

print("Part 1:", part1)
print("Part 2:", part2)
