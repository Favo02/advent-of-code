# https://adventofcode.com/2023/day/8
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

# ======================================================
# generalized version for day8 part2:
#   the other solution (./day08.py) is works only for very specific inputs
#   in which a lot of assumptions should be made:
#     - each A encounters only one Z
#     - the cycle from Zi back to Zi doesnt encounter any other A or Z
#     - the time needed to reach Zi from Ai is the same as the time to cycle back to Zi
#     - the sequence of instructions (L and Rs) is synced with cycles
#   in this solution, a few assumptions can be removed
#
# run with pypy (~40s) for better performances than cpython (~500s)
# ======================================================

lines = []
for line in fin:
  lines.append(line.rstrip())

# cycles could start at differnt times:
#     the time needed to reach Zi from Ai could be different
#     than the time to reach again Zi from Zi
# cycles is a list of tuples: (first_reach, time_to_cycle) where
#     first_reach =     time needed to reach Zi the first time from Ai
#     time_to_cycle =   time needed to reach again Zi from Zi
# this funcion will find the minimum point in time where all the cycles are
# "synced" and all paths are on a Z, with an "optimized" bruteforce approach
def synced_lcm(cycles, time):
  times_to_cycle = [t[1] for t in cycles]
  # pick the biggest to start bruteforcing
  jump_size = max(times_to_cycle)

  is_lcm = False

  while not is_lcm:

    is_lcm = True
    for first_reach, cycle in cycles:
      # sync the time (removing time needed to first reach Zi (first_reach))
      # and then check lcm
      if (time - first_reach) % cycle != 0:
        is_lcm = False

    if not is_lcm:
      time += jump_size

  # lowest time that is a lcm of all synced cycles
  return time

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
incomplete_cycles = {}
cycles = {}

while len(cycles) < len(starts):
  dir = 0 if (dirs[time % len(dirs)] == "L") else 1

  for i, cur in enumerate(starts):
    if cur in cycles:
      continue

    starts[i] = graph[cur][dir]

    if starts[i][2] == "Z":
      if starts[i] in incomplete_cycles:
        cycles[starts[i]] = (time+1, time - incomplete_cycles[starts[i]])
      else:
        incomplete_cycles[starts[i]] = time

  time += 1

part2 = synced_lcm(list(cycles.values()), time)
print("Part 2 (generalized):", part2)
