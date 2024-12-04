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
    if cell == "A":
      queue.append((x, y))

ROWS = len(field)
COLS = len(field[0])

while queue:
  x, y = queue.popleft()

  found = []
  for dx, dy in [(-1,-1), (1,-1), (1,1), (-1,1)]:
    nx, ny = x + dx, y + dy
    if not (0 <= nx < COLS): continue
    if not (0 <= ny < ROWS): continue
    if field[ny][nx] not in ["M", "S"]: continue
    found.append(field[ny][nx])

  # print(x, y, found)

  if len(found) < 4:
    # print("short")
    continue

  # print(x, y)
  count = Counter(found)
  if "M" not in count or count["M"] != 2: continue
  if "S" not in count or count["S"] != 2: continue
  valid = False
  for a, b in zip(found, found[1:]):
    if a == b: valid = True

  if valid:
    print(x, y, "valid: ", valid)
    res += 1

print("RES:", res)
