import advent

advent.setup(2023, 10)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
# fin = advent.get_input("easy_input2.txt")
lines = advent.get_lines(fin)

PIPES = {
  "|": [(0,-1), (0,+1)],
  "-": [(-1,0), (+1,0)],
  "L": [(0,-1), (+1,0)],
  "J": [(0,-1), (-1,0)],
  "7": [(0,+1), (-1,0)],
  "F": [(0,+1), (+1,0)]
}

ADJS = [(0,-1),(0,+1),(-1,0),(+1,0)]

def pointing_to(type, pipe, target):
  for dx,dy in PIPES[type]:
    if (pipe[0]+dx, pipe[1]+dy) == target:
      return True
  return False

def main_loop(matrix, start):
  path = [start]

  for dx,dy in ADJS:
    x,y = start[0]+dx, start[1]+dy
    if matrix[y][x] in PIPES and pointing_to(matrix[y][x], (x,y), start):
      path.append((start[0]+dx, start[1]+dy))
      # print(x,y, matrix[y][x])
      break

  cx,cy = path[-1]
  while (cx,cy) != start:
    for dx,dy in PIPES[matrix[cy][cx]]:
      if (cx+dx, cy+dy) != path[-2]:
        path.append((cx+dx, cy+dy))
        break
    cx,cy = path[-1]

  return path[:-1]

part1 = 0
part2 = 0

matrix = []
start = (-1,-1)
for y, line in enumerate(lines):
  if "S" in line:
    start = (line.index("S"), y)
  matrix.append(line)

path = main_loop(matrix, start)
part1 = len(path)//2





advent.print_answer(1, part1)
advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
