import sys
from collections import deque

grid = list(map(str.strip, sys.stdin.read().strip().splitlines()))

ROWS, COLS = len(grid), len(grid[0])

q = deque()

start = None
for y, line in enumerate(grid):
    if "S" in line:
        start = line.index("S"), y
        q.append(start)
        break

res = 0
seen = set()
while q:
    x, y = q.popleft()
    if (x, y) in seen: continue
    seen.add((x, y))
    if not (0 <= x < COLS): continue
    if not (0 <= y+1 < ROWS): continue
    if grid[y+1][x] == "^":
        q.append((x-1, y+1))
        q.append((x+1, y+1))
        res += 1
    else:
        q.append((x, y+1))


print(res)
