from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def pprint():
  for f in field:
    print("".join(f))
  print()

def bfs(sx, sy):
  seen = set()

  queue = deque()
  queue.append((0, sx, sy))

  while queue:
    dist, x, y = queue.popleft()

    if (x, y) == (COLS, ROWS):
      return dist

    for dx, dy in [(0,1),(0,-1),(1,0),(-1,0)]:
      nx, ny = x+dx, y+dy

      if not (0 <= nx <= COLS): continue
      if not (0 <= ny <= ROWS): continue
      if (nx, ny) in seen: continue
      if (nx, ny) in bytes: continue

      seen.add((nx, ny))
      queue.append((dist+1, nx, ny))

  assert False

res = 0

bytes = set()
for y, line in enumerate(fin):
  line = line.rstrip()
  bytes.add(tuple(map(int, line.split(","))))
  # if y == 11: break
  if y == 1023: break

# assert len(bytes) == 12
assert len(bytes) == 1024

# ROWS, COLS = 6, 6
ROWS, COLS = 70, 70

field = [["#" if (x, y) in bytes else "." for x in range(COLS+1)] for y in range(ROWS+1)]

pprint()

res = bfs(0, 0)

print("RES:", res)
