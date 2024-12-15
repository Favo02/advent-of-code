from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def pprint():
  for f in field:
    print("".join(f))
  print()

def move(x, y, dx, dy):
  nx, ny = dx+x, dy+y
  if field[ny][nx] == "#": return False
  if field[ny][nx] == ".":
    field[ny][nx], field[y][x] = field[y][x], field[ny][nx]
    return True
  if move(nx, ny, dx, dy):
    field[ny][nx], field[y][x] = field[y][x], field[ny][nx]
    return True
  return False

def gps():
  res = 0
  for y, row in enumerate(field):
    for x, cell in enumerate(row):
      if cell == "O":
        res += 100*y + x
  return res

res = 0

fieldd, instructions = fin.read().strip().split("\n\n")
instructions = instructions.replace("\n", "")

cx = cy = None
field = []
for y, line in enumerate(fieldd.split("\n")):
  line = line.strip()
  field.append(list(line))
  if "@" in line:
    cx, cy = (line.find("@"), y)

ROWS = len(field)
COLS = len(field[0])

print(cx)
pprint()
print(instructions)

DIRS = {"<": (-1,0), "^": (0,-1), ">": (+1,0), "v": (0,+1)}

i = 0
while i < len(instructions):
  dx, dy = DIRS[instructions[i]]
  if move(cx, cy, dx, dy):
    cx, cy = cx+dx, cy+dy
  # pprint()
  i += 1

res = gps()

print("RES:", res)
