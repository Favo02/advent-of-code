# https://adventofcode.com/2024/day/25
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def to_pins(pattern):
  transpose = [[row[i] for row in pattern] for i in range(len(pattern[0]))]
  return [col.count("#") for col in transpose]

def fit(l, k):
  return all(ll + kk <= 7 for ll, kk in zip(l, k))

patterns = [p.split("\n") for p in fin.read().strip().split("\n\n")]

keys = [to_pins(p) for p in patterns if all(c == "#" for c in p[-1])]
locks = [to_pins(p) for p in patterns if all(c == "#" for c in p[0])]

part1 = 0
for l in locks:
  for k in keys:
    if fit(l,k):
      part1 += 1

print("Part 1:", part1)
