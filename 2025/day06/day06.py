# https://adventofcode.com/2025/day/6
# https://github.com/Favo02/advent-of-code

import sys
from itertools import groupby, zip_longest
from functools import reduce
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

lines = fin.read().strip().splitlines()

ops = [o for o in lines[-1].split() if o]
p1grid = [list(map(int, e)) for e in zip(*(filter(None, l.split()) for l in lines[:-1]))]
p2grid = [list(map(int, e)) for v, e in groupby(map(lambda e: "".join(e).strip(), zip_longest(*map(list, lines[:-1]), fillvalue=" ")), lambda x: x == "") if not v]

assert len(ops) == len(p1grid) == len(p2grid)

OPS = {
    "*": (lambda a, b: a * b, 1),
    "+": (lambda a, b: a + b, 0)
}

for op, els1, els2 in zip(ops, p1grid, p2grid):
    apply, start = OPS[op]
    part1 += reduce(apply, els1, start)
    part2 += reduce(apply, els2, start)

print("Part 1:", part1)
print("Part 2:", part2)
