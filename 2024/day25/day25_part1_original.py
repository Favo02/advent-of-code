from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def tocode(a):
  cols = [0] * len(a[0])
  for row in a:
    for c, cell in enumerate(row):
      if cell == "#":
        cols[c] += 1
  return tuple(c-1 for c in cols)

res = 0

alll = fin.read().split("\n\n")
print(alll)

locks = []
keys = []

for l in alll:
  if l[0] == "#":
    locks.append((tocode(l.strip().split("\n"))))
  else:
    keys.append((tocode(l.strip().split("\n"))))

print(locks)
print(keys)

def check(l, k):
  assert len(l) == len(k)
  for cl, ck in zip(l, k):
    if cl+1 + ck+1 > 7:
      return False
  return True


H = 7

res = 0

for ii, l in enumerate(locks):
  for i, k in enumerate(keys):
    print(l, k, check(l,k))
    if check(l,k):
      res += 1


print(res)
