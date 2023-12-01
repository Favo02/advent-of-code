# https://adventofcode.com/202?/day/?
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = 0
part2 = 0

for line in fin:
  print(line.rstrip())

print("part1:", part1)
print("part2:", part2)
