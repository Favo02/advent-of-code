from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(sx, sy):
  dist = defaultdict(lambda: float("inf"))
  dist[(sx, sy)] = 0

  queue = deque()
  queue.append((sx, sy))

  while queue:
    x, y = queue.popleft()

    for dx, dy in [(0,1), (0,-1), (1,0), (-1,0)]:
      nx, ny = x+dx, y+dy

      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if field[ny][nx] == "#": continue
      if (nx, ny) in dist: continue
      dist[(nx, ny)] = dist[(x, y)] + 1
      queue.append((nx, ny))

  return dist

def mandist(x1, y1, x2, y2):
  return abs(x1-x2) + abs(y1-y2)

res = 0

field = []
start = end = None

for y, line in enumerate(fin):
  line = line.strip()
  #line = line.strip()[1:-1]
  field.append(list(line))
  if "S" in line: start = line.index("S"), y#-1
  if "E" in line: end = line.index("E"), y#-1

# field = field[1:-1]

ROWS = len(field)
COLS = len(field[0])

FROMSTART = bfs(*start)
FROMEND = bfs(*end)

# print(start, end)

nocheat = FROMSTART[end]

print("NO CHEAT:", nocheat)

results = defaultdict(int)
unique = set()

for y1, row1 in enumerate(field):
  for x1, cell1 in enumerate(row1):
    if FROMSTART[(x1, y1)] == float("inf"): continue


    for y2, row2 in enumerate(field):
      for x2, cell2 in enumerate(row2):
        if FROMEND[(x1, y1)] == float("inf"): continue

        cheat = mandist(x1, y1, x2, y2)

        if cheat <= 20:
          if (d := FROMSTART[(x1, y1)] + FROMEND[(x2, y2)] + mandist(x1, y1, x2, y2)) <= nocheat-100:
            results[d] += 1




res = 0
for saved in sorted(results.keys()):
  if saved >= 100:
    res += results[saved]
    # print(saved, results[saved])


print(res)
