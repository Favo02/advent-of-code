# https://adventofcode.com/2024/day/12
# https://github.com/Favo02/advent-of-code

from collections import deque
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(sx, sy):
  type = field[sy][sx]

  area = set([(sx, sy)])
  perimeter = set()

  queue = deque()
  queue.append((sx, sy))
  while queue:
    x, y = queue.popleft()
    for dx, dy in [(0,1),(0,-1),(1,0),(-1,0)]:
      nx, ny = x+dx, y+dy

      if not (0 <= nx < COLS) or\
         not (0 <= ny < ROWS) or\
         field[ny][nx] != type:
        perimeter.add((nx, ny, (dx, dy)))
        continue

      if (nx, ny) in area: continue

      area.add((nx, ny))
      queue.append((nx, ny))

  return area, perimeter

def sides(perimeter):
  sid = 0

  while perimeter:
    x, y, dir = perimeter.pop()
    sid += 1

    if (x+1, y, dir) in perimeter or (x-1, y, dir) in perimeter:
      deltas = [(1,0), (-1,0)] # horizontal side
    else:
      deltas = [(0,1), (0,-1)] # vertical side

    queue = deque()
    queue.append((x, y, dir))
    while queue:
      x, y, dir = queue.popleft()
      for dx, dy in deltas:
        if (x+dx, y+dy, dir) in perimeter:
          perimeter.remove((x+dx, y+dy, dir))
          queue.append((x+dx, y+dy, dir))

  return sid

part1 = 0
part2 = 0

field = []

for line in fin:
  field.append(line.strip())

ROWS = len(field)
COLS = len(field[0])

used = set()
for y, row in enumerate(field):
  for x, cell in enumerate(row):
    if (x, y) in used: continue
    a, p = bfs(x, y)
    part1 += len(a) * len(p)
    part2 += len(a) * sides(p)

    used.update(a)

print("Part 1:", part1)
print("Part 2:", part2)
