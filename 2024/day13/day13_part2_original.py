from heapq import heappop, heappush
from collections import defaultdict, Counter, deque
from functools import reduce, lru_cache
from math import ceil
import sys
import re
from sympy import Symbol, nsolve
fin = open(sys.argv[1]) if len(sys.argv) > 1 else sys.stdin

macs = list(fin.read().strip().split("\n\n"))
res = 0

for i, mac in enumerate(macs):
  a,b,target = mac.split("\n")

  ax, ay = map(int, re.findall(r'\d+', a))
  bx, by = map(int, re.findall(r'\d+', b))
  # tx, ty = map(lambda n: int(n), re.findall(r'\d+', target))
  tx, ty = map(lambda n: int(n) + 10000000000000, re.findall(r'\d+', target))

  n = Symbol('n')
  m = Symbol('m')

  f1 = n*ax + m*bx - tx
  f2 = n*ay + m*by - ty

  print(f"--- {i}/{len(macs)} ---")

  x = list(nsolve((f1, f2), (n, m), (0, 0), prec=25))

  valid = True
  for xx in x:
    if abs(xx - round(xx)) > 0.00000000006:
      valid = False

  if valid:
    res += round(x[0]) * 3 + round(x[1])
    print(round(x[0]), round(x[1]))
    # print("VALID")



print("RES:", res)
