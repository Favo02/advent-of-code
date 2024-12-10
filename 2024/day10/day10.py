# https://adventofcode.com/2024/day/10
# https://github.com/Favo02/advent-of-code

from collections import deque
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(sx, sy):

  p1 = set()
  p2 = 0

  queue = deque()
  queue.append((sx, sy))

  while queue:
    x, y = queue.popleft()
    for dx, dy in [(0,1),(0,-1),(1,0),(-1,0)]:
      nx, ny = x+dx, y+dy

      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if not field[ny][nx].isdigit(): continue
      if int(field[ny][nx]) != int(field[y][x]) + 1: continue

      if field[ny][nx] == "9":
        p1.add((nx, ny))
        p2 += 1
      else:
        queue.append((nx, ny))

  return len(p1), p2

field = []
starts = set()

for y, line in enumerate(fin):
  line = line.rstrip()
  field.append(line)

  for x, cell in enumerate(line):
    if cell == "0":
      starts.add((x, y))

ROWS = len(field)
COLS = len(field[0])

part1 = part2 = 0

for sx, sy in starts:
  p1, p2 = bfs(sx, sy)
  part1 += p1
  part2 += p2

print("Part 1:", part1)
print("Part 2:", part2)
