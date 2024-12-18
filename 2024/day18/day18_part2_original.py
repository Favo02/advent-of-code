from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(sx, sy):
  seen = set()

  queue = deque()
  queue.append((0, sx, sy))

  while queue:
    dist, x, y = queue.popleft()

    if (x, y) == (COLS, ROWS):
      return True

    for dx, dy in [(0,1),(0,-1),(1,0),(-1,0)]:
      nx, ny = x+dx, y+dy

      if not (0 <= nx <= COLS): continue
      if not (0 <= ny <= ROWS): continue
      if (nx, ny) in seen: continue
      if (nx, ny) in bytes: continue

      seen.add((nx, ny))
      queue.append((dist+1, nx, ny))

  return False

# ROWS, COLS = 6, 6
ROWS, COLS = 70, 70

bytes = set()
for y, line in enumerate(fin):
  line = line.rstrip()
  bytes.add(tuple(map(int, line.split(","))))
  if not bfs(0, 0):
    print(line)
    break
