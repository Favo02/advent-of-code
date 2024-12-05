from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = 0

ordering = defaultdict(set)

for y, line in enumerate(fin):
  if line == "\n": break
  a, b = map(int, line.split("|"))
  ordering[b].add(a)

updates = []
for y, line in enumerate(fin):
  updates.append(list(map(int, line.split(","))))

for i, update in enumerate(updates):
  banned = set()
  valid = True

  for elem in update:
    if elem in banned:
      valid = False
      break

    print(ordering[elem])
    banned.update(ordering[elem])

  if valid:
    assert len(update)%2 == 1
    res += update[len(update) // 2]

print(updates)

print(res)
