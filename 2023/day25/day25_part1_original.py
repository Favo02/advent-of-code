import advent
import networkx as nx

advent.setup(2023, 25)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

part1 = 0
part2 = 0

graph = nx.Graph()
nodes = set()
for line in lines:
  s, o = line.split(": ")
  o = o.split()
  graph.add_node(s)
  nodes.add(s)
  for oo in o:
    nodes.add(oo)
    graph.add_node(oo)
    graph.add_edge(s, oo, capacity=1)
    graph.add_edge(oo, s, capacity=1)

print(graph)
for s in nodes:
  for e in nodes:
    if s == e: continue
    cut_value, partitions = nx.minimum_cut(graph, s, e)
    print(cut_value, len(partitions))
    if cut_value == 3:
      part1 = len(partitions[0]) * len(partitions[1])
      break
  else:
    continue
  break

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
