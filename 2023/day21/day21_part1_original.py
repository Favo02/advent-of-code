import advent
import sys
from collections import deque

advent.setup(2023, 21)
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin
lines = advent.get_lines(fin)

def printm(m):
  for r in m:
    for c in r:
      print(c, end="")
    print()
  print()

def bfs(matrix, cur, left, seen):
  queue = deque([(cur, left)])

  while queue:
    cur, left = queue.popleft()

    if cur in seen:
      continue
    if left%2 == 0:
      seen.add(cur)
    if left == 0:
      continue

    W, H = len(matrix[0]), len(matrix)

    x, y = cur
    for dx, dy in DIRS:
      if not ((0 <= x+dx < W) and (0 <= y+dy < H)):
        continue
      if matrix[y+dy][x+dx] == "#":
        continue

      queue.append(((x+dx, y+dy), left-1))
  return seen

part1 = 0
part2 = 0

DIRS = [(0,+1),(0,-1),(+1,0),(-1,0)]

matrix = []
start = None

for y, line in enumerate(lines):
  line = list(line)
  if "S" in line:
    start = (line.index("S"), y)
    line[line.index("S")] = "."
  matrix.append(line)

# printm(matrix)
# print(start)

res = bfs(matrix, start, 64, set())
part1 = len(res)

print(res)


advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
