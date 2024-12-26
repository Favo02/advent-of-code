# https://adventofcode.com/2024/day/23
# https://github.com/Favo02/advent-of-code

from collections import defaultdict
from itertools import combinations
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

# https://www.geeksforgeeks.org/maximal-clique-problem-recursive-solution/
def bron_kerbosch(R, P, X, graph):
  if not P and not X:
    yield frozenset(R)
  while P:
    v = P.pop()
    yield from bron_kerbosch(
      R.union({v}),
      P.intersection(graph[v]),
      X.intersection(graph[v]),
      graph
    )
    X.add(v)

def size3(cliques):
  for cl in cliques:
    if len(cl) == 3:
      yield cl
    if len(cl) > 3:
      yield from (frozenset(c) for c in combinations(cl, 3))

entries = []
keys = set()
for y, line in enumerate(fin):
  a, b = line.strip().split("-")
  keys.add(a)
  keys.add(b)
  entries.append((a,b))

ktoi = {k: i for i, k in enumerate(sorted(list(keys)))}

graph = defaultdict(set)
for a, b in entries:
  graph[a].add(b)
  graph[b].add(a)

cliques = list(bron_kerbosch(set(), set(graph.keys()), set(), graph))

part1 = sum(1 for cl in set(size3(cliques)) if any(pc[0] == "t" for pc in cl))
part2 = ",".join(sorted(max((len(cl), cl) for cl in cliques)[1]))

print("Part 1:", part1)
print("Part 2:", part2)
