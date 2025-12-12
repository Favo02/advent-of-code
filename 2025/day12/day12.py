# https://adventofcode.com/2025/day/12
# https://github.com/Favo02/advent-of-code

import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = 0

*shapes, regions = fin.read().strip().split("\n\n")
shapes = [s.split("\n")[1:] for s in shapes]
areas = [len([True for line in shape for c in line if c == "#"]) for shape in shapes]
regions = [(tuple(map(int, size.split("x"))), list(map(int, qties.split()))) for size, qties in [r.split(": ") for r in regions.split("\n")]]
part1 = len([True for ((x, y), qties) in regions if x*y >= sum(a*q for a,q in zip(areas, qties))])

print("Part 1:", part1)
