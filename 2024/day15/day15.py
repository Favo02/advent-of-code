# https://adventofcode.com/2024/day/15
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

# edit inplace the field
def move_p1(field, x, y, dx, dy):
  nx, ny = dx+x, dy+y
  if field[y][x] == "#" or field[ny][nx] == "#":
    return False
  if field[ny][nx] == ".":
    field[ny][nx], field[y][x] = field[y][x], field[ny][nx]
    return True
  if move_p1(field, nx, ny, dx, dy):
    field[ny][nx], field[y][x] = field[y][x], field[ny][nx]
    return True
  return False

# does not edit the field, adds the valid moves in `moves` set
def move_p2(field, x, y, dx, dy):
  nx, ny = dx+x, dy+y
  if field[y][x] == "#" or field[ny][nx] == "#":
    return False

  if dy == 0 or field[y][x] == "@":
    if field[ny][nx] == "." or move_p2(field, nx, ny, dx, dy):
      moves.add((x, y, nx, ny))
      return True
    return False

  pair = -1 if field[y][x] == "]" else +1

  if (field[ny][nx] == "." or move_p2(field, nx, ny, dx, dy)) and\
     (field[ny][nx+pair] == "." or move_p2(field, nx+pair, ny, dx, dy)):
    moves.add((x, y, nx, ny))
    moves.add((x+pair, y, nx+pair, ny))
    return True
  return False

def gps(field):
  res = 0
  for y, row in enumerate(field):
    for x, cell in enumerate(row):
      if cell == "[" or cell == "O":
        res += 100*y + x
  return res

raw_field, instructions = fin.read().strip().split("\n\n")
instructions = instructions.replace("\n", "")

x1 = y1 = None
x2 = y2 = None

field_p1 = []
field_p2 = []

EXPAND = {".": "..", "#": "##", "O": "[]", "@": "@."}

for y, line in enumerate(raw_field.split("\n")):
  line = line.strip()
  row_p1 = list(line)
  row_p2 = [char for cell in line for char in EXPAND[cell]]
  if "@" in row_p1: x1, y1 = row_p1.index("@"), y
  if "@" in row_p2: x2, y2 = row_p2.index("@"), y
  field_p1.append(row_p1)
  field_p2.append(row_p2)

DIRS = {"<": (-1,0), "^": (0,-1), ">": (+1,0), "v": (0,+1)}

for dir in instructions:
  dx, dy = DIRS[dir]

  if move_p1(field_p1, x1, y1, dx, dy):
    x1, y1 = x1+dx, y1+dy

  moves = set()
  if move_p2(field_p2, x2, y2, dx, dy):
    for ox, oy, nx, ny in moves:
      field_p2[oy][ox], field_p2[ny][nx] = field_p2[ny][nx], field_p2[oy][ox]
    x2, y2 = x2+dx, y2+dy

part1 = gps(field_p1)
part2 = gps(field_p2)

print("Part 1:", part1)
print("Part 2:", part2)
