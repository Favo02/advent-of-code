# https://adventofcode.com/2023/day/25
# https://github.com/Favo02/advent-of-code

import sys
from networkx import Graph, minimum_cut
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def solve(graph):
  for start in graph.nodes:
    for end in graph.nodes:
      if start == end:
        continue
      edges_cutted, partitions = minimum_cut(graph, start, end)
      if edges_cutted == 3:
        return len(partitions[0]) * len(partitions[1])

part1 = 0
part2 = 0

graph = Graph()

for line in fin:
  node, adjs = line.strip().split(": ")
  for adj in adjs.split():
    graph.add_edge(node, adj, capacity=1)
    graph.add_edge(adj, node, capacity=1)

part1 = solve(graph)

print("Part 1:", part1)
