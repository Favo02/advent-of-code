from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def pprint():
  for f in field:
    print("".join(f))

def dist(x1, y1, x2, y2):
  return x1-x2, y1-y2

res = set()

field = []
antennas = defaultdict(list)

for y, line in enumerate(fin):
  line = line.rstrip()
  field.append(list(line))

  for x, cell in enumerate(line):
    if cell != ".":
      antennas[cell].append((x, y))

ROWS = len(field)
COLS = len(field[0])

for freq, ant in antennas.items():
  print(freq, antennas[freq])
  for i, (x1, y1) in enumerate(ant):
    for ii, (x2, y2) in enumerate(ant):
      if x1 == x2 and y1 == y2: continue

      dx, dy = dist(x1, y1, x2, y2)
      nx, ny = y2 + 2*dy, x2 + 2*dx
      if not (0 <= ny < ROWS): continue
      if not (0 <= nx < COLS): continue

      res.add((nx, ny))

print("RES:", len(res))
