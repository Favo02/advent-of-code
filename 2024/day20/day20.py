# https://adventofcode.com/2024/day/20
# https://github.com/Favo02/advent-of-code

from collections import deque, defaultdict
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(sx, sy):
  dist = defaultdict(lambda: float("inf"))
  dist[(sx, sy)] = 0

  queue = deque()
  queue.append((sx, sy))

  while queue:
    x, y = queue.popleft()

    for dx, dy in [(0,1), (0,-1), (1,0), (-1,0)]:
      nx, ny = x+dx, y+dy

      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if field[ny][nx] == "#": continue
      if (nx, ny) in dist: continue
      dist[(nx, ny)] = dist[(x, y)] + 1
      queue.append((nx, ny))
  return dist

field = []
start = end = None

for y, line in enumerate(fin):
  line = line.strip()
  field.append(line)
  if "S" in line: start = line.find("S"), y
  if "E" in line: end = line.find("E"), y

ROWS = len(field)
COLS = len(field[0])

fromstart = bfs(*start)
fromend = bfs(*end)

nocheat = fromstart[end]
part1 = part2 = 0

for y1 in range(ROWS):
  for x1 in range(COLS):
    if field[y1][x1] == "#": continue

    for y2 in range(y1-20, y1+21):
      if not (0 <= y2 < ROWS): continue
      for x2 in range(x1-20, x1+21):
        if not (0 <= x2 < COLS): continue
        if field[y2][x2] == "#": continue

        cheat = abs(x1-x2) + abs(y1-y2)

        if cheat <= 2:
          if (d := fromstart[(x1, y1)] + fromend[(x2, y2)] + cheat) <= nocheat-100:
            part1 += 1

        if cheat <= 20:
          if (d := fromstart[(x1, y1)] + fromend[(x2, y2)] + cheat) <= nocheat-100:
            part2 += 1

print("Part 1:", part1)
print("Part 2:", part2)
