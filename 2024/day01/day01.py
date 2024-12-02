from sys import stdin
from collections import Counter

lefts = []
rights = []

for r, line in enumerate(stdin):
  l, r = map(int, line.strip().split("  "))
  lefts.append(l)
  rights.append(r)

part1 = sum(abs(l - r) for l, r in zip(sorted(lefts), sorted(rights)))

rights = Counter(rights)
part2 = sum(l * rights[l] for l in lefts)

print("Part 1:", part1)
print("Part 2:", part2)
