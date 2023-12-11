# https://adventofcode.com/2023/day/11
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def expand_galaxies(gals, empty_rows, empty_cols, MULT):
  # expand rows
  ei = 0
  for gi, (gx, gy) in enumerate(gals):
    while ei < len(empty_rows) and empty_rows[ei] < gy:
      ei += 1
    gals[gi] = (gx, gy + (ei * MULT))
  # expand cols
  gals.sort()
  ei = 0
  for gi, (gx, gy) in enumerate(gals):
    while ei < len(empty_cols) and empty_cols[ei] < gx:
      ei += 1
    gals[gi] = (gx + (ei * MULT), gy)
  return gals

def distances(galaxies):
  man_dist = lambda g1, g2: abs(g1[0] - g2[0]) + abs(g1[1] - g2[1])
  dists = {}
  for i,g1 in enumerate(galaxies):
    for g2 in galaxies[i+1:]:
      d = man_dist(g1, g2)
      dists[(g1,g2)] = d
  return dists

part1 = 0
part2 = 0

P1_MULTIPLIER = 2-1 # part1
P2_MULTIPLIER = 1000000-1 # part2
# P2_MULTIPLIER = 10-1 # part2 example

universe = []
for line in fin:
  universe.append(line.rstrip())

empty_rows = [y for y, row in enumerate(universe) if all(r == '.' for r in row)]
empty_cols = [x for x in range(len(universe[0])) if all(row[x] == '.' for row in universe)]

galaxies = [(x,y) for y, row in enumerate(universe) for x, cell in enumerate(row) if cell == '#']

for part2 in [False, True]:
  MULT = P2_MULTIPLIER if part2 else P1_MULTIPLIER
  gals = expand_galaxies(galaxies.copy(), empty_rows, empty_cols, MULT)

  if part2:
    part2 = sum(distances(gals).values())
  else:
    part1 = sum(distances(gals).values())

print("Part 1:", part1)
print("Part 2:", part2)
