from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(x, y):
  poss = defaultdict(int)
  poss[(x, y)] = 1
  ends = set()
  seen = set()
  queue = deque()
  queue.append((x, y))
  while queue:
    x, y = queue.popleft()
    for dx, dy in [(0,1),(0,-1),(1,0),(-1,0)]:
      nx, ny = x+dx, y+dy
      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if field[ny][nx] == ".": continue
      if int(field[ny][nx]) != int(field[y][x]) + 1: continue
      poss[(nx, ny)] += poss[(x, y)]
      if (nx, ny) in seen: continue
      seen.add((nx, ny))
      if (field[ny][nx]) == "9":
        ends.add((nx, ny))
      else:
        queue.append((nx, ny))

  res = 0
  for rx, ry in ends:
    print(rx, ry, poss[(rx, ry)])
    res += poss[(rx, ry)]
  return res

res = 0

field = []
starts = set()
for y, line in enumerate(fin):
  line = line.rstrip()
  field.append(line)

  for x, cell in enumerate(line):
    if cell == "0":
      starts.add((x, y))

ROWS = len(field)
COLS = len(field[0])

for sx, sy in starts:
  print(sx, sy, bfs(sx, sy))
  res += bfs(sx, sy)

print("RES:", res)
