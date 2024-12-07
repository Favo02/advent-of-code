from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
import math
import sys
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

def conc(a, b):
  return a*(10**len(str(b))) + b

def solve(target, equation, i, curres):
  if i == len(equation):
    return [curres]

  mult = solve(target, equation, i+1, curres * equation[i])
  add = solve(target, equation, i+1, curres + equation[i])
  con = solve(target, equation, i+1, conc(curres, equation[i]))

  return mult + add + con

res = 0

field = []
for y, line in enumerate(fin):
  line = line.rstrip()

  res, vals = line.split(": ")
  res = int(res)
  vals = list(map(int, vals.split(" ")))

  field.append((res, vals))

p1 = 0

for res, eq in field:
  poss = solve(res, eq, 0, 0)
  if res in poss:
    p1 += res


print("RES:", p1)
