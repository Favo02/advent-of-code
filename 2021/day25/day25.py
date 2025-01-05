# https://adventofcode.com/2021/day/25
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

field = []

for line in fin:
  field.append(list(line.strip()))

ROWS = len(field)
COLS = len(field[0])

time = 0
while True:
  moves = False

  newfield = [row.copy() for row in field]

  for y, row in enumerate(field):
    for x, cell in enumerate(row):
      if cell == ">" and field[y][(x+1) % COLS] == ".":
        newfield[y][x] = "."
        newfield[y][(x+1) % COLS] = ">"
        moves = True

  field = newfield
  newfield = [row.copy() for row in field]

  for y, row in enumerate(field):
    for x, cell in enumerate(row):
      if cell == "v" and field[(y+1) % ROWS][x] == ".":
        newfield[y][x] = "."
        newfield[(y+1) % ROWS][x] = "v"
        moves = True

  field = newfield

  time += 1
  if not moves: break

print("Part 1:", time)
