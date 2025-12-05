# https://adventofcode.com/2025/day/5
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

ranges, ingrs = sys.stdin.read().strip().split("\n\n")

OPEN, CLOSE = True, False

double = []
single = []

for r in ranges.strip().splitlines():
    a, b = map(int, r.split("-"))
    double.append((a, b+1))
    single.append((a, OPEN, b+1))
    single.append((b+1, CLOSE, a))

part1 = sum(any(a <= i < b for a, b in double) for i in map(int, ingrs.splitlines()))

single.sort()
open, close = float("inf"), 0
for me, what, other in single:
    if what is OPEN:
        open = min(open, me)
        close = max(close, other)
    if what is CLOSE and close == me and open != float("inf"):
        part2 += close - open
        open, close = float("inf"), 0

print("Part 1:", part1)
print("Part 2:", part2)
