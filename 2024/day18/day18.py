# https://adventofcode.com/2024/day/18
# https://github.com/Favo02/advent-of-code

from collections import deque
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def bfs(sx, sy):
  seen = set()

  queue = deque()
  queue.append((0, sx, sy))

  while queue:
    dist, x, y = queue.popleft()

    if (x, y) == (COLS, ROWS):
      return dist

    for dx, dy in [(0,1), (0,-1), (1,0), (-1,0)]:
      nx, ny = x+dx, y+dy

      if not (0 <= nx <= COLS): continue
      if not (0 <= ny <= ROWS): continue
      if (nx, ny) in seen: continue
      if (nx, ny) in bytes: continue

      seen.add((nx, ny))
      queue.append((dist+1, nx, ny))

  return -1

# ROWS = COLS = 6 # example
ROWS = COLS = 70

# BYTES_P1 = 12 # example
BYTES_P1 = 1024

part1 = part2 = None

bytes = set()
for line in fin:
  bytes.add(tuple(map(int, line.strip().split(","))))
  if len(bytes) == BYTES_P1:
    part1 = bfs(0, 0)
  if bfs(0, 0) == -1:
    part2 = line

  if part1 and part2: break

print("Part 1:", part1)
print("Part 2:", part2)
