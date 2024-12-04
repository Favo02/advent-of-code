from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0
queue = deque()

field = []
for y, line in enumerate(fin):
  print(line.rstrip())
  field.append(line)

  for x, cell in enumerate(line):
    if cell == "X":
      queue.append((x, y, "X", None, None))

ROWS = len(field)
COLS = len(field[0])

advance = {"X": "M", "M": "A", "A": "S"}

while queue:
  x, y, progress, dirx, diry = queue.popleft()
  if progress == "S":
    res += 1
    continue

  if progress == "X":
    for dx, dy in [(0,1), (0,-1), (1,0), (-1,0), (-1,-1), (-1,1), (1,-1), (1,1)]:
      nx, ny = x + dx, y + dy
      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if field[ny][nx] != advance[progress]: continue
      queue.append((nx, ny, advance[progress], dx, dy))

  else:
    nx, ny = x + dirx, y + diry
    if not (0 <= nx < COLS): continue
    if not (0 <= ny < ROWS): continue
    if field[ny][nx] != advance[progress]: continue
    queue.append((nx, ny, advance[progress], dirx, diry))



print("RES:", res)
