# https://adventofcode.com/2025/day/10
# https://github.com/Favo02/advent-of-code

import sys
from collections import deque
from functools import reduce
from z3 import Int, Optimize, Sum
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = part2 = 0

lines = list(map(lambda l: tuple(l.strip().split()), sys.stdin.read().strip().splitlines()))

def solveP1(buttons, target):
    q = deque([(0, 0)])
    seen = set([0])
    while q:
        t, cur = q.popleft()
        for b in buttons:
            next = reduce(lambda a, b: a ^ b, (1 << i for i in b), cur)
            if next == target: return t+1
            if next in seen: continue
            seen.add(next)
            q.append((t+1, next))
    assert False

def solveP2(buttons, target):
    buts = [Int(i) for i in range(len(buttons))]
    cost = Int('cost')
    opt = Optimize()

    for it, t in enumerate(target):
        opt.add(Sum([buts[ib] for ib, b in enumerate(buttons) if it in b]) == t)

    for b in buts: opt.add(b >= 0)
    opt.add(cost == Sum(buts))
    opt.minimize(cost)
    opt.check()
    return opt.model()[cost].as_long()

for l in lines:
    buttons = [tuple(map(int, b[1:-1].split(","))) for b in l[1:-1]]
    lights = reduce(lambda a, b: a | b, (1 << i for i, c in enumerate(l[0][1:-1]) if c == "#"), 0)
    joltage = tuple(map(int, l[-1][1:-1].split(",")))
    part1 += solveP1(buttons, lights)
    part2 += solveP2(buttons, joltage)

print("Part 1:", part1)
print("Part 2:", part2)
