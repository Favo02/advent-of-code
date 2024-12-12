from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(sx, sy):
  queue = deque()
  queue.append((sx, sy))
  type = field[sy][sx]
  seen = set()
  seen.add((sx, sy))
  while queue:
    x, y = queue.popleft()
    for dx, dy in [(0,1),(0,-1),(1,0),(-1,0)]:
      nx, ny = x+dx, y+dy
      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if field[ny][nx] != type: continue
      if (nx, ny) in seen: continue
      seen.add((nx, ny))
      queue.append((nx, ny))
  return seen

def perimeter(area):
  per = set()
  for x, y in area:
    # for dx, dy in [(0,1),(0,-1),(1,0),(-1,0),(1,1),(1,-1),(-1,1),(-1,-1)]:
    for dx, dy in [(0,1),(0,-1),(1,0),(-1,0)]:
      nx, ny = x+dx, y+dy
      if (nx, ny) in area: continue
      per.add((nx, ny, (dx, dy)))
  return per

def sides(per):
  sid = 0

  while per:
    x, y, dir = per.pop()
    sid += 1

    if (x+1, y, dir) in per or (x-1, y, dir) in per: deltas = [(1,0), (-1,0)]
    else: deltas = [(0,1), (0,-1)]

    print("L", x, y, dir, deltas)
    queue = deque()
    queue.append((x, y, dir))
    while queue:
      x, y, dir = queue.popleft()
      print(x, y)
      for dx, dy in deltas:
        if (x+dx, y+dy, dir) in per:
          per.remove((x+dx, y+dy, dir))
          queue.append((x+dx, y+dy, dir))
  return sid

res = 0

field = []
for y, line in enumerate(fin):
  line = line.strip()
  field.append(line)

for f in field:
  print(f)

ROWS = len(field)
COLS = len(field[0])

used = set()

for y, row in enumerate(field):
  for x, cell in enumerate(row):
    if (x,y) in used: continue
    print("---------", cell)
    area = bfs(x, y)
    per = perimeter(area)
    sid = sides(per)
    print(field[y][x], sid)
    used.update(area)
    res += len(area)*sid

assert len(used) == ROWS*COLS

print("RES:", res)
