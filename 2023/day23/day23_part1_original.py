import advent
import sys
from collections import deque, defaultdict
import heapq

sys.setrecursionlimit(10**5)
advent.setup(2023, 23)
fin = advent.get_input()
# fin = advent.get_input("easy_input.txt")
lines = advent.get_lines(fin)

def printm(m):
  for r in m:
    for c in r:
      print(c, end="")
    print()
  print()

part1 = 0
part2 = 0

matrix = []
start = None
end = None
for i,line in enumerate(lines):
  if i == 0:
    start = (line.index("."), 0)
  if i == len(lines)-1:
    end = (line.index("."), len(matrix)-1)
  matrix.append(list(line))

assert start
assert end

W, H = len(matrix[0]), len(matrix)
DIRS = [(0,1),(0,-1),(1,0),(-1,0)]
ALLOWED_DIR = {
  ">": (+1, 0),
  "<": (-1, 0),
  "^": (0, -1),
  "v": (0, +1),
}

dist = defaultdict(int)
visited = set()
def brutesolve(cur, cur_sum):
  if cur in visited:
    return
  visited.add(cur)
  if dist[cur] < cur_sum:
    dist[cur] = cur_sum

  x,y = cur
  for dx,dy in DIRS:
    if not ((0 <= x+dx < W) and (0 <= y+dy < H)):
      continue
    assert matrix[y+dy][x+dx] in "#.^v<>"
    if matrix[y+dy][x+dx] == "#":
      continue
    if matrix[y+dy][x+dx] in ALLOWED_DIR:
      if (dx,dy) != ALLOWED_DIR[matrix[y+dy][x+dx]]:
        continue
    brutesolve((x+dx,y+dy), cur_sum+1)
  visited.remove(cur)

brutesolve(start, 1)
print(dist[end])

advent.print_answer(1, part1)
# advent.submit_answer(1, part1)

advent.print_answer(2, part2)
# advent.submit_answer(2, part2)
