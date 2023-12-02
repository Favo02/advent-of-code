# https://adventofcode.com/202?/day/?
# https://github.com/Favo02/advent-of-code

import sys
from functools import reduce
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def parsePart(str):
  count = { "blue": 0, "green": 0, "red": 0 }

  tokens = str.split(" ")
  tokens = [t for t in tokens if len(t) > 0]

  for i in range(0, len(tokens), 2):
    color = tokens[i+1]
    color = color.replace(",", "")
    color = color.replace(";", "")

    count[color] += int(tokens[i])
  return count

part1 = 0
part2 = 0

PART1_MAX = {
  "red":    12,
  "green":  13,
  "blue":   14
}

for id, line in enumerate(fin):
  tokens = line.rstrip().split(": ")
  parts = tokens[1].split(";")

  part1_valid = True
  part2_value = 1

  extractionMax = { "blue": 0, "green": 0, "red": 0 }
  for p in parts:
    values = parsePart(p)

    for color, count in values.items():
      if count > PART1_MAX[color]:
        part1_valid = False
      if count > extractionMax[color]:
        extractionMax[color] = count

  if part1_valid: part1 += (id+1)
  part2 += reduce(lambda x, acc: x * acc, extractionMax.values())

print("Part 1:", part1)
print("Part 2:", part2)
