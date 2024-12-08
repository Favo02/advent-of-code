# https://adventofcode.com/2024/day/8
# https://github.com/Favo02/advent-of-code

from collections import defaultdict
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

field = []
antennas = defaultdict(list)

for y, line in enumerate(fin):
  line = line.strip()
  field.append(line)

  for x, cell in enumerate(line):
    if cell != ".":
      antennas[cell].append((x, y))

ROWS = len(field)
COLS = len(field[0])

part1 = set()
part2 = set()

for freq, ant in antennas.items():

  for x1, y1 in ant:
    for x2, y2 in ant:
      if x1 == x2 and y1 == y2: continue

      dx, dy = x1-x2, y1-y2

      for mult in range(1, max(ROWS, COLS)):
        nx, ny = y2 + mult*dy, x2 + mult*dx
        if not (0 <= ny < ROWS): break
        if not (0 <= nx < COLS): break

        if abs(mult) == 2: part1.add((nx, ny))
        part2.add((nx, ny))

print("Part 1:", len(part1))
print("Part 2:", len(part2))
