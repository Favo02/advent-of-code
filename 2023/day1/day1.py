# https://adventofcode.com/2023/day/1
# https://github.com/Favo02/advent-of-code

import sys

def parseInput():
  lines = []
  for line in sys.stdin:
    lines.append(line.rstrip())
  return lines

def findFirstAndLast(line, keys):
  firstI, lastI = None, None
  firstVal, lastVal = None, None

  for e in keys:
    newi = line.find(e)
    newi2 = line.rfind(e)

    if (newi != -1) and (firstI == None or newi < firstI):
      firstI = newi
      firstVal = e
    if (newi2 != -1) and (lastI == None or newi2 > lastI):
      lastI = newi2
      lastVal = e

  return firstVal, lastVal

lines = parseInput()

values = {"1":1,"2":2,"3":3,"4":4,"5":5,"6":6,"7":7,"8":8,"9":9,"0":0,"one":1, "two":2, "three":3, "four":4, "five":5, "six":6, "seven":7, "eight":8, "nine":9}

part1 = 0
part2 = 0

for l in lines:

  f1, l1 = findFirstAndLast(l, list(values.keys())[:10])
  part1 += int(f"{values[f1]}{values[l1]}")

  f2, l2 = findFirstAndLast(l, list(values.keys()))
  part2 += int(f"{values[f2]}{values[l2]}")

print("part1:",part1)
print("part2:",part2)
