# https://adventofcode.com/2024/day/19
# https://github.com/Favo02/advent-of-code

from functools import cache
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

@cache
def solve(tow, i=0):
  if i == len(tow): return 1
  if i > len(tow): return 0

  res = 0

  for avail in available:
    if len(tow) - i < len(avail):
      continue
    if tow.startswith(avail, i):
      res += solve(tow, i+len(avail))

  return res

available, tobuild = fin.read().strip().split("\n\n")
available = available.split(", ")
tobuild = tobuild.split("\n")

part1 = part2 = 0

for towel in tobuild:
  ways = solve(towel)
  if ways > 0: part1 += 1
  part2 += ways

print("Part 1:", part1)
print("Part 2:", part2)
