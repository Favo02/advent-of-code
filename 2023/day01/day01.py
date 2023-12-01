# https://adventofcode.com/202?/day/?
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def first(line, keys):
  pos = [line.find(str(k)) for k in keys]
  return keys[pos.index(min(pos, key=lambda x: x if x >= 0 else 10**9))]

DIGITS = [d for d in range(10)]

DIGITS_W = {str(d): d for d in DIGITS}
DIGITS_W.update({
  "one":    1,
  "two":    2,
  "three":  3,
  "four":   4,
  "five":   5,
  "six":    6,
  "seven":  7,
  "eight":  8,
  "nine":   9
})

part1 = 0
part2 = 0

for l in fin:
  f1 = first(l, DIGITS)
  l1 = first(l[::-1], DIGITS)
  part1 += f1*10 + l1

  f2 = DIGITS_W[first(l, list(DIGITS_W.keys()))]
  l2 = DIGITS_W[first(l[::-1], list(DIGITS_W.keys()))]
  part2 += f2*10 + l2

print("Part 1:", part1)
print("Part 2:", part2)
