# https://adventofcode.com/2023/day/16
# https://github.com/Favo02/advent-of-code

import sys
sys.setrecursionlimit(10**5)
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def light(start, dir):
  if (start, dir) in calls:
    return
  else:
    calls.add((start,dir))

  x, y = start
  if not ((0 <= x < len(matrix[0])) and (0 <= y < len(matrix))):
    return

  lighted[y][x] = 1
  dx, dy = dir
  mirror = matrix[y][x]

  # 1 reflection mirror (/\)
  if matrix[y][x] in MIRRORS_1:
    nx, ny = MIRRORS_1[mirror][(-dx,-dy)]
    light((x+nx, y+ny), (nx, ny))
  # 2 reflections mirror (|-)
  elif matrix[y][x] in MIRRORS_2:
    (dx1, dy1), (dx2, dy2) = MIRRORS_2[mirror]
    light((x+dx1, y+dy1), (dx1, dy1))
    light((x+dx2, y+dy2), (dx2, dy2))
  # no reflections (.)
  else:
    light((x+dx, y+dy), dir)

part1 = 0
part2 = 0

MIRRORS_1 = {
  "\\": {
    (-1,0): (0,+1),
    (0,+1): (-1,0),
    (0,-1): (+1,0),
    (+1,0): (0,-1),
  },
  "/": {
    (-1,0): (0,-1),
    (0,+1): (+1,0),
    (0,-1): (-1,0),
    (+1,0): (0,+1),
  },
}
MIRRORS_2 = {
  "|": [(0,-1), (0,+1)],
  "-": [(-1,0), (+1,0)]
}

matrix = [line for line in fin.read().split("\n") if line]
calls = set()

W, H = len(matrix[0]), len(matrix)

for x in range(W):
  for y, dy in [(0, +1), (H-1, -1)]:
    lighted = [[0]*W for _ in matrix]
    calls.clear()
    light((x, y), (0, dy))
    res = sum([sum(row) for row in lighted])
    part2 = max(part2, res)

for y in range(H):
  for x, dx in [(0, +1), (W-1, -1)]:
    lighted = [[0]*W for _ in matrix]
    calls.clear()
    light((x, y), (dx, 0))
    res = sum([sum(row) for row in lighted])
    part2 = max(part2, res)

print("Part 1:", part1)
print("Part 2:", part2)
