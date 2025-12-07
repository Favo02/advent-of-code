import sys
from collections import deque
from functools import cache

grid = list(map(str.strip, sys.stdin.read().strip().splitlines()))

seen = set()

@cache
def ways(x, y):
    if not (0 <= x < COLS): return 0
    if not (0 <= y+1 < ROWS): return 0
    # if (x, y) in seen: return 0
    # seen.add((x, y))

    while y+1 < ROWS:
        if grid[y+1][x] == "^":
            return ways(x-1, y+1) + ways(x+1, y+1)
        y += 1
    return 1


ROWS, COLS = len(grid), len(grid[0])

q = deque()

start = None
for y, line in enumerate(grid):
    if "S" in line:
        start = line.index("S"), y
        q.append(start)
        break

print(ways(*start))
