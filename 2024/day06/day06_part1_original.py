from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0
start = None
field = []
for y, line in enumerate(fin):
  line = line.rstrip()
  field.append(line)

  for x, cell in enumerate(line):
    if cell == "^":
      start = x, y

ROWS = len(field)
COLS = len(field[0])

direc = [(0,-1), (1,0), (0,1), (-1, 0)]
dir = 0

x, y = start
seen = set()



print(field)
print(start)

while True:
  seen.add((x, y))

  dx, dy = direc[dir]
  nx, ny = x + dx, y + dy

  if not ((0 <= nx < COLS) and (0 <= ny < ROWS)):
    break

  if field[ny][nx] == "#":
    dir = (dir+1)%4
    continue

  x, y = nx, ny



print("RES:", len(seen))
