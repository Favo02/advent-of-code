from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(sx, sy, end):
  seen = set()

  queue = deque()
  queue.append((0, sx, sy))

  res = 0

  while queue:
    dist, x, y = queue.popleft()

    if (x, y) == end:
      return dist

    for dx, dy in [(0,1), (0,-1), (1,0), (-1,0)]:
      nx, ny = x+dx, y+dy

      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if field[ny][nx] == "#": continue
      if (nx, ny) in seen: continue
      seen.add((nx, ny))

      queue.append((dist+1, nx, ny))

  return res

res = 0

field = []
start = end = None

for y, line in enumerate(fin):
  line = line.strip()
  field.append(list(line))
  if "S" in line: start = line.index("S"), y
  if "E" in line: end = line.index("E"), y

ROWS = len(field)
COLS = len(field[0])

nocheats = bfs(*start, end)

for r, row in enumerate(field):
  for c, cell in enumerate(row):
    if cell == "#":
      field[r][c] = "."
      dist = bfs(*start, end)
      if dist <= nocheats - 100:
        res += 1
      field[r][c] = "#"

print("RES:", res)
