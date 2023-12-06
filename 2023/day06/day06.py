# https://adventofcode.com/2023/day/6
# https://github.com/Favo02/advent-of-code

import sys
from math import sqrt, floor
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = 1
part2 = 0

p1times = [int(n) for n in input().split()[1:]]
p1records = [int(n) for n in input().split()[1:]]

p2time = int("".join(map(str, p1times)))
p2record = int("".join(map(str, p1records)))

def equation(time, rec):
  # (time +- (time^2 - 4*rec)) / 2
  delta = sqrt(time**2 - 4*rec)
  x1 = (time + delta) / 2
  x2 = (time - delta) / 2

  # rounding to int
  x1 = floor(x1)-1 if x1 == floor(x1) else floor(x1)
  x2 = max(0, floor(x2))

  return (x1, x2)

for time, rec in zip(p1times, p1records):
  x1, x2 = equation(time, rec)
  part1 *= (x1-x2)

x1, x2 = equation(p2time, p2record)
part2 = x1-x2

print("Part 1:", part1)
print("Part 2:", part2)
