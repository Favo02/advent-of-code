from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

# https://www.geeksforgeeks.org/maximal-clique-problem-recursive-solution/
def bron_kerbosch(R, P, X, graph):
    if not P and not X:
        yield R
    while P:
        v = P.pop()
        yield from bron_kerbosch(
            R.union({v}),
            P.intersection(graph[v]),
            X.intersection(graph[v]),
            graph
        )
        X.add(v)

entries = []
keys = set()
for y, line in enumerate(fin):
  a, b = line.strip().split("-")
  keys.add(a)
  keys.add(b)
  entries.append((a,b))

keys = sorted(list(keys))
ktoi = {k: i for i, k in enumerate(keys)}
print(keys)
print(ktoi)

graph = defaultdict(set)
for a, b in entries:
    graph[a].add(b)
    graph[b].add(a)

cliques = list(bron_kerbosch(set(), set(graph.keys()), set(), graph))
best = None
for c in cliques:
    if best is None or len(c) > len(best):
        best = c

print(",".join(sorted(best)))
