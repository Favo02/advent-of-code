# https://adventofcode.com/2024/day/6
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

start = None
field = []

for y, line in enumerate(fin):
  line = line.rstrip()
  field.append(list(line))
  if "^" in line:
    start = line.find("^"), y

ROWS = len(field)
COLS = len(field[0])

DIRS = [(0,-1), (1,0), (0,1), (-1, 0)]

def walk(x, y):
  dir = 0
  cycles = set()

  while True:
    if (x, y, dir) in cycles:
      return True, float("inf")

    cycles.add((x, y, dir))

    dx, dy = DIRS[dir]
    nx, ny = x + dx, y + dy

    if not (0 <= nx < COLS): break
    if not (0 <= ny < ROWS): break

    if field[ny][nx] == "#":
      dir = (dir+1)%4
      continue

    x, y = nx, ny

  return False, set((x, y) for x, y, _ in cycles)

cycle, seen = walk(*start)
assert not cycle
part1 = len(seen)

part2 = 0
for x, y in seen:
  if field[y][x] not in "#^":
    field[y][x] = "#"
    cycle, _ = walk(*start)
    if cycle: part2 += 1
    field[y][x] = "."

print("Part 1:", part1)
print("Part 2:", part2)
