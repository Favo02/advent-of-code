from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

res = set()

graph = defaultdict(set)
for y, line in enumerate(fin):
  a, b = line.strip().split("-")
  graph[a].add(b)
  graph[b].add(a)

for a in graph:
  for b in graph:
    for c in graph:

      if b in graph[a] and a in graph[b] and \
         b in graph[c] and c in graph[b] and \
         a in graph[c] and c in graph[a]:
        if a.startswith("t") or b.startswith("t") or c.startswith("t"):
          res.add(frozenset([a,b,c]))

print("RES:", len(res))
