# https://adventofcode.com/2024/day/16
# https://github.com/Favo02/advent-of-code

from heapq import heappop, heappush
from collections import defaultdict
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

# east, north, west, south
DIRS = [(1, 0), (0, -1), (-1, 0), (0, 1)]

def change_dir_cost(dir, newdir):
  return min(abs(dir - newdir), 4 - newdir + dir, 4 - dir + newdir) * 1000

def dijkstra(sx, sy, ex, ey):

  p1 = None

  # dist, x, y, direction
  queue = []
  heappush(queue, (0, sx, sy, 0))

  # x, y, direction
  dist = defaultdict(lambda: float("inf"))
  dist[(sx, sy, 0)] = 0

  while queue:
    d, x, y, dir = heappop(queue)

    if d != dist[(x, y, dir)]:
      assert d > dist[(x, y, dir)]
      continue

    if p1 is None and (x, y) == (ex, ey):
      p1 = (d, dir)

    for ndir, (dx, dy) in enumerate(DIRS):
      nx, ny = x+dx, y+dy
      if not (0 <= nx < COLS): continue
      if not (0 <= ny < ROWS): continue
      if field[ny][nx] == "#": continue

      ndist = dist[(x, y, dir)] + change_dir_cost(dir, ndir) + 1

      if ndist < dist[(nx, ny, ndir)]:
        dist[(nx, ny, ndir)] = ndist
        heappush(queue, (ndist, nx, ny, ndir))

  return *p1, dist

def backtrace(dist, x, y, dir):

  path = set()
  path.add((x, y))

  for dx, dy in DIRS:
    nx, ny = x+dx, y+dy
    if not (0 <= nx < COLS): continue
    if not (0 <= ny < ROWS): continue
    if field[ny][nx] == "#": continue

    for ndir in range(4):
      ndist = dist[(x, y, dir)] - change_dir_cost(dir, ndir) - 1

      if dist[(nx, ny, ndir)] == ndist:
        path.add((nx, ny))
        path.update(backtrace(dist, nx, ny, ndir))

  return path

sx = sy = None
ex = ey = None

field = []
for y, line in enumerate(fin):
  line = line.strip()
  field.append(line)
  if "S" in line: sx, sy = line.find("S"), y
  if "E" in line: ex, ey = line.find("E"), y

ROWS = len(field)
COLS = len(field[0])

part1, edir, dist = dijkstra(sx, sy, ex, ey)
part2 = len(backtrace(dist, ex, ey, edir))

print("Part 1:", part1)
print("Part 2:", part2)
