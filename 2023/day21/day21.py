# https://adventofcode.com/2023/day/21
# https://github.com/Favo02/advent-of-code

import sys
from collections import deque
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def pacman(x, y):
  while not (0 <= x < SIZE):
    x = (x + SIZE) if x < 0 else (x - SIZE)
  while not (0 <= y < SIZE):
    y = (y + SIZE) if y < 0 else (y - SIZE)
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
      if not ((0 <= x+dx < SIZE) and (0 <= y+dy < SIZE)):
        px, py = pacman(x+dx, y+dy)
      if matrix[py][px] == "#":
        continue
      queue.append(((x+dx, y+dy), left-1))
  return len(seen)

def find_increase_factor(nums):
  if all([i == j for i, j in zip(nums, nums[1:])]):
    return nums[0]
  diffs = [i - j for i, j in zip(nums, nums[1:])]
  return find_increase_factor(diffs)

part1 = 0
part2 = 0

matrix = []
start = None

for y, line in enumerate(fin):
  line = list(line.rstrip())
  if "S" in line:
    start = (line.index("S"), y)
    line[line.index("S")] = "."
  matrix.append(line)

DIRS = [(0,+1),(0,-1),(+1,0),(-1,0)]
SIZE = len(matrix)
assert len(matrix[0]) == len(matrix)

part1 = bfs(matrix, start, 64)

STEPS =       26501365
FULL_GRIDS =  STEPS // SIZE + 1
REMAINDER =   STEPS % SIZE
assert REMAINDER == SIZE // 2

progression = []
for i in range(4):
  dist = REMAINDER + SIZE*i
  progression.append(bfs(matrix, start, dist))

STEPS_DONE = len(progression)
MAGIC_NUMBER = find_increase_factor(progression[1:])

num = progression[-1]
diff = progression[-1] - progression[-2]
for i in range(FULL_GRIDS - STEPS_DONE):
  diff += MAGIC_NUMBER
  num += diff

part2 = num

print("Part 1:", part1)
print("Part 2:", part2)
