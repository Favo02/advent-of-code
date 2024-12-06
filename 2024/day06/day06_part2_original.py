from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

DIRS = [(0,-1), (1,0), (0,1), (-1, 0)]

def walk(x, y, dir):
  seen = set()

  while True:
    if (x, y, dir) in seen:
      return True
    seen.add((x, y, dir))

    dx, dy = DIRS[dir]
    nx, ny = x + dx, y + dy

    if not ((0 <= nx < COLS) and (0 <= ny < ROWS)):
      break

    if field[ny][nx] == "#":
      dir = (dir+1)%4
      continue

    x, y = nx, ny
  return False

res = 0
start = None
obstacles = set()
field = []
for y, line in enumerate(fin):
  line = line.rstrip()
  field.append(list(line))

  for x, cell in enumerate(line):
    if cell == "#":
      obstacles.add((x, y))
    if cell == "^":
      start = x, y

ROWS = len(field)
COLS = len(field[0])

res = 0

for y, row in enumerate(field):
  for x, cell in enumerate(row):
    if field[y][x] not in "#^":
      field[y][x] = "#"
      if walk(*start, 0):
        res += 1
      field[y][x] = "."



print(res)
