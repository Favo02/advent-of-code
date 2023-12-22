import advent
import sys
from collections import deque

advent.setup(2023, 21)
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin
lines = advent.get_lines(fin)

def pacman(x, y):
  while not (0 <= x < N):
    x = (x + N) if x < 0 else (x - N)
  while not (0 <= y < N):
    y = (y + N) if y < 0 else (y - N)
  return x, y

def bfs(matrix, cur, left):
  seen = set()
  queue = deque([(cur, left)])

  while queue:
    cur, left = queue.popleft()

    if (cur in seen) or (left < 0):
      continue
    if left%2 == 0:
      seen.add(cur)

    x, y = cur
    for dx, dy in DIRS:
      px, py = x+dx, y+dy
      if not ((0 <= x+dx < N) and (0 <= y+dy < N)):
        px, py = pacman(x+dx, y+dy)
      if matrix[py][px] == "#":
        continue
      queue.append(((x+dx, y+dy), left-1))
  return len(seen)

part1 = 0
part2 = 0

matrix = []
start = None

for y, line in enumerate(lines):
  line = list(line)
  if "S" in line:
    start = (line.index("S"), y)
    line[line.index("S")] = "."
  matrix.append(line)

DIRS = [(0,+1),(0,-1),(+1,0),(-1,0)]
N = len(matrix)

progression = []
for i in range(10):
  dist = N//2 + N*i
  progression.append(bfs(matrix, start, dist))
  print(dist, "\t", progression[-1])

print("DO THE MATHS :))))")

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
