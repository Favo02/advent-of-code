# https://adventofcode.com/2024/day/13
# https://github.com/Favo02/advent-of-code

import sys
import re
from sympy import Symbol, nsolve
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

part1 = 0
part2 = 0

for configuration in fin.read().strip().split("\n\n"):
  a, b, target = configuration.split("\n")
  ax, ay = map(int, re.findall(r"\d+", a))
  bx, by = map(int, re.findall(r"\d+", b))
  tx, ty = map(int, re.findall(r"\d+", target))

  n = Symbol("n")
  m = Symbol("m")

  f1 = n*ax + m*bx - tx
  f2 = n*ay + m*by - ty
  at, bt = nsolve((f1, f2), (n, m), (0, 0), prec=25)

  if abs(at - round(at)) < 1e-6 and abs(bt - round(bt)) < 1e-6:
    part1 += round(at)*3 + round(bt)

  f1 = n*ax + m*bx - (tx+10000000000000)
  f2 = n*ay + m*by - (ty+10000000000000)

  at, bt = nsolve((f1, f2), (n, m), (0, 0), prec=25)

  if abs(at - round(at)) < 1e-6 and abs(bt - round(bt)) < 1e-6:
    part2 += round(at)*3 + round(bt)

print("Part 1:", part1)
print("Part 2:", part2)
