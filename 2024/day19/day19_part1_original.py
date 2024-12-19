from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

@lru_cache(None)
def solve(tow, i):
  if i == len(tow): return True
  if i > len(tow): return False

  for avail in available:
    if len(tow) - i < len(avail): continue
    if tow.startswith(avail, i):
      if solve(tow, i+len(avail)):
        return True

  return False


available = input().split(", ")
input()

res = 0

towels = []
for y, line in enumerate(fin):
  line = line.rstrip()
  towels.append(line)
  # if y % 10 == 0: print(y)
  if solve(line, 0):
    res += 1

print(res)
