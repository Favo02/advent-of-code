# https://adventofcode.com/2024/day/1
# https://github.com/Favo02/advent-of-code

from collections import Counter
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

lefts = []
rights = []

for line in fin:
  l, r = map(int, line.strip().split("  "))
  lefts.append(l)
  rights.append(r)

part1 = sum(abs(l - r) for l, r in zip(sorted(lefts), sorted(rights)))

rights = Counter(rights)
part2 = sum(l * rights[l] for l in lefts)

print("Part 1:", part1)
print("Part 2:", part2)
