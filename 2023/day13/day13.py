# https://adventofcode.com/2023/day/13
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

# bruteforce all smudges, ignoring the the old valid reflection without smudges
def smudge(matrix):
  old_reflection = reflection(matrix)
  for row in matrix:
    for x,c in enumerate(row):
      if c == "#":
        row[x] = "."
      else:
        row[x] = "#"
      res = reflection(matrix, ignore=old_reflection)
      row[x] = c
      if res > 0:
        return res
  assert False, "No smudge found"

# find a valid reflection, both horizontally and vertically
# ignore the reflection that gives as result "ignore"
def reflection(matrix, ignore=None):
  for x in range(1, len(matrix[0])):
    if is_specular(matrix, x):
      if x != ignore:
        return x
  for y in range(1, len(matrix)):
    if is_specular(matrix, y, vertical=True):
      if y*100 != ignore:
        return y*100
  return 0

# check if a matrix is specular around a pivot (ignoring extra lines)
def is_specular(matrix, pivot, vertical=False):
  if vertical:
    transpose = lambda m: [[row[x] for row in m] for x in range(len(m[0]))]
    matrix = transpose(matrix)

  for row in matrix:
    left = row[:pivot]
    right = row[pivot:]
    L = min(len(left), len(right))
    if left[::-1][:L] != right[:L]:
      return False
  return True

part1 = 0
part2 = 0

matrix = []
for line in fin:
  if line == "\n":
    part1 += reflection(matrix)
    part2 += smudge(matrix)
    matrix = []
  else:
    matrix.append(list(line.rstrip()))
part1 += reflection(matrix)
part2 += smudge(matrix)

print("Part 1:", part1)
print("Part 2:", part2)
