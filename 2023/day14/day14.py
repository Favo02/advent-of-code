# https://adventofcode.com/2023/day/14
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

# tilt the dish and moves all the rocks
# works only for negative dirx and diry (north and west)
def tilt_nw(matrix, dirx, diry):
  for y, row in enumerate(matrix):
    for x, v in enumerate(row):
      if v == "O":
        cx, cy = x, y
        while (0-dirx <= cx) and\
              (0-diry <= cy) and\
              matrix[cy+diry][cx+dirx] not in ["O", "#"]:
          cx += dirx
          cy += diry
        matrix[y][x] = "."
        matrix[cy][cx] = "O"
  return matrix

# tilt the dish and moves all the rocks
# works only for positive dirx and diry (south and east)
def tilt_se(matrix, dirx, diry):
  for y in range(len(matrix)-1, -1, -1):
    row = matrix[y]
    for x in range(len(row)-1, -1, -1):
      if matrix[y][x] == "O":
        cx, cy = x, y
        while (cx < len(row)-dirx) and\
              (cy < len(matrix)-diry) and\
              matrix[cy+diry][cx+dirx] not in ["O", "#"]:
          cx+=dirx
          cy+=diry
        matrix[y][x] = "."
        matrix[cy][cx] = "O"
  return matrix

# perform a full cycle (tilt N, W, S, E)
def cycle(matrix):
  for tx,ty in [(0,-1), (-1,0)]:
    matrix = tilt_nw(matrix, tx,ty)
  for tx,ty in [(0,+1), (+1,0)]:
    matrix = tilt_se(matrix, tx,ty)
  return matrix

# get the load on the north support
def north_load(m):
  return sum([H-y
              for y in range(H)
              for x in range(W)
              if m[y][x] == "O"])

part1 = 0
part2 = 0

matrix = []
for line in fin:
  matrix.append(list(line.rstrip()))

H, W = len(matrix[0]), len(matrix)

part1_matrix = [row[:] for row in matrix]
part1_matrix = tilt_nw(part1_matrix, 0, -1)
part1 = north_load(part1_matrix)

CYCLES = 1000000000
rocks_table = {}
cycle_start = cycle_end = -1

for time in range(CYCLES):
  rocks = frozenset({(x,y)
                     for y in range(H)
                     for x in range(W)
                     if matrix[y][x] == "O"})
  if rocks in rocks_table:
    cycle_start = rocks_table[rocks]
    cycle_end = time
    break
  rocks_table[rocks] = time
  matrix = cycle(matrix)

cycles_remaining = (CYCLES-cycle_start) % (cycle_end-cycle_start)

for time in range(cycles_remaining):
  matrix = cycle(matrix)

part2 = north_load(matrix)

print("Part 1:", part1)
print("Part 2:", part2)
