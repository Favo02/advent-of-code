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

moves = set()
def move(x, y, dx, dy):
  nx, ny = dx+x, dy+y
  if field[y][x] == "#" or field[ny][nx] == "#": return False

  # HORIZONTAL
  if dy == 0 or field[y][x] == "@":
    if field[ny][nx] == "." or move(nx, ny, dx, dy):
      moves.add((x, y, nx, ny))
      # field[ny][nx], field[y][x] = field[y][x], field[ny][nx]
      return True
    return False

  # VERTICAL
  assert dx == 0
  pair = -1 if field[y][x] == "]" else +1

  if (field[ny][nx] == "." or move(nx, ny, dx, dy)) and\
      (field[ny][nx+pair] == "." or move(nx+pair, ny, dx, dy)):

    moves.add((x, y, nx, ny))
    moves.add((x+pair, y, nx+pair, ny))
    # field[ny][nx], field[y][x] = field[y][x], field[ny][nx]
    # field[ny][nx+pair], field[y][x+pair] = field[y][x+pair], field[ny][nx+pair]
    return True
  return False

def gps():
  res = 0
  for y, row in enumerate(field):
    for x, cell in enumerate(row):
      if cell == "[":
        res += 100*y + x
  return res

res = 0

fieldd, instructions = fin.read().strip().split("\n\n")
instructions = "".join(instructions.split("\n"))

expand = {".": "..", "#": "##", "O": "[]", "@": "@."}

cx = cy = None
field = []
for y, line in enumerate(fieldd.split("\n")):
  line = line.strip()
  row = []
  for x, cell in enumerate(line):
    row += list(expand[cell])
    if cell == "@":
      cx, cy = (x*2, y)
  field.append(row)

ROWS = len(field)
COLS = len(field[0])

# print(cx)
# pprint()
# print(instructions)

DIRS = {"<": (-1,0), "^": (0,-1), ">": (+1,0), "v": (0,+1)}

i = 0
while i < len(instructions):
# while i < 100:
  # pprint()
  print(instructions[i], DIRS[instructions[i]])
  dx, dy = DIRS[instructions[i]]
  if valid := move(cx, cy, dx, dy):
    print(moves)
    for ox, oy, nx, ny in moves:
      field[oy][ox], field[ny][nx] = field[ny][nx], field[oy][ox]
      # pprint()
    cx, cy = cx+dx, cy+dy
  moves = set()
  print(valid)
  i += 1

pprint()
res = gps()

print("RES:", res)
