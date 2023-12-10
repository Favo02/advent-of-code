# https://adventofcode.com/2023/day/10
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

# pretty print a matrix
def print_matrix(matrix):
  for l in matrix:
    for c in l:
      print(c, end="")
    print()
  print()

# check if a pipe is connected (pointing to) a target point
def is_connected_to(pipe_type, pipe_pos, target):
  px, py = pipe_pos
  for dx, dy in PIPES[pipe_type]:
    if (px+dx, py+dy) == target:
      return True
  return False

# get the main loop connected to start (and going back to start)
def main_loop(matrix, start):
  path = [start]
  sx, sy = start

  # get first pipe attached to start
  for dx, dy in ADJS:
    x, y = sx+dx, sy+dy
    if matrix[y][x] in PIPES and is_connected_to(matrix[y][x], (x, y), start):
      path.append((x, y))
      break

  # follow the pipes until coming back to start
  x, y = path[-1]
  while (x, y) != start:
    for dx, dy in PIPES[matrix[y][x]]:
      # not already in loop (not going back)
      if (x+dx, y+dy) != path[-2]:
        path.append((x+dx, y+dy))
        break
    x, y = path[-1]
  return path[:-1]

# replace the S with the correct type of pipe to close the loop
def replace_start(matrix, loop):
  sx, sy = loop[0]
  fx, fy = loop[1]
  lx, ly = loop[-1]
  dirs = {(fx-sx, fy-sy), (lx-sx, ly-sy)}

  type = None
  for pipe_type, pipe_dirs in PIPES.items():
    if dirs == set(pipe_dirs):
      type = pipe_type
      break
  assert type != None, "Cannot replace start pipe"

  matrix[sy][sx] = type
  return matrix

# replace unused pipes with '.'
def remove_unused(matrix, loop):
  for y, row in enumerate(matrix):
    for x, cell in enumerate(row):
      if ((x,y) not in loop) and (cell != "."):
        row[x] = "."
  return matrix

# zoom x2 the matrix, both width and height are doubled
# each cell of the matrix becomes a 2x2, the pipes are expanded too
# and add a border of size 1 in each direction, filled with '.'
def expand(matrix):
  HOR = ["L", "F", "-"] # pipes that expand horizontally
  VER = ["7", "F", "|"] # pipes that expand vertically

  SIZE_X =  2*(len(matrix[0]))+2

  new_matrix = []
  new_matrix.append(["."] * SIZE_X) # border

  for row in matrix:
    row1 = ["."] # border
    row2 = ["."] # border

    for cell in row:
      row1.append(cell)
      row1.append("-" if cell in HOR else ".")
      row2.append("|" if cell in VER else ".")
      row2.append(".")

    row1.append(".") # border
    row2.append(".") # border
    new_matrix.append(row1)
    new_matrix.append(row2)

  new_matrix.append(["."] * SIZE_X) # border
  return new_matrix

# flood the border of the matrix, that will flood into every cell that is not
# internal of the loop (because of the x2 expansion, this will flood also the
# cells "landlocked" by the loop but not inside to it
def flood(matrix):
  queue = [ (0,0) ]

  while queue:
    x, y = queue.pop()

    for dx, dy in ADJS:
      if (0 <= y+dy < len(matrix)) and (0 <= x+dx < len(matrix[0])):
        if matrix[y+dy][x+dx] == ".":
          matrix[y+dy][x+dx] = "_"
          queue.append((x+dx,y+dy))
  return matrix

# reverse operation of expand: halve both the height and the width of the matrix
# by discarding each even row and column
def shrink(matrix):
  return [row[1::2] for row in matrix[1::2]]

part1 = 0
part2 = 0

PIPES = {
  "|": [(0,-1), (0,+1)],
  "-": [(-1,0), (+1,0)],
  "L": [(0,-1), (+1,0)],
  "J": [(0,-1), (-1,0)],
  "7": [(0,+1), (-1,0)],
  "F": [(0,+1), (+1,0)]
}

ADJS = [
  (0,-1),
  (0,+1),
  (-1,0),
  (+1,0)
]

matrix = []
start = None

for y, line in enumerate(fin):
  line = line.rstrip()
  matrix.append(list(line))
  if "S" in line:
    start = (line.index("S"), y)

loop = main_loop(matrix, start)
part1 = len(loop) // 2
# print("original")
# print_matrix(matrix)

matrix = replace_start(matrix, loop)
# print("replace start")
# print_matrix(matrix)

matrix = remove_unused(matrix, loop)
# print("remove unused")
# print_matrix(matrix)

matrix = expand(matrix)
# print("expand")
# print_matrix(matrix)

matrix = flood(matrix)
# print("flood")
# print_matrix(matrix)

matrix = shrink(matrix)
# print("shrink")
# print_matrix(matrix)

for row in matrix:
  for cell in row:
    if cell == ".":
      part2 += 1

print("Part 1:", part1)
print("Part 2:", part2)
