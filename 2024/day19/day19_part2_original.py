from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

@lru_cache(None)
def solve(tow, i):
  if i == len(tow): return 1
  if i > len(tow): return 0

  res = 0

  for avail in available:
    if len(tow) - i < len(avail): continue
    if tow.startswith(avail, i):
      sub = solve(tow, i+len(avail))
      res += sub

  return res


available = input().split(", ")
input()

res = 0

towels = []
for line in fin:
  line = line.rstrip()
  towels.append(line)
  rr = solve(line, 0)
  print(line, rr)
  res += rr

print(res)
