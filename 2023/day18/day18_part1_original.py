import advent
from operator import itemgetter

advent.setup(2023, 18)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# fin = advent.get_input("easy_input2.txt")
lines = advent.get_lines(fin)

def printm(m):
  for r in m:
    for c in r:
      print("#" if c else "_", end="")
    print()
  print()

part1 = 0
part2 = 0

DIRS = {
  "U": (0, -1),
  "D": (0, +1),
  "L": (-1, 0),
  "R": (+1, 0)
}

cur = (0,0)
border = [cur]

for line in lines:
  dir, amt, color = line.split()
  for i in range(int(amt)):
    dirx, diry = DIRS[dir]
    curx, cury = cur
    cur = (curx + dirx, cury + diry)
    border.append(cur)

# print(border)

bounds_x = min(border)[0], max(border)[0]
bounds_y = min(border, key=itemgetter(1))[1], max(border, key=itemgetter(1))[1]
# print(bounds_x, bounds_y)

W = abs(bounds_x[0]) + abs(bounds_x[1]) + 1
H = abs(bounds_y[0]) + abs(bounds_y[1]) + 1

trasl_x, trasl_y = -(bounds_x[0]), -(bounds_y[0])
# print(W, H)

matrix = [[0 for _ in range(W)] for _ in range(H)]

for x,y in border:
  matrix[y+trasl_y][x+trasl_x] = 1

# printm(matrix)

def bfs(matrix, start):
  queue = [start]
  while queue:
    cx, cy = queue.pop(0)
    for dx, dy in DIRS.values():
      nx, ny = cx + dx, cy + dy
      if not (0 <= nx < len(matrix[0]) and 0 <= ny < len(matrix)):
        continue
      if matrix[ny][nx] == 1:
        continue
      matrix[ny][nx] = 1
      queue.append((nx,ny))
  return matrix

matrix = bfs(matrix, (155,1)) # hardcoded bfs start :)
# printm(matrix)

part1 = sum(sum(row) for row in matrix)

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
