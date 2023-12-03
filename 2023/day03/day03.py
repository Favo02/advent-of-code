# https://adventofcode.com/2023/day/3
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin
lines = []
for l in fin:
  lines.append(l.rstrip())

# search for a symbol in the 8 adjacents to x,y
def checkAdj(lines, x, y):
  for dy in range(-1,2):
    if (y == 0 and dy == -1) or (y == len(lines)-1 and dy == 1): continue

    for dx in range(-1,2):
      if (x == 0 and dx == -1) or (x == len(lines[0])-1 and dx == 1): continue
      if dx == dy == 0: continue

      if (lines[y+dy][x+dx] != ".") and (not lines[y+dy][x+dx].isdigit()):
        return True
  return False

# checks for a symbol all around items inside buffer
def checkBuf(lines, lineI, colI, buffer):
  colI -= len(buffer) # bring start to start of buffer
  for i in range(len(buffer)):
    if checkAdj(lines, x=colI+i, y=lineI):
      return True
  return False

# reconstruct a number given a line and a position of the number
# return starting coordinate of the number and the number
def reconstructNumber(line, x):
  start, end = x, x

  while start > 0 and line[start-1].isdigit():
    start -= 1
  while end < len(line)-1 and line[end+1].isdigit():
    end += 1

  return start, int("".join(line[start:end+1]))

# search for numbers in the 8 adjacents of x,y
def recSearch(lines,x,y):
  numbers = []
  foundNumberCoords = set()

  for dy in range(-1,2):
    if (y == 0 and dy == -1) or (y == len(lines)-1 and dy == 1): continue

    for dx in range(-1,2):
      if (x == 0 and dx == -1) or (x == len(lines[0])-1 and dx == 1): continue
      if dx == dy == 0: continue

      if lines[y+dy][x+dx].isdigit():
        xCoord, num = reconstructNumber(lines[y+dy],x+dx)
        numCoord = (xCoord, y+dy)

        if numCoord not in foundNumberCoords:
          foundNumberCoords.add(numCoord)
          numbers.append(num)

  return numbers

part1 = 0
part2 = 0

for lineI, line in enumerate(lines):
  buffer = []

  for colI, char in enumerate(line):

    if char.isdigit():
      buffer.append(char)

    # check valid buffer when another symbol found or end of line
    if ((not char.isdigit()) or (colI == len(line)-1)) and buffer:
      if checkBuf(lines, lineI, colI, buffer):
        part1 += int("".join(buffer))
      buffer = []

    if char == '*':
      adjNums = recSearch(lines, x=colI, y=lineI)
      if len(adjNums) == 2:
        part2 += adjNums[0] * adjNums[1]

print("Part 1:", part1)
print("Part 2:", part2)

