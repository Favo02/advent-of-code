# https://adventofcode.com/2024/day/4
# https://github.com/Favo02/advent-of-code

from collections import deque
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0
queue_p1 = deque()
queue_p2 = deque()

field = []
for y, line in enumerate(fin):
  line = line.strip()
  field.append(line)

  for x, cell in enumerate(line):
    if cell == "X":
      queue_p1.append((x, y, "X", 0, 0))
    if cell == "A":
      queue_p2.append((x, y))

ROWS = len(field)
COLS = len(field[0])

advance = {"X": "M", "M": "A", "A": "S"}

while queue_p1:
  x, y, progress, dirx, diry = queue_p1.popleft()

  if progress == "S":
    part1 += 1
    continue
  elif progress == "X":
    adjs = [(0,1), (0,-1), (1,0), (-1,0), (-1,-1), (-1,1), (1,-1), (1,1)]
  else:
    adjs = [(dirx, diry)]

  for dx, dy in adjs:
    nx, ny = x + dx, y + dy
    if not (0 <= nx < COLS): continue
    if not (0 <= ny < ROWS): continue
    if field[ny][nx] != advance[progress]: continue
    queue_p1.append((nx, ny, advance[progress], dx, dy))

while queue_p2:
  x, y = queue_p2.popleft()

  found = []
  for dx, dy in [(-1,-1), (1,-1), (1,1), (-1,1)]:
    nx, ny = x + dx, y + dy
    if not (0 <= nx < COLS): continue
    if not (0 <= ny < ROWS): continue
    if field[ny][nx] not in ["M", "S"]: continue
    found.append(field[ny][nx])

  if found.count("M") != 2: continue
  if found.count("S") != 2: continue
  if not any(a == b for a, b in zip(found, found[1:])): continue

  part2 += 1

print("Part 1:", part1)
print("Part 2:", part2)
