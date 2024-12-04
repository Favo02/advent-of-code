from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0

field = []
for y, line in enumerate(fin):
  line = line.rstrip()
  field.append(line)

  for x, cell in enumerate(line):
    pass

ROWS = len(field)
COLS = len(field[0])

print("RES:", res)
