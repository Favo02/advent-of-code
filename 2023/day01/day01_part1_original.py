# https://adventofcode.com/2023/day/1
# https://github.com/Favo02/advent-of-code

import sys

def parseInput():
  lines = []
  for line in sys.stdin:
    lines.append(line.rstrip())
  return lines

lines = parseInput()
count = 0

summ = 0
for l in lines:

  d =[]
  count = 0
  for ll in l:
    if ll.isdigit():
      d.append(ll)
  print(d)
  if len(d) >= 1:
    summ += int(f"{d[0]}{d[-1]}")

print(summ)
