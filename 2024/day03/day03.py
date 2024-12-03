# https://adventofcode.com/2024/day/3
# https://github.com/Favo02/advent-of-code

import re
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0
enabled = True
regex = re.compile(r'(mul|do(n\'t)?)\(((\d{1,3}),(\d{1,3}))?\)')

lines = []
for line in sys.stdin:
  for matchh in regex.finditer(line.strip()):
    match matchh.group(1):

      case "mul":
        val = int(matchh.group(4)) * int(matchh.group(5))
        part1 += val
        if enabled: part2 += val

      case "do": enabled = True
      case "don't": enabled = False

print("Part 1:", part1)
print("Part 2:", part2)
