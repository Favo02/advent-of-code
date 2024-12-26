from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def generate(s):

  r1 = s*64
  s ^= r1
  s %= 16777216

  r2 = s//32
  s ^= r2
  s %= 16777216

  r3 = s * 2048
  s ^= r3
  s %= 16777216

  return s

res = 0

field = []
for y, line in enumerate(fin):
  line = line.strip()
  n = int(line)

  for _ in range(2000):
    n = generate(n)
  print(n)
  res += n

print("RES:", res)
