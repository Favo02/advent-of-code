# https://adventofcode.com/2021/day/20
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def get_lines():
  lines = []
  for line in fin:
    lines.append(line.rstrip())
  return lines

def adjs(image, x, y, turn):
  res = []

  out = '.' if turn % 2 == 0 else '#'

  for yy in range(y-1,y+2):
    for xx in range(x-1,x+2):
      # print(xx, yy, end="->")

      if not (0 <= xx < len(image[0])):
        res.append(out)
        # print(".")
        continue
      if not (0 <= yy < len(image)):
        res.append(out)
        # print(".")
        continue

      # print(image[yy][xx])
      res.append(image[yy][xx])
  return res

def toBin(arr):
  return "".join(['0' if a == '.' else '1' for a in arr])

def toDec(binStr):
  return int(binStr, 2)

part1 = 0
part2 = 0

lines = get_lines()

ALGO = lines[0]
image = lines[2:]

for times in range(50):
  pixels = 0
  newImg = []

  for r in range(-2, len(image)+2):
    newRow = []
    for c in range(-2, len(image[0])+2):
      ind = toDec(toBin(adjs(image, c,r, times)))
      pixel = ALGO[ind]
      newRow.append(pixel)
      if pixel == "#": pixels += 1

    newImg.append(newRow)

  image = newImg
  if times == 1: part1 = pixels
  if times == 49: part2 = pixels

print("Part 1:", part1)
print("Part 2:", part2)
