import advent
import sys
sys.setrecursionlimit(10**5)

advent.setup(2023, 16)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def printm(m):
  for r in m:
    for c in r:
      print(c, end="")
    print()
  print()

def light(start, dir):
  if (start, dir) in calls:
    print("break already done")
    return
  else:
    calls.add((start,dir))

  x, y = start
  if not (0 <= x < len(matrix[0])):
    print(f"break bounds, {x, y}")
    return
  if not (0 <= y < len(matrix)):
    print(f"break bounds, {x, y}")
    return

  dx, dy = dir
  lighted[y][x] = 1
  mirror = matrix[y][x]
  print(f"call {start}, dir {dir}, mirr {mirror}")

  if matrix[y][x] in MIRR1:
    nx, ny = MIRR1[mirror][(-dx,-dy)]
    light((x+nx, y+ny), (nx,ny))
  elif matrix[y][x] in MIRR2:
    (dx1, dy1), (dx2, dy2) = MIRR2[mirror]
    light((x+dx1, y+dy1), (dx1,dy1))
    light((x+dx2, y+dy2), (dx2,dy2))
  else:
    light((x+dx,y+dy), dir)

part1 = 0
part2 = 0

MIRR1 = {
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
MIRR2 = {
  "|": [(0,-1), (0,+1)],
  "-": [(-1,0), (+1,0)]
}

matrix = []
for line in lines:
  matrix.append(line)

lighted = [[0]*len(matrix[0]) for _ in matrix]

calls = set()
light((0,0), (+1,0))
printm(matrix)
printm(lighted)

part1 = sum([sum(row) for row in lighted])



advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
