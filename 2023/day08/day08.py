# https://adventofcode.com/2023/day/8
# https://github.com/Favo02/advent-of-code

import sys
from math import lcm
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = 0
part2 = 0

lines = []
for line in fin:
  lines.append(line.rstrip())

dirs = lines[0]
starts = []
ends = []
graph = {}

for line in lines[2:]:
  line = line.replace(")", "").replace("(", "")
  fromm, to = line.split(" = ")
  left, right = to.split(", ")

  if fromm[2] == "A":
    starts.append(fromm)
  if fromm[2] == "Z":
    ends.append(fromm)
  graph[fromm] = (left, right)

time = 0
dist = {}

while len(dist) < len(starts):
  dir = 0 if (dirs[time % len(dirs)] == "L") else 1

  for i, cur in enumerate(starts):
    if cur not in dist and cur[2] == "Z":
      dist[starts[i]] = time
    else:
      starts[i] = graph[cur][dir]

  time += 1

part1 = dist["ZZZ"]
part2 = lcm(*dist.values()) # unpack list

print("Part 1:", part1)
print("Part 2:", part2)
