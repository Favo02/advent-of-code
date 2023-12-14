import advent

advent.setup(2023, 14)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def printm(m):
  print()
  for r in m:
    print("".join(r))

def tilt_nw(matrix, dirx, diry):
  for y,row in enumerate(matrix):
    for x,v in enumerate(row):
      if v == "O":
        cx, cy = x, y
        while (0-dirx <= cx) and\
              (0-diry <= cy) and\
              matrix[cy+diry][cx+dirx] not in ["O", "#"]:
          cx+=dirx
          cy+=diry
        matrix[y][x] = "."
        matrix[cy][cx] = "O"
  return matrix

def tilt_se(matrix, dirx, diry):
  for y in range(len(matrix)-1, -1, -1):
    row = matrix[y]
    for x in range(len(row)-1, -1, -1):
      v = matrix[y][x]
      if v == "O":
        cx, cy = x, y
        while (cx < len(row)-dirx) and\
              (cy < len(matrix)-diry) and\
              matrix[cy+diry][cx+dirx] not in ["O", "#"]:
          cx+=dirx
          cy+=diry
        cx = min(max(0, cx), len(row)-1)
        cy = min(max(0, cy), len(matrix)-1)
        matrix[y][x] = "."
        matrix[cy][cx] = "O"
  return matrix

def cycle(matrix):
  # matrix = [row[:] for row in matrix]
  for tx,ty in [(0,-1), (-1,0)]:
    matrix = tilt_nw(matrix, tx,ty)
    # printm(matrix)
  # print("---" * 10)
  for tx,ty in [(0,+1), (+1,0)]:
    matrix = tilt_se(matrix, tx,ty)
    # printm(matrix)
  return matrix

def get_hash(matrix):
  balls = set()
  for y, row in enumerate(matrix):
    for x, val in enumerate(row):
      if val == "O":
        balls.add((x,y))
  return frozenset(balls)

def load_north(matrix):
  res = 0
  H = len(matrix)
  for y, row in enumerate(matrix):
    for x, val in enumerate(row):
      if val == "O":
        res += H-y
  return res

part1 = 0
part2 = 0

matrix = []
for line in lines:
  matrix.append(list(line))

CYCLES = 1000000000
hashes = {}
start_cycle = now = -1

for i in range(CYCLES):
  hash = get_hash(matrix)
  if hash in hashes:
    start_cycle = hashes[hash]
    now = i
    break
  hashes[hash] = i
  matrix = cycle(matrix)
  if i % 10000 == 0:print(i)

rem = (CYCLES-start_cycle) % (now-start_cycle)

for i in range(rem):
  matrix = cycle(matrix)

part2 = load_north(matrix)

print()
advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
