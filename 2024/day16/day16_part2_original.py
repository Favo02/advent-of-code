from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def pprint():
  for f in field:
    print("".join(f))

DIRS = ((1, 0), (0, -1), (-1, 0), (0, 1))

def distcost(dir, newdir):
  return min(abs(dir - newdir), (dir + (4 - newdir)), (4-dir) + newdir) * 1000

def dijkstra(sx, sy):

  queue = []
  heappush(queue, (0, sx, sy, 0))

  dist = defaultdict(lambda: float("inf"))
  dist[(sx, sy, 0)] = 0

  while queue:
    d, x, y, dir = heappop(queue)
    if d != dist[(x, y, dir)]:
      assert d > dist[(x, y, dir)]
      continue

    # if field[y][x] == "E": return d

    for ndir, (dx, dy) in enumerate(DIRS):
      nx, ny = x+dx, y+dy
      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if field[ny][nx] == "#": continue
      ndist = dist[(x, y, dir)] + distcost(dir, ndir) + 1
      if ndist < dist[(nx, ny, ndir)]:
        dist[(nx, ny, ndir)] = ndist
        heappush(queue, (ndist, nx, ny, ndir))

  return dist

best = set()
def back(dist, x, y, dir):
  # print(f"{x=} {y=} {dir=} {dist[(x, y, dir)]=}")
  # pprint()
  for ndir, (dx, dy) in enumerate(DIRS):
    nx, ny = x+dx, y+dy
    if not (0 <= nx < COLS): continue
    if not (0 <= ny < ROWS): continue
    if field[ny][nx] == "#": continue

    for nnndir in range(4):
      dcost = distcost(dir, nnndir)

      # print(f"  {nx=} {ny=} {nnndir=} {dist[(nx, ny, nnndir)]=}")

      if dist[(nx, ny, nnndir)] == dist[(x, y, dir)] - 1 - dcost:
        field[ny][nx] = "O"
        best.add((nx, ny))
        back(dist, nx, ny, nnndir)

res = 0

start = None

field = []
for y, line in enumerate(fin):
  line = line.strip()
  field.append(list(line))

  if "S" in line:
    start = line.index("S"), y
  if "E" in line:
    ex, ey = line.index("E"), y

ROWS = len(field)
COLS = len(field[0])

dist = dijkstra(*start)

besttt = res = float("inf")

for d in range(4):
  if dist[(ex, ey, d)] < res:
    res = dist[(ex, ey, d)]
    besttt = ex, ey, d


back(dist, *besttt)

print(len(best) + 1)


# print("RES:", dijkstra(*start))
